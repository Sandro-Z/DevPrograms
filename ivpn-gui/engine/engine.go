package engine

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	slog "log"

	"git.ana/dorbmon/ivpn-gui/common/pool"
	timeoutReader "git.ana/dorbmon/ivpn-gui/common/timeoutreader"
	"git.ana/dorbmon/ivpn-gui/core"
	"git.ana/dorbmon/ivpn-gui/core/device"
	"git.ana/dorbmon/ivpn-gui/core/device/tun"
	"git.ana/dorbmon/ivpn-gui/core/option"
	"git.ana/dorbmon/ivpn-gui/engine/mirror"
	"git.ana/dorbmon/ivpn-gui/log"
	"git.ana/dorbmon/ivpn-gui/metadata"
	"git.ana/dorbmon/ivpn-gui/proxy"
	"github.com/docker/go-units"
	"github.com/elazarl/goproxy"
	socks5 "github.com/things-go/go-socks5"
	"github.com/things-go/go-socks5/statute"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
)

var (
	_engineMu sync.Mutex

	// _defaultKey holds the default key for the engine.
	_defaultKey *Key

	// _defaultProxy holds the default proxy for the engine.
	_defaultProxy proxy.Proxy

	// _defaultDevice holds the default device for the engine.
	_defaultDevice device.Device

	// _defaultStack holds the default stack for the engine.
	_defaultStack *stack.Stack

	httpProxyServer *http.Server
)

// Start starts the default engine up.
func Start() {
	if err := start(); err != nil {
		log.Fatalf("[ENGINE] failed to start: %v", err)
	}
}

// Stop shuts the default engine down.
func Stop() {
	if httpProxyServer != nil {
		if err := httpProxyServer.Shutdown(context.Background()); err != nil {
			log.Fatalf("[ENGINE] failed to stop: %v", err)
		}
		httpProxyServer = nil
	}
	if err := stop(); err != nil {
		log.Fatalf("[ENGINE] failed to stop: %v", err)
	}
}

// Insert loads *Key to the default engine.
func Insert(k *Key) {
	_engineMu.Lock()
	_defaultKey = k
	_engineMu.Unlock()
}

func start() error {
	_engineMu.Lock()
	level, err := log.ParseLevel(_defaultKey.LogLevel)
	if err != nil {
		return err
	}
	log.SetGUIMode(!_defaultKey.NoGUI)
	log.SetLevel(level)
	if _defaultKey == nil {
		return errors.New("empty key")
	}

	if err := netstack(_defaultKey); err != nil {
		return err
	}

	_engineMu.Unlock()
	return nil
}

func stop() (err error) {
	_engineMu.Lock()
	if _defaultDevice != nil {
		err = _defaultDevice.Close()
	}
	if _defaultStack != nil {
		_defaultStack.Close()
		_defaultStack.Wait()
	}
	_engineMu.Unlock()
	return err
}

func createDevice(s string, mtu uint32) (device.Device, error) {
	if !strings.Contains(s, "://") {
		s = fmt.Sprintf("%s://%s", tun.Driver /* default driver */, s)
	}

	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	name := u.Host
	driver := strings.ToLower(u.Scheme)

	switch driver {
	case tun.Driver:
		return tun.Open(name, mtu)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driver)
	}
}

