package distro

import (
	"setup/helper"
)

func (f *DistroHelper) detectSensors() error {
	_ = helper.Run("clear")

	f.Log.Info("Detecting hardware sensors")

	err := helper.Run("sudo", "sensors-detect")
	if err != nil {
		f.Log.Error("Sensors detect", err.Error())
		return err
	}

	return nil
}
