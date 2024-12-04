package tunnel

import (
	"errors"
	"io"
	"net"
	"sync"
	"syscall"
	"time"

	"git.ana/dorbmon/ivpn-gui/common/pool"
	timeoutReader "git.ana/dorbmon/ivpn-gui/common/timeoutreader"
	"git.ana/dorbmon/ivpn-gui/core/adapter"
	"git.ana/dorbmon/ivpn-gui/log"
	M "git.ana/dorbmon/ivpn-gui/metadata"
	"git.ana/dorbmon/ivpn-gui/proxy"
)

const (
	tcpWaitTimeout = 5 * time.Second
)

func handleTCPConn(localConn adapter.TCPConn) {
	defer localConn.Close()

	id := localConn.ID()
	metadata := &M.Metadata{
		Network: M.TCP,
		SrcIP:   net.IP(id.RemoteAddress),
		SrcPort: id.RemotePort,
		DstIP:   net.IP(id.LocalAddress),
		DstPort: id.LocalPort,
	}

	targetConn, err := proxy.Dial(metadata)
	if err != nil {
		log.Warnf("[TCP] dial %s: %v", metadata.DestinationAddress(), err)
		return
	}

	defer targetConn.Close()
	log.Infof("[TCP] %s <-> %s", localConn.LocalAddr().String(), localConn.RemoteAddr().String())
	relay(localConn, targetConn) /* relay connections */
}

// relay copies between left and right bidirectionally.
func relay(left, right net.Conn) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		//right.SetReadDeadline(time.Now().Add(tcpWaitTimeout))
		lReader := timeoutReader.NewTimeoutReader(left)
		lReader.SetTimeout(tcpWaitTimeout)
		if err := copyBuffer(right, lReader); err != nil {
			if err.Error() == "timeout" {
				return
			}
			log.Warnf("[TCP] %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		rReader := timeoutReader.NewTimeoutReader(right)
		rReader.SetTimeout(tcpWaitTimeout)
		left.SetReadDeadline(time.Now().Add(tcpWaitTimeout))
		if err := copyBuffer(left, rReader); err != nil {
			if err.Error() == "timeout" {
				return
			}
			log.Warnf("[TCP] %v", err)
		}
	}()

	wg.Wait()
}

func copyBuffer(dst io.Writer, src io.Reader) error {
	buf := pool.Get(pool.RelayBufferSize)
	defer pool.Put(buf)

	_, err := io.CopyBuffer(dst, src, buf)
	if ne, ok := err.(net.Error); ok && ne.Timeout() {
		return nil /* ignore I/O timeout */
	} else if errors.Is(err, syscall.EPIPE) {
		return nil /* ignore broken pipe */
	} else if errors.Is(err, syscall.ECONNRESET) {
		return nil /* ignore connection reset by peer */
	}
	return err
}
