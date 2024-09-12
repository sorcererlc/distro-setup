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
		Base []struct {
			Name string `yaml:"name"`
			Url  string `yaml:"url"`
		} `yaml:"base"`
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
	err := f.installRepos()
	if err != nil {
		return err
	}

	err = f.loadPackages()
	if err != nil {
		return err
	}

	p := f.Packages.Base
	if f.Config.Flatpak.Packages.Devel {
		p = append(p, f.Packages.Devel...)
	}
	if f.Config.Flatpak.Packages.Extras {
		p = append(p, f.Packages.Extras...)
	}
	if f.Config.Flatpak.Packages.Misc {
		p = append(p, f.Packages.Misc...)
	}

	for _, pkg := range p {
		err := f.installPackage(pkg)
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

func (f *FlatpakHelper) installPackage(p string) error {
	f.Log.Info("Install Flatpak package", p)

	err := helper.Run("flatpak", "install", "-y", p)
	if err != nil {
		f.Log.Warn("Install Flatpak package", p, err.Error())
	}

	return nil
}

func (f *FlatpakHelper) installRepos() error {
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

	for _, r := range f.Repos.Base {
		err = helper.Run("flatpak", "remote-add", "--if-not-exists", r.Name, r.Url)
		if err != nil {
			f.Log.Error("Install Flatpak repo", err.Error())
			return err
		}
	}

	return nil
}
