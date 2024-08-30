package packages

import (
	"errors"
	"fmt"
	"os"
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
		Log:  log.NewLog("packages.log"),
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

	err = f.installParu()
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

	if f.Conf.Options.WindowManager == "hyprland" {
		p = append(p, pkg.Hyprland...)
	}

	if f.Conf.Options.WindowManager == "sway" {
		p = append(p, pkg.Sway...)
	}

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

	err := helper.Run("sudo", "pacman", "-Syu")
	if err != nil {
		f.Log.Error("Update packages", err.Error())
		return err
	}

	return nil
}

func (f *ArchHelper) installRepos(r []string) error {
	if len(r) == 0 {
		return nil
	}

	for i := 0; i < len(r); i++ {
		r[i] = fmt.Sprintf(r[i], f.Env.OS.VersionId)
	}
	f.Log.Info("Installing repositories", strings.Join(r, ", "))

	args := []string{"sudo", "pacman", "-Sy", "--needed", "--noconfirm"}
	args = append(args, r...)
	err := helper.Run(args...)
	if err != nil {
		f.Log.Error("Install repositories", err.Error())
		return err
	}

	return nil
}

func (f *ArchHelper) removePackages(p []string) error {
	if len(p) == 0 {
		return nil
	}

	f.Log.Info("Removing packages", strings.Join(p, ", "))

	args := []string{"sudo", "pacman", "-Rs"}
	args = append(args, p...)
	err := helper.Run(args...)
	if err != nil {
		f.Log.Error("Remove packages", err.Error())
		return err
	}

	return nil
}

func (f *ArchHelper) installPackages(p []string) error {
	f.Log.Info("Installing packages", strings.Join(p, ", "))

	args := []string{"sudo", "pacman", "-Sy", "--needed", "--noconfirm"}
	args = append(args, p...)
	err := helper.Run(args...)
	if err != nil {
		f.Log.Error("Install packages", err.Error())
		return err
	}

	// for _, pk := range p {
	// 	i := f.checkInstalledPackage(pk)
	// 	if !i {
	// 		f.Log.Error("Package " + pk + " failed to install. Aborting setup.")
	// 		return errors.New("")
	// 	}
	// }

	return nil
}

func (f *ArchHelper) installAurPackages(p []string) error {
	f.Log.Info("Installing AUR packages", strings.Join(p, ", "))

	args := []string{"paru", "-Sy"}
	args = append(args, p...)
	err := helper.Run(args...)
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
	err := helper.Run("pacman", "-Q", p)
	if err != nil {
		return false
	}

	return true
}

func (f *ArchHelper) installParu() error {
	_, err := os.Stat("/usr/bin/paru")
	if err == nil {
		f.Log.Info("Paru is already installed. Moving on.")
		return nil
	}

	f.Log.Info("Cloning paru repo")

	err = helper.Run("git", "clone", "https://aur.archlinux.org/paru.git")
	if err != nil {
		f.Log.Error("Paru repo clone", err.Error())
		return err
	}

	err = helper.Run("./scripts/paru.sh")
	if err != nil {
		f.Log.Error("Install paru", err.Error())
		return err
	}

	return nil
}
