package distro

import (
	"os"
	"setup/helper"
	"strings"

	"gopkg.in/yaml.v3"
)

func (f *DistroHelper) setupNetworkShares() error {
	f.Log.Info("Setting up network shares")

	_, err := os.Stat(f.Conf.SharesFile)
	if err != nil {
		f.Log.Warn("Shares file not found. Skipping network share configuration.")
		return nil
	}

	fs, err := os.ReadFile(f.Conf.SharesFile)
	if err != nil {
		f.Log.Error("Load shares file", err.Error())
		return err
	}

	sf := []string{}
	err = yaml.Unmarshal(fs, &sf)
	if err != nil {
		f.Log.Error("Parse shares file", err.Error())
		return err
	}

	err = helper.Run("echo", "\n\n# Network shares", "|", "sudo", "tee", "-a", "/etc/fstab")

	for _, s := range sf {
		d := strings.Split(s, " ")[1]
		err = helper.Run("sudo", "mkdir", "-p", d)
		if err != nil {
			f.Log.Error("Create mount point")
			return err
		}

		err = helper.Run("sudo", "chown", "-R", f.Env.User.Username+":"+f.Env.User.Username, d)
		if err != nil {
			f.Log.Error("Change mount point owner", err.Error())
			return err
		}

		err = helper.Run("echo", s, "|", "sudo", "tee", "-a", "/etc/fstab")
		if err != nil {
			f.Log.Error("Add network share")
			return err
		}
	}

	return nil
}
