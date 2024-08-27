package packages

import (
	"os"
	"setup/log"
	"setup/types"

	"gopkg.in/yaml.v3"
)

type Pkg struct {
	Conf *types.Config
	Env  *types.Environment
	Log  *log.Log
}

func NewPkg(c *types.Config, e *types.Environment) *Pkg {
	p := Pkg{
		Conf: c,
		Env:  e,
		Log:  log.NewLog("packages.log"),
	}

	return &p
}

func (p *Pkg) loadPackageFile(fs string) (*types.Packages, error) {
	p.Log.Info("Loading package file", fs)

	f, err := os.ReadFile(fs)
	if err != nil {
		p.Log.Error("Load package file", err.Error())
		return nil, err
	}

	pkg := types.Packages{}
	err = yaml.Unmarshal(f, &pkg)
	if err != nil {
		return nil, err
	}

	return &pkg, nil
}

func (p *Pkg) loadPackages(distro string) (*types.Packages, error) {
	pkg, err := p.loadPackageFile(p.Env.Cwd + "/packages/common/packages.yml")
	if err != nil {
		return nil, err
	}

	dPkg, err := p.loadPackageFile(p.Env.Cwd + "/packages/" + distro + "/packages.yml")
	if err != nil {
		return nil, err
	}

	pkg.Repo = append(pkg.Repo, dPkg.Repo...)
	pkg.Base = append(pkg.Base, dPkg.Base...)
	pkg.Hyprland = append(pkg.Hyprland, dPkg.Hyprland...)
	pkg.Sway = append(pkg.Sway, dPkg.Sway...)
	pkg.Nvidia = append(pkg.Nvidia, dPkg.Nvidia...)
	pkg.Sddm = append(pkg.Sddm, dPkg.Sddm...)
	pkg.Bluetooth = append(pkg.Bluetooth, dPkg.Bluetooth...)
	pkg.Extras = append(pkg.Extras, dPkg.Extras...)
	pkg.Remove = append(pkg.Remove, dPkg.Remove...)
	pkg.Fonts = append(pkg.Fonts, dPkg.Fonts...)

	if dPkg.Aur != nil {
		pkg.Aur = dPkg.Aur
	}
	if dPkg.AurExtra != nil {
		pkg.AurExtra = dPkg.AurExtra
	}

	return pkg, nil
}

func (p *Pkg) SetupPackages() error {
	pkg, err := p.loadPackages(p.Env.OS.Id)
	if err != nil {
		return err
	}

	switch p.Env.OS.Id {
	case "fedora":
		d := NewFedoraHelper(p.Conf, p.Env)
		err := d.SetupPackages(pkg)
		if err != nil {
			return err
		}
		break
	case "arch":
		d := NewArchHelper(p.Conf, p.Env)
		err := d.SetupPackages(pkg)
		if err != nil {
			return err
		}
		break
	}

	return nil
}
