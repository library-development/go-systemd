package systemd

import (
	"fmt"
	"os/exec"
)

func DisableService(name string) error {
	out, err := exec.Command("systemctl", "disable", name).CombinedOutput()
	if err != nil {
		return fmt.Errorf("systemctl disable %s: %v: %s", name, err, out)
	}
	return nil
}
