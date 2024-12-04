package proxy

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"

	M "git.ana/dorbmon/ivpn-gui/metadata"
	"git.ana/dorbmon/ivpn-gui/proxy/proto"
	"golang.org/x/net/websocket"
)

type WebSocket struct {
	addr  string
	token string
	url   string
	proto proto.Proto
}

func NewWebSocket(s string) (Proxy, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	p := &WebSocket{}
	if u.Scheme == "ws" {
		p.proto = proto.WebSocket
	} else if u.Scheme == "wss" {
		p.proto = proto.WebSocketSecure
	} else {
		return nil, fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}
	p.addr = u.Host
	p.token = "Bearer " + u.User.String()
	u.User = nil
	p.url = u.String()
	return p, nil
}

func (p *WebSocket) Addr() string {
	return p.addr
}

func (p *WebSocket) Proto() proto.Proto {
	return p.proto
}

func (p *WebSocket) DialContext(ctx context.Context, metadata *M.Metadata) (c net.Conn, err error) {
	headers := http.Header{
		"Authorization": {p.token},
		"Dest":          {net.JoinHostPort(metadata.DstIP.String(), strconv.FormatUint(uint64(metadata.DstPort), 10))},
	}
	config, err := websocket.NewConfig(p.url, "http://localhost/")
	if err != nil {
		return nil, fmt.Errorf("connect to %s: %w", p.url, err)
	}

	config.Header = headers
	config.Protocol = []string{"proxy"}
	c, err = websocket.DialConfig(config)
	if err != nil {
		return nil, fmt.Errorf("connect to %s: %w", p.url, err)
	}

	return
}
