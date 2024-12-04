// Package proxy provides implementations of proxy protocols.
package proxy

import (
	"context"
	"net"
	"time"

	"git.ana/dorbmon/ivpn-gui/metadata"
	"git.ana/dorbmon/ivpn-gui/proxy/proto"
)

const (
	tcpConnectTimeout = 5 * time.Second
)

var _defaultDialer Dialer = nil

type Dialer interface {
	DialContext(context.Context, *metadata.Metadata) (net.Conn, error)
}

type Proxy interface {
	Dialer
	Addr() string
	Proto() proto.Proto
}

// SetDialer sets default Dialer.
func SetDialer(d Dialer) {
	_defaultDialer = d
}

// Dial uses default Dialer to dial TCP.
func Dial(m *metadata.Metadata) (net.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tcpConnectTimeout)
	defer cancel()
	return _defaultDialer.DialContext(ctx, m)
}

// DialContext uses default Dialer to dial TCP with context.
func DialContext(ctx context.Context, m *metadata.Metadata) (net.Conn, error) {
	return _defaultDialer.DialContext(ctx, m)
}
