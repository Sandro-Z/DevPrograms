package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"git.ana/dorbmon/ivpn-gui/engine"
	"git.ana/dorbmon/ivpn-gui/gui"
)

var (
	key = new(engine.Key)
)

func init() {
	flag.BoolVar(&key.NoGUI, "no-gui", false, "Disable GUI")
	flag.IntVar(&key.Mark, "fwmark", 0, "Set firewall MARK (Linux only)")
	flag.IntVar(&key.MTU, "mtu", 0, "Set device maximum transmission unit (MTU)")
	flag.StringVar(&key.Device, "device", "", "Use this device [driver://]name")
	flag.StringVar(&key.Interface, "interface", "", "Use network INTERFACE (Linux/MacOS only)")
	flag.StringVar(&key.LogLevel, "loglevel", "info", "Log level [debug|info|warning|error|silent]")
	flag.StringVar(&key.Proxy, "proxy", "", "Use this proxy [protocol://]host[:port]")
	flag.StringVar(&key.TCPSendBufferSize, "tcp-sndbuf", "", "Set TCP send buffer size for netstack")
	flag.StringVar(&key.TCPReceiveBufferSize, "tcp-rcvbuf", "", "Set TCP receive buffer size for netstack")
	flag.BoolVar(&key.TCPModerateReceiveBuffer, "tcp-auto-tuning", false, "Enable TCP receive buffer auto-tuning")
	flag.BoolVar(&key.EnableSocks5, "enable-socks5", false, "Enable Socks5")
	flag.StringVar(&key.Socks5Addr, "socks5-address", "", "socks5 listen address")
	flag.StringVar(&key.Socks5Username, "socks5-username", "", "socks5 username")
	flag.StringVar(&key.Socks5Password, "socks5-password", "", "socks5 password")
	flag.BoolVar(&key.EnableHttpProxy, "enable-http", false, "Enable Http Proxy")
	flag.StringVar(&key.HttpProxyAddr, "http-proxy-addr", ":1080", "Http Proxy Address")
	flag.Parse()
}

func main() {
	if !key.NoGUI {
		gui.GUIMain()
		return
	}
	engine.Insert(key)
	engine.Start()
	defer engine.Stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
