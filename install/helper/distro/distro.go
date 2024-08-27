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

		err = f.addUserToGroup(g)
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

	if f.Conf.Packages.Dotfiles {
		err := f.setupDotfiles(f.Conf.DotFilesRepo)
		if err != nil {
			return err
		}
	}

	return nil
}

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

func (f *DistroHelper) enableService(s string) error {
	f.Log.Info("Enabling service", s)

	err := helper.Run("sudo", "systemctl", "enable", s)
	if err != nil {
		f.Log.Error("Enable service", s, err.Error())
		return err
	}

	return nil
}

func (f *DistroHelper) setupDotfiles(p types.GitPackage) error {
	f.Log.Info("Setting up dotfiles")

	args := []string{"git", "clone", "--depth", "1"}
	if p.Tag != "" {
		args = append(args, "--branch", p.Tag)
	}
	args = append(args, p.Url)

	err := helper.Run(args...)
	if err != nil {
		f.Log.Error("Clone dotfiles repo", err.Error())
		return err
	}

	_ = helper.Run("cd", "dotfiles")

	err = helper.Run("make", "run")
	if err != nil {
		f.Log.Error("Setup dotfiles", err.Error())
		return err
	}

	return nil
}
