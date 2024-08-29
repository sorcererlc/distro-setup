package packages

import (
	"os"
	"setup/helper"
	"setup/log"
	"setup/types"
	"strings"

	"gopkg.in/yaml.v3"
)

type FlatpakHelper struct {
	Config *types.Config
	Env    *types.Environment
	Log    *log.Log
	Repos  struct {
		Base []string `yaml:"base"`
	}
	Packages struct {
		Base   []string `yaml:"base"`
		Devel  []string `yaml:"devel"`
		Extras []string `yaml:"extras"`
		Misc   []string `yaml:"misc"`
	}
}

func NewFlatpakHelper(c *types.Config, e *types.Environment) *FlatpakHelper {
	f := FlatpakHelper{
		Config: c,
		Env:    e,
		Log:    log.NewLog("flatpak.log"),
	}

	return &f
}

func (f *FlatpakHelper) InstallPackages() error {
	err := f.loadPackages()
	if err != nil {
		return nil
	}

	err = f.installPackageGroup(f.Packages.Base)
	if err != nil {
		return err
	}
	if f.Config.Flatpak.Packages.Devel {
		err := f.installPackageGroup(f.Packages.Devel)
		if err != nil {
			return err
		}
	}
	if f.Config.Flatpak.Packages.Extras {
		err := f.installPackageGroup(f.Packages.Extras)
		if err != nil {
			return err
		}
	}
	if f.Config.Flatpak.Packages.Misc {
		err := f.installPackageGroup(f.Packages.Misc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *FlatpakHelper) loadPackages() error {
	fs, err := os.ReadFile(f.Env.Cwd + "/packages/flatpak/packages.yml")
	if err != nil {
		f.Log.Error("Read Flatpak packages", err.Error())
		return nil
	}

	err = yaml.Unmarshal(fs, &f.Packages)
	if err != nil {
		f.Log.Error("Parse Flatpak packages", err.Error())
		return nil
	}

	return nil
}

func (f *FlatpakHelper) installPackageGroup(g []string) error {
	f.Log.Info("Installing Flatpak packages", strings.Join(g, ", "))

	args := []string{"flatpak", "install", "-y"}
	args = append(args, g...)
	err := helper.Run(args...)
	if err != nil {
		f.Log.Error("Install Flatpak packages", err.Error())
		return err
	}

	return nil
}

func (f *FlatpakHelper) InstallRepos() error {
	fs, err := os.ReadFile(f.Env.Cwd + "/packages/flatpak/repos.yml")
	if err != nil {
		f.Log.Error("Read Flatpak repo file", err.Error())
		return err
	}

	err = yaml.Unmarshal(fs, &f.Repos)
	if err != nil {
		f.Log.Error("Parse Flatpak repo file", err.Error())
		return err
	}

	args := []string{"flatpak", "remote-add", "--if-not-exists"}
	args = append(args, f.Repos.Base...)
	err = helper.Run(args...)
	if err != nil {
		f.Log.Error("Install Flatpak repos", err.Error())
		return nil
	}

	return nil
}
