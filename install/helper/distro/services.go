package distro

import (
	"setup/helper"
)

func (f *DistroHelper) enableService(s string) error {
	f.Log.Info("Enabling service", s)

	err := helper.Run("sudo", "systemctl", "enable", s)
	if err != nil {
		f.Log.Error("Enable service", s, err.Error())
		return err
	}

	return nil
}
