package systemd

import (
	"fmt"
	"os"
)

func AddService(name, description, execStart string) error {
	serviceFile := "/etc/systemd/system/" + name + ".service"
	_, err := os.Stat(serviceFile)
	if err == nil {
		return fmt.Errorf("service %s already exists", name)
	}
	b := []byte(fmt.Sprintf(`[Unit]
	Description=%s
	
	[Service]
	ExecStart=%s
	
	[Install]
	WantedBy=multi-user.target
`, description, execStart))
	err = os.WriteFile(serviceFile, b, 0644)
	if err != nil {
		return err
	}

	// Reload systemd.
	err = DaemonReload()
	if err != nil {
		return err
	}

	// Enable the service.
	err = EnableService(name)
	if err != nil {
		return err
	}

	// Start the service.
	err = StartService(name)
	if err != nil {
		return err
	}

	return nil
}
