package distro

import (
	"setup/helper"
	"strings"
)

func (f *DistroHelper) setupFirewall() error {
	f.Log.Info("Setting up firewall rules")

	err := helper.Run("sudo", "ufw", "reset")
	for _, r := range f.DistroConfig.Firewall {
		a := strings.Split(r, " ")
		err = helper.Run("sudo", "ufw", a[0], a[1])
	}
	err = helper.Run("sudo", "ufw", "enable")
	if err != nil {
		f.Log.Warn("Firewall rule setup", err.Error())
	}

	return nil
}
