package systemd

import (
	"fmt"
	"os/exec"
)

func DaemonReload() error {
	out, err := exec.Command("systemctl", "daemon-reload").CombinedOutput()
	if err != nil {
		return fmt.Errorf("systemctl daemon-reload: %v: %s", err, out)
	}
	return nil
}
