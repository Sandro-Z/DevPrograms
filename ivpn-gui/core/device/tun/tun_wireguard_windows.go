package tun

import (
	"fmt"
	"os/exec"

	"git.ana/dorbmon/ivpn-gui/log"
)

const (
	offset     = 0
	defaultMTU = 0 /* auto */
)

func (t *TUN) SetIPRoute(ip string, routes []string) error {
	log.Debugf("SetIPRoute: dev=%s, ip=%s, routes=%v", t.name, ip, routes)
	if err := exec.Command("netsh", "interface", "ipv4", "set", "address", t.name, "static", ip).Run(); err != nil {
		return fmt.Errorf("netsh: %w", err)
	}
	for _, route := range routes {
		if err := exec.Command("netsh", "interface", "ipv4", "add", "route", route, t.name).Run(); err != nil {
			return fmt.Errorf("netsh: %w", err)
		}
	}
	return nil
}
