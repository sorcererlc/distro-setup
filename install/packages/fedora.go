package packages

import (
	"errors"
	"fmt"
	"os"
	"setup/helper"
	"setup/log"
	"setup/types"
	"strings"

	"gopkg.in/yaml.v3"
)

type FedoraHelper struct {
	Conf      *types.Config
	Env       *types.Environment
	Log       *log.Log
	CoprRepos struct {
		Copr []string `yaml:"copr"`
	}
}

func NewFedoraHelper(c *types.Config, e *types.Environment) *FedoraHelper {
	f := FedoraHelper{
		Conf: c,
		Env:  e,
		Log:  log.NewLog("packages.log"),
	}

	return &f
}

func (f *FedoraHelper) SetupPackages(pkg *types.Packages) error {
	err := f.updateDistro()
	if err != nil {
		return err
	}

	err = f.enableCoprRepos()
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

	err = f.setupNwgLook(pkg.Git["nwg-look"])
	if err != nil {
		return err
	}

	return nil
}

func (f *FedoraHelper) updateDistro() error {
	f.Log.Info("Updating packages")

	err := helper.Run("sudo", "dnf", "upgrade", "-y")
	if err != nil {
		f.Log.Error("Update packages", err.Error())
		return err
	}

	return nil
}

func (f *FedoraHelper) installRepos(r []string) error {
	for i := 0; i < len(r); i++ {
		r[i] = fmt.Sprintf(r[i], f.Env.OS.VersionId)
	}
	f.Log.Info("Installing repositories", strings.Join(r, ", "))

	args := []string{"sudo", "dnf", "install", "-y"}
	args = append(args, r...)
	err := helper.Run(args...)
	if err != nil {
		f.Log.Error("Install repositories", err.Error())
		return err
	}

	return nil
}

func (f *FedoraHelper) enableCoprRepos() error {
	f.Log.Info("Enabling Copr repositories", strings.Join(f.CoprRepos.Copr, ", "))

	fs, err := os.ReadFile(f.Env.Cwd + "/packages/fedora/repos.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fs, &f.CoprRepos)

	args := []string{"sudo", "dnf", "copr", "enable", "-y"}
	args = append(args, f.CoprRepos.Copr...)
	err = helper.Run(args...)
	if err != nil {
		f.Log.Error("Enable Copr repositories", err.Error())
		return err
	}

	return nil
}

func (f *FedoraHelper) removePackages(p []string) error {
	f.Log.Info("Removing packages", strings.Join(p, ", "))

	args := []string{"sudo", "dnf", "remove", "-y"}
	args = append(args, p...)
	err := helper.Run(args...)
	if err != nil {
		f.Log.Error("Remove packages", err.Error())
		return err
	}

	return nil
}

func (f *FedoraHelper) installPackages(p []string) error {
	f.Log.Info("Installing packages", strings.Join(p, ", "))

	args := []string{"sudo", "dnf", "install", "-y"}
	args = append(args, p...)
	err := helper.Run(args...)
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

func (f *FedoraHelper) checkInstalledPackage(p string) bool {
	err := helper.Run("dnf", "list", "installed", p)
	if err != nil {
		return false
	}

	return true
}

func (f *FedoraHelper) setupNwgLook(p types.GitPackage) error {
	err := helper.Run("git", "clone", "--reursive", "--depth", "1", "--branch", p.Tag, p.Url)
	if err != nil {
		f.Log.Error("Clone nwg-look repo", err.Error())
		return err
	}

	_ = helper.Run("cd", "nwg-look")

	err = helper.Run("make", "build")
	if err != nil {
		f.Log.Error("Build nwg-look", err.Error())
		return err
	}

	err = helper.Run("sudo", "make", "install")
	if err != nil {
		f.Log.Error("Install nwg-look", err.Error())
		return err
	}

	_ = helper.Run("cd", f.Env.Cwd)
	_ = helper.Run("rm", "-rf", "nwg-look")

	return nil
}
