//go:build !linux && !windows

package tun

import (
	"fmt"
	"os/exec"

	"git.ana/dorbmon/ivpn-gui/log"
)

const (
	offset     = 4 /* 4 bytes TUN_PI */
	defaultMTU = 1500
)

func (t *TUN) SetIPRoute(ip string, routes []string) error {
	log.Debugf("SetIPRoute: dev=%s, ip=%s, routes=%v", t.name, ip, routes)
	if err := exec.Command("ifconfig", t.name, ip, ip, "up").Run(); err != nil {
		return fmt.Errorf("ifconfig: %w", err)
	}
	for _, route := range routes {
		if err := exec.Command("route", "add", "-net", route, ip).Run(); err != nil {
			return fmt.Errorf("route: %w", err)
		}
	}
	return nil
}
