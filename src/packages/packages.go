package packages

import (
	"os"
	"setup/log"
	"setup/types"

	"gopkg.in/yaml.v3"
)

type Handler interface {
	SetupDistro() error
	InstallBasePackages() error
	InstallExtraPackages() error
	InstallNvidia() error
	InstallSddm() error
	InstallHyprland() error
	InstallBluetooth() error
}

type Pkg struct {
	Conf    *types.Config
	Env     *types.Environment
	Log     *log.Log
	Handler Handler
}

func NewPkg(c *types.Config, e *types.Environment) (*Pkg, error) {
	p := Pkg{
		Conf: c,
		Env:  e,
		Log:  log.NewLog("packages.log"),
	}

	pkg, err := p.loadPackages(e.OS.Id)
	if err != nil {
		return nil, err
	}

	switch e.OS.Id {
	case "fedora":
		p.Handler = NewFedoraHelper(c, e, pkg)
		break
	case "arch":
		p.Handler = NewArchHelper(c, e, pkg)
		break
	}

	return &p, nil
}

func (p *Pkg) loadPackageFile(fs string) (*types.Packages, error) {
	p.Log.Debug("Loading package file", fs)

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

	if pkg.Git == nil {
		pkg.Git = make(map[string]types.GitPackage)
	}
	for k, v := range dPkg.Git {
		pkg.Git[k] = v
	}

	if dPkg.Aur != nil {
		pkg.Aur = dPkg.Aur
	}
	if dPkg.AurExtra != nil {
		pkg.AurExtra = dPkg.AurExtra
	}

	return pkg, nil
}

func (p *Pkg) SetupDistro() error {
	err := p.Handler.SetupDistro()
	if err != nil {
		return err
	}
	return nil
}

func (p *Pkg) InstallBasePackages() error {
	err := p.Handler.InstallBasePackages()
	if err != nil {
		return err
	}
	return nil
}

func (p *Pkg) InstallExtraPackages() error {
	err := p.Handler.InstallExtraPackages()
	if err != nil {
		return err
	}
	return nil
}

func (p *Pkg) InstallNvidia() error {
	err := p.Handler.InstallNvidia()
	if err != nil {
		return err
	}
	return nil
}

func (p *Pkg) InstallSddm() error {
	err := p.Handler.InstallSddm()
	if err != nil {
		return err
	}
	return nil
}

func (p *Pkg) InstallHyprland() error {
	err := p.Handler.InstallHyprland()
	if err != nil {
		return err
	}
	return nil
}

func (p *Pkg) InstallBluetooth() error {
	err := p.Handler.InstallBluetooth()
	if err != nil {
		return err
	}
	return nil
}
