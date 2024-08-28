package distro

import (
	"setup/helper"
	"strings"
)

func (f *DistroHelper) setupUdevRule(n string, r string, fs string) error {
	f.Log.Info("Setting up udev rule", n)

	r = strings.ReplaceAll(r, "$USER_GID", f.Env.User.Gid)

	err := f.writeFile(fs, r, false, true)
	if err != nil {
		f.Log.Error("Setup udev rule", n, err.Error())
		return err
	}

	return nil
}

func (f *DistroHelper) reloadUdevRules() error {
	f.Log.Info("Reloading udev rules")

	err := helper.Run("sudo", "udevadm", "control", "--reload")
	if err != nil {
		f.Log.Error("Reload udev rules", err.Error())
		return err
	}

	return nil
}
