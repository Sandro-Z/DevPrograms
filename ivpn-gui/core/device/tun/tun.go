// Package tun provides TUN which implemented device.Device interface.
package tun

import (
	"git.ana/dorbmon/ivpn-gui/core/device"
)

const Driver = "tun"

func (t *TUN) Type() string {
	return Driver
}

var _ device.Device = (*TUN)(nil)