func netstack(k *Key) (err error) {
	if k.Proxy == "" {
		fmt.Println(*k)
		return errors.New("empty proxy")
	}
	if k.Device == "" {
		return errors.New("empty device")
	}

	if _defaultProxy, err = proxy.NewWebSocket(k.Proxy); err != nil {
		return
	}
	proxy.SetDialer(_defaultProxy)

	if _defaultDevice, err = createDevice(k.Device, uint32(k.MTU)); err != nil {
		return
	}
	if err = _defaultDevice.SetIPRoute("10.59.0.1", []string{"10.58.0.0/16"}); err != nil {
		log.Warnf("[ENGINE] failed to set IP route: %v", err)
	}

	var opts []option.Option
	if k.TCPModerateReceiveBuffer {
		opts = append(opts, option.WithTCPModerateReceiveBuffer(true))
	}

	if k.TCPSendBufferSize != "" {
		size, err := units.RAMInBytes(k.TCPSendBufferSize)
		if err != nil {
			return err
		}
		opts = append(opts, option.WithTCPSendBufferSize(int(size)))
	}

	if k.TCPReceiveBufferSize != "" {
		size, err := units.RAMInBytes(k.TCPReceiveBufferSize)
		if err != nil {
			return err
		}
		opts = append(opts, option.WithTCPReceiveBufferSize(int(size)))
	}

	if _defaultStack, err = core.CreateStack(&core.Config{
		LinkEndpoint:     _defaultDevice,
		TransportHandler: &mirror.Tunnel{},
		PrintFunc: func(format string, v ...any) {
			log.Warnf("[STACK] %s", fmt.Sprintf(format, v...))
		},
		Options: opts,
	}); err != nil {
		return
	}

	if k.EnableSocks5 {
		// then start socks5 server
		opts := []socks5.Option{socks5.WithLogger(socks5.NewLogger(slog.New(os.Stdout, "socks5: ", slog.LstdFlags))), socks5.WithConnectHandle(Socks5ConnectHandler)}
		if _defaultKey.Socks5Username != "" {
			opts = append(opts, socks5.WithAuthMethods([]socks5.Authenticator{socks5.UserPassAuthenticator{
				Credentials: socks5.StaticCredentials{
					_defaultKey.Socks5Username: _defaultKey.Socks5Password,
				},
			}}))
		}
		server := socks5.NewServer(opts...)
		go func() {
			if err := server.ListenAndServe("tcp", k.Socks5Addr); err != nil {
				log.Errorf("[SOCKS5] %s", err.Error())
			}
		}()
	}
	if k.EnableHttpProxy {
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		proxy.OnRequest()
		httpProxyServer = &http.Server{
			Addr:    "localhost:5000",
			Handler: proxy,
		}
		proxy.OnRequest().DoFunc(httpProxyHandler)
		go func() {
			if err := httpProxyServer.ListenAndServe(); err != nil {
				log.Errorf("[HTTP PROXY] %s", err.Error())
			}
		}()
	}
	log.Infof(
		"[STACK] %s://%s <-> %s://%s",
		_defaultDevice.Type(), _defaultDevice.Name(),
		_defaultProxy.Proto(), _defaultProxy.Addr(),
	)
	return nil
}
func Socks5ConnectHandler(ctx context.Context, writer io.Writer, request *socks5.Request) error {
	c, err := _defaultProxy.DialContext(ctx, &metadata.Metadata{
		Network: metadata.TCP,
		DstIP:   request.DestAddr.IP,
		DstPort: uint16(request.DestAddr.Port),
	})
	if err != nil {
		return err
	}
	defer c.Close()
	log.Infof("[socks5 TCP] %s <-> %s", request.LocalAddr.String(), request.DestAddr.String())
	if err := socks5.SendReply(writer, statute.RepSuccess, request.LocalAddr); err != nil {
		return err
	}
	// go start transfer
	newReader := timeoutReader.NewTimeoutReader(request.Reader)
	newReader.SetTimeout(5 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		buf := pool.Get(pool.RelayBufferSize)
		defer pool.Put(buf)
		if _, err := io.CopyBuffer(writer, c, buf); err != nil {
			if err.Error() == "timeout" {
				return
			}
			log.Warnf("[TCP] %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		buf := pool.Get(pool.RelayBufferSize)
		defer pool.Put(buf)
		if _, err := io.CopyBuffer(c, newReader, buf); err != nil {
			if err.Error() == "timeout" {
				return
			}
			log.Warnf("[TCP] %v", err)
		}
	}()
	wg.Wait()
	return nil
}
