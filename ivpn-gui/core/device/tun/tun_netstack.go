//go:build linux

package tun

import (
	"fmt"
	"net"
	"os/exec"
	"unsafe"

	"git.ana/dorbmon/ivpn-gui/core/device"
	"git.ana/dorbmon/ivpn-gui/log"
	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/tcpip/link/fdbased"
	"gvisor.dev/gvisor/pkg/tcpip/link/rawfile"
	"gvisor.dev/gvisor/pkg/tcpip/link/tun"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
)

type TUN struct {
	stack.LinkEndpoint

	fd   int
	mtu  uint32
	name string
}

func Open(name string, mtu uint32) (device.Device, error) {
	t := &TUN{name: name, mtu: mtu}

	if len(t.name) >= unix.IFNAMSIZ {
		return nil, fmt.Errorf("interface name too long: %s", t.name)
	}

	fd, err := tun.Open(t.name)
	if err != nil {
		return nil, fmt.Errorf("create tun: %w", err)
	}
	t.fd = fd

	if t.mtu > 0 {
		if err := setMTU(t.name, t.mtu); err != nil {
			return nil, fmt.Errorf("set mtu: %w", err)
		}
	}

	_mtu, err := rawfile.GetMTU(t.name)
	if err != nil {
		return nil, fmt.Errorf("get mtu: %w", err)
	}
	t.mtu = _mtu

	ep, err := fdbased.New(&fdbased.Options{
		FDs: []int{fd},
		MTU: t.mtu,
		// TUN only, ignore ethernet header.
		EthernetHeader: false,
		// SYS_READV support only for TUN fd.
		PacketDispatchMode: fdbased.Readv,
		// TAP/TUN fd's are not sockets and using the WritePackets calls results
		// in errors as it always defaults to using SendMMsg which is not supported
		// for tap/tun device fds.
		//
		// This CL changes WritePackets to gracefully degrade to using writev instead
		// of sendmmsg if the underlying fd is not a socket.
		//
		// Fixed: https://github.com/google/gvisor/commit/f33d034fecd7723a1e560ccc62aeeba328454fd0
		MaxSyscallHeaderBytes: 0x00,
	})
	if err != nil {
		return nil, fmt.Errorf("create endpoint: %w", err)
	}
	t.LinkEndpoint = ep

	return t, nil
}

func (t *TUN) Name() string {
	return t.name
}

func (t *TUN) Close() error {
	return unix.Close(t.fd)
}

// Ref: wireguard tun/tun_linux.go setMTU.
func setMTU(name string, n uint32) error {
	// open datagram socket
	fd, err := unix.Socket(
		unix.AF_INET,
		unix.SOCK_DGRAM,
		0,
	)
	if err != nil {
		return err
	}

	defer unix.Close(fd)

	const ifReqSize = unix.IFNAMSIZ + 64

	// do ioctl call
	var ifr [ifReqSize]byte
	copy(ifr[:], name)
	*(*uint32)(unsafe.Pointer(&ifr[unix.IFNAMSIZ])) = n
	_, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(fd),
		uintptr(unix.SIOCSIFMTU),
		uintptr(unsafe.Pointer(&ifr[0])),
	)

	if errno != 0 {
		return fmt.Errorf("failed to set MTU: %w", errno)
	}

	return nil
}

func (t *TUN) SetIPRoute(ip string, routes []string) error {
	return setIPRoute(t.name, ip, routes)
}

func setIPRoute(dev string, ip string, routes []string) error {
	log.Debugf("SetIPRoute: dev=%s, ip=%s, routes=%v", dev, ip, routes)
	if err := bringInterfaceUP(dev); err != nil {
		return fmt.Errorf("netlink: %w", err)
	}
	if err := setIP(dev, ip); err != nil {
		return fmt.Errorf("netlink: %w", err)
	}
	for _, route := range routes {
		if err := setRoute(dev, route); err != nil {
			return fmt.Errorf("ip: %w", err)
		}
	}
	return nil
}

func bringInterfaceUP(name string) error {
	const ifReqSize = unix.IFNAMSIZ + 64
	var ifr [ifReqSize]byte
	copy(ifr[:], name)
	*(*uint32)(unsafe.Pointer(&ifr[unix.IFNAMSIZ])) = unix.IFF_UP
	return callIOCTL(&ifr, unix.SIOCSIFFLAGS)
}

func setIP(name string, ip string) error {
	nip := net.ParseIP(ip)
	nip = nip.To4()

	addr := unix.RawSockaddrInet4{}
	addr.Family = unix.AF_INET
	copy(addr.Addr[:], nip)
	const ifReqSize = unix.IFNAMSIZ + 64
	var ifr [ifReqSize]byte
	copy(ifr[:], name)
	*(*unix.RawSockaddrInet4)(unsafe.Pointer(&ifr[unix.IFNAMSIZ])) = addr
	return callIOCTL(&ifr, unix.SIOCSIFADDR)
}

func setRoute(name string, route string) error {
	//TODO:
	return exec.Command("ip", "route", "add", route, "dev", name).Run()
	//TODO:not implemented
	// return nil
}

func callIOCTL(ifr *[unix.IFNAMSIZ + 64]byte, t uintptr) error {
	// open datagram socket
	fd, err := unix.Socket(
		unix.AF_INET,
		unix.SOCK_DGRAM,
		0,
	)
	if err != nil {
		return err
	}

	defer unix.Close(fd)

	// do ioctl call
	_, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(fd),
		uintptr(t),
		uintptr(unsafe.Pointer(&ifr[0])),
	)

	if errno != 0 {
		return fmt.Errorf("failed to do ioctl call MTU: %w", errno)
	}

	return nil
}