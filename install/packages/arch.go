package packages

import (
	"errors"
	"fmt"
	"setup/helper"
	"setup/log"
	"setup/types"
	"strings"
)

type ArchHelper struct {
	Conf      *types.Config
	Env       *types.Environment
	Log       *log.Log
	CoprRepos struct {
		Copr []string `yaml:"copr"`
	}
}

func NewArchHelper(c *types.Config, e *types.Environment) *ArchHelper {
	f := ArchHelper{
		Conf: c,
		Env:  e,
		Log:  log.NewLog("arch_packages.log"),
	}

	return &f
}

func (f *ArchHelper) SetupPackages(pkg *types.Packages) error {
	err := f.updateDistro()
	if err != nil {
		return err
	}

	err = f.removePackages(pkg.Remove)
	if err != nil {
		return err
	}

	err = f.installRepos(pkg.Repo)
	if err != nil {
		return err
	}

	p := pkg.Base
	if f.Conf.Packages.Extras {
		p = append(p, pkg.Extras...)
	}
	if f.Conf.Packages.Sddm {
		p = append(p, pkg.Sddm...)
	}
	if f.Conf.Packages.Bluetooth {
		p = append(p, pkg.Bluetooth...)
	}
	if f.Conf.Packages.Nvidia {
		p = append(p, pkg.Nvidia...)
	}
	p = append(p, pkg.Fonts...)

	err = f.installPackages(p)
	if err != nil {
		return err
	}

	p = pkg.Aur
	if f.Conf.Packages.Extras {
		p = append(p, pkg.AurExtra...)
	}

	err = f.installAurPackages(p)
	if err != nil {
		return err
	}

	return nil
}

func (f *ArchHelper) updateDistro() error {
	f.Log.Info("Updating packages")

	c := helper.Cmd{
		Bin:  "sudo",
		Args: []string{"pacman", "-Syu"},
	}

	err := helper.ExecuteCommand(c)
	if err != nil {
		f.Log.Error("Update packages", err.Error())
		return err
	}

	return nil
}

func (f *ArchHelper) installRepos(r []string) error {
	for i := 0; i < len(r); i++ {
		r[i] = fmt.Sprintf(r[i], f.Env.OS.VersionId)
	}
	f.Log.Info("Installing repositories", strings.Join(r, ", "))

	c := helper.Cmd{
		Bin:  "sudo",
		Args: []string{"pacman", "-Sy"},
	}
	c.Args = append(c.Args, r...)

	err := helper.ExecuteCommand(c)
	if err != nil {
		f.Log.Error("Install repositories", err.Error())
		return err
	}

	return nil
}

func (f *ArchHelper) removePackages(p []string) error {
	f.Log.Info("Removing packages", strings.Join(p, ", "))

	c := helper.Cmd{
		Bin:  "sudo",
		Args: []string{"pacman", "-Rsy"},
	}
	c.Args = append(c.Args, p...)

	err := helper.ExecuteCommand(c)
	if err != nil {
		f.Log.Error("Remove packages", err.Error())
		return err
	}

	return nil
}

func (f *ArchHelper) installPackages(p []string) error {
	f.Log.Info("Installing packages", strings.Join(p, ", "))

	c := helper.Cmd{
		Bin:  "sudo",
		Args: []string{"pacman", "-Sy"},
	}
	c.Args = append(c.Args, p...)

	err := helper.ExecuteCommand(c)
	if err != nil {
		f.Log.Error("Install packages", err.Error())
		return err
	}

	for _, pk := range p {
		i := f.checkInstalledPackage(pk)
		if !i {
			f.Log.Error("Package " + pk + " failed to install. Aborting setup.")
			return errors.New("")
		}
	}

	return nil
}

func (f *ArchHelper) installAurPackages(p []string) error {
	f.Log.Info("Installing AUR packages", strings.Join(p, ", "))

	c := helper.Cmd{
		Bin:  "paru",
		Args: []string{"-Sy"},
	}
	c.Args = append(c.Args, p...)

	err := helper.ExecuteCommand(c)
	if err != nil {
		f.Log.Error("Install AUR packages", err.Error())
		return err
	}

	for _, pk := range p {
		i := f.checkInstalledPackage(pk)
		if !i {
			f.Log.Error("Package " + pk + " failed to install. Aborting setup.")
			return errors.New("")
		}
	}
	return nil
}

func (f *ArchHelper) checkInstalledPackage(p string) bool {
	c := helper.Cmd{
		Bin:  "pacman",
		Args: []string{"-Q", p},
	}

	err := helper.ExecuteCommand(c)
	if err != nil {
		return false
	}

	return true
}
