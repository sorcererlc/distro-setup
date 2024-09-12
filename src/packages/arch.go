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
	Packages  *types.Packages
	Log       *log.Log
	CoprRepos struct {
		Copr []string `yaml:"copr"`
	}
}

func NewArchHelper(c *types.Config, e *types.Environment, p *types.Packages) *ArchHelper {
	f := ArchHelper{
		Conf:     c,
		Env:      e,
		Packages: p,
		Log:      log.NewLog("packages.log"),
	}

	return &f
}

func (f *ArchHelper) SetupDistro() error {
	err := f.updateDistro()
	if err != nil {
		return err
	}

	err = f.removePackages(f.Packages.Remove)
	if err != nil {
		return err
	}

	err = f.installRepos(f.Packages.Repo)
	if err != nil {
		return err
	}

	err = f.installParu()
	if err != nil {
		return err
	}

	err = f.setupPacman()
	if err != nil {
		return err
	}

	return nil
}

func (f *ArchHelper) InstallBasePackages() error {
	err := f.installPackages(f.Packages.Base)
	if err != nil {
		return err
	}

	err = f.installPackages(f.Packages.Fonts)
	if err != nil {
		return err
	}

	err = f.installAurPackages(f.Packages.Aur)
	if err != nil {
		return err
	}

	return nil
}

func (f *ArchHelper) InstallExtraPackages() error {
	err := f.installPackages(f.Packages.Extras)
	if err != nil {
		return err
	}

	err = f.installAurPackages(f.Packages.AurExtra)
	if err != nil {
		return err
	}

	return nil
}

func (f *ArchHelper) InstallNvidia() error {
	err := f.installPackages(f.Packages.Nvidia)
	if err != nil {
		return err
	}

	err = f.setupNvidia()
	if err != nil {
		return err
	}

	return nil
}

func (f *ArchHelper) InstallSddm() error {
	err := f.installPackages(f.Packages.Sddm)
	if err != nil {
		return err
	}

	return nil
}

func (f *ArchHelper) InstallHyprland() error {
	err := f.installPackages(f.Packages.Hyprland)
	if err != nil {
		return err
	}

	return nil
}

func (f *ArchHelper) InstallBluetooth() error {
	err := f.installPackages(f.Packages.Bluetooth)
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

	for _, pk := range p {
		err := helper.Run("sudo", "pacman", "-Sy", "--needed", "--noconfirm", pk)
		if err != nil {
			f.Log.Error("Install packages", err.Error())
			return err
		}

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

	for _, pk := range p {
		args := []string{"paru", "-Sy", pk}
		err := helper.Run(args...)
		if err != nil {
			f.Log.Error("Install AUR packages", err.Error())
			return err
		}

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

func (f *ArchHelper) setupNvidia() error {
	f.Log.Info("Setting up NVIDIA driver")

	err := helper.Run("./scripts/arch-nvidia.sh")
	if err != nil {
		f.Log.Error("Setup NVIDIA driver", err.Error())
		return err
	}

	return nil
}

func (f *ArchHelper) setupPacman() error {
	f.Log.Info("Setting up pacman")

	err := helper.Run("./scripts/pacman.sh")
	if err != nil {
		f.Log.Error("Setup pacman", err.Error())
		return err
	}

	return nil
}
