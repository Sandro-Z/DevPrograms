package gui

import (
	"fmt"
	"os"

	"os/exec"
)

func amAdmin() bool {
	return os.Getuid() == 0
}
func runMeElevated() {
	path, _ := os.Executable()
	cmd := exec.Command("sudo", path)
	fmt.Println(cmd.Run())
}
