package distro

import (
	"os"
	"setup/helper"
	"setup/log"
	"setup/types"

	"gopkg.in/yaml.v3"
)

type DistroHelper struct {
	Conf         *types.Config
	Env          *types.Environment
	DistroConfig *types.DistroConfig
	Log          *log.Log
}

func NewDistroHelper(c *types.Config, e *types.Environment) (*DistroHelper, error) {
	f := DistroHelper{
		Conf: c,
		Env:  e,
		Log:  log.NewLog("distro-setup.log"),
	}

	err := f.LoadConfig()
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func (f *DistroHelper) LoadConfig() error {
	c := types.DistroConfig{}

	fs, err := os.ReadFile("./distro_config/distro.yml")
	if err != nil {
		f.Log.Error("Load distro config file", err.Error())
		return err
	}

	err = yaml.Unmarshal(fs, &c)

	f.DistroConfig = &c

	return nil
}

func (f *DistroHelper) SetupDistro() error {
	for _, g := range f.DistroConfig.Groups {
		err := f.createGroup(g)
		if err != nil {
			return err
		}

		err = f.addUserToGroup(g, f.Env.User.Username)
		if err != nil {
			return err
		}
	}

	if f.Conf.Packages.Bluetooth {
		for _, s := range f.DistroConfig.Services.Bluetooth {
			err := f.enableService(s)
			if err != nil {
				return err
			}
		}
	}

	if f.Conf.Packages.Sddm {
		for _, s := range f.DistroConfig.Services.Sddm {
			err := f.enableService(s)
			if err != nil {
				return err
			}
		}
	}

	err := f.detectSensors()
	if err != nil {
		return err
	}

	if f.Conf.Options.Firewall {
		err = f.setupFirewall()
		if err != nil {
			return err
		}
	}

	for _, r := range f.DistroConfig.Udev {
		err := f.setupUdevRule(r.Name, r.Rule, r.File)
		if err != nil {
			return err
		}
	}

	err = f.reloadUdevRules()
	if err != nil {
		return err
	}

	if f.Conf.Options.NetworkShares {
		err = f.setupNetworkShares()
		if err != nil {
			return err
		}
	}

	err = f.setupShell()
	if err != nil {
		return err
	}

	if f.Conf.Packages.Dotfiles {
		err := f.setupDotfiles(f.Conf.DotFilesRepo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *DistroHelper) writeFile(fs string, s string, a bool, su bool) error {
	args := []string{"echo", s}

	if su {
		args = append(args, "|", "sudo", "tee")
		if a {
			args = append(args, "-a")
		}
	} else {
		if a {
			args = append(args, ">>")
		} else {
			args = append(args, ">")
		}
	}

	args = append(args, fs)

	err := helper.Run(args...)
	if err != nil {
		return err
	}

	return nil
}
