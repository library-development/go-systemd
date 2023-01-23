package systemd

import (
	"fmt"
	"os/exec"
)

func StartService(name string) error {
	out, err := exec.Command("systemctl", "start", name).CombinedOutput()
	if err != nil {
		return fmt.Errorf("systemctl start %s: %v: %s", name, err, out)
	}
	return nil
}
