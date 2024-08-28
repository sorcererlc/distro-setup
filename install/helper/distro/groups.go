package distro

import (
	"setup/helper"
)

func (f *DistroHelper) createGroup(g string) error {
	f.Log.Info("Adding group", g)

	err := helper.Run("sudo", "groupadd", g)
	if err != nil {
		f.Log.Error("Add group", g, err.Error())
		return err
	}

	return nil
}

func (f *DistroHelper) addUserToGroup(g string) error {
	f.Log.Info("Adding user to group", g)

	err := helper.Run("sudo", "usermod", "-aG", g, "$USER")
	if err != nil {
		f.Log.Error("Add user to group", g, err.Error())
		return err
	}

	return nil
}
