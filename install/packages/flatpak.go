package packages

import (
	"fmt"
	"os"
	"setup/helper"
	"setup/types"
	"strings"

	"gopkg.in/yaml.v3"
)

type FlatpakHelper struct {
	Cwd    string
	Config *types.Config
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

func NewFlatpakHelper(c *types.Config) (*FlatpakHelper, error) {
	var err error
	f := FlatpakHelper{
		Config: c,
	}

	f.Cwd, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func (f *FlatpakHelper) loadPackages() error {
	fs, err := os.ReadFile(f.Cwd + "/packages/flatpak/packages.yml")
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal(fs, &f.Packages)
	if err != nil {
		return nil
	}

	return nil
}

func (f *FlatpakHelper) installPackageGroup(g []string) error {
	fmt.Printf("Installing Flatpak packages %s\n", strings.Join(g, ", "))

	c := helper.Cmd{
		Bin:  "flatpak",
		Args: []string{"install", "-y"},
	}
	c.Args = append(c.Args, g...)

	err := helper.ExecuteCommand(c)
	if err != nil {
		fmt.Printf("Error installing Flatpak packages. Aborting setup.\n%s", err.Error())
		return err
	}

	return nil
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

func (f *FlatpakHelper) InstallRepos() error {
	fs, err := os.ReadFile(f.Cwd + "/packages/flatpak/repos.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fs, &f.Repos)

	c := helper.Cmd{
		Bin:  "flatpak",
		Args: []string{"remote-add", "--if-not-exists"},
	}
	c.Args = append(c.Args, f.Repos.Base...)

	err = helper.ExecuteCommand(c)
	if err != nil {
		return nil
	}

	return nil
}
