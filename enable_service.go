package systemd

import (
	"fmt"
	"os/exec"
)

func EnableService(name string) error {
	out, err := exec.Command("systemctl", "enable", name).CombinedOutput()
	if err != nil {
		return fmt.Errorf("systemctl enable %s: %v: %s", name, err, out)
	}
	return nil
}
