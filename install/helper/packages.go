package helper

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Packages struct {
	Repo      []string `yaml:"repo"`
	Base      []string `yaml:"base"`
	Hyprland  []string `yaml:"hyprland"`
	Sway      []string `yaml:"sway"`
	Nvidia    []string `yaml:"nvidia"`
	Sddm      []string `yaml:"sddm"`
	Bluetooth []string `yaml:"bluetooth"`
	Extras    []string `yaml:"extras"`
	Aur       []string `yaml:"aur,omitempty"`
	AurExtra  []string `yaml:"aur_extra,omitempty"`
	Remove    []string `yaml:"remove,omitempty"`
	Fonts     []string `yaml:"fonts"`
}

func LoadPackages(distro string) (*Packages, *Packages, error) {
	c, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	fs := c + "/packages/common/packages.yml"
	println("Loading package file " + fs)

	f, err := os.ReadFile(fs)
	if err != nil {
		return nil, nil, err
	}

	p := Packages{}
	err = yaml.Unmarshal(f, &p)
	if err != nil {
		return nil, nil, err
	}

	fs = c + "/packages/" + distro + "/packages.yml"
	println("Loading package file " + fs)

	f, err = os.ReadFile(fs)
	if err != nil {
		return nil, nil, err
	}

	dp := Packages{}
	err = yaml.Unmarshal(f, &dp)
	if err != nil {
		return nil, nil, err
	}

	return &p, &dp, nil
}

func InstallRepos(distro string, dv string, pkg []string) error {
	switch distro {
	case "fedora":
		d, err := NewFedoraHelper()
		err = d.InstallRepos(dv, pkg)
		if err != nil {
			return err
		}

		err = d.EnableCoprRepos()
		if err != nil {
			return err
		}
	}

	return nil
}

func InstallPackages(distro string, pkg []string) error {
	switch distro {
	case "fedora":
		d := FedoraHelper{}
		err := d.InstallPackages(pkg)
		if err != nil {
			return err
		}
		break
	case "arch":
		d := ArchHelper{}
		err := d.InstallPackages(pkg)
		if err != nil {
			return err
		}
		break
	}

	return nil
}

func RemovePackages(distro string, pkg []string) error {
	switch distro {
	case "fedora":
		d := FedoraHelper{}
		err := d.RemovePackages(pkg)
		if err != nil {
			return err
		}
		break
	case "arch":
		d := ArchHelper{}
		err := d.RemovePackages(pkg)
		if err != nil {
			return err
		}
		break
	}

	return nil
}
