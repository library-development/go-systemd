package systemd

import (
	"fmt"
	"os"
)

func RemoveService(name string) error {
	err := StopService(name)
	if err != nil {
		return err
	}
	err = DisableService(name)
	if err != nil {
		return err
	}
	serviceFile := "/etc/systemd/system/" + name + ".service"
	if err := os.Remove(serviceFile); err != nil {
		return fmt.Errorf("systemctl remove %s: %v", name, err)
	}
	err = DaemonReload()
	if err != nil {
		return err
	}
	return nil
}
