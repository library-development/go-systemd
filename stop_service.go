package systemd

import (
	"fmt"
	"os/exec"
)

func StopService(name string) error {
	out, err := exec.Command("systemctl", "stop", name).CombinedOutput()
	if err != nil {
		return fmt.Errorf("systemctl stop %s: %v: %s", name, err, out)
	}
	return nil
}
