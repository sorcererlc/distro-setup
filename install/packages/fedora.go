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
		Log:  log.NewLog("fedora_packages.log"),
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

	return nil
}

func (f *FedoraHelper) updateDistro() error {
	f.Log.Info("Updating packages")

	c := helper.Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "upgrade", "-y"},
	}

	err := helper.ExecuteCommand(c)
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

	c := helper.Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "install", "-y"},
	}
	c.Args = append(c.Args, r...)

	err := helper.ExecuteCommand(c)
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

	c := helper.Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "copr", "enable", "-y"},
	}
	c.Args = append(c.Args, f.CoprRepos.Copr...)

	err = helper.ExecuteCommand(c)
	if err != nil {
		f.Log.Error("Enable Copr repositories", err.Error())
		return err
	}

	return nil
}

func (f *FedoraHelper) removePackages(p []string) error {
	f.Log.Info("Removing packages", strings.Join(p, ", "))

	c := helper.Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "remove", "-y"},
	}
	c.Args = append(c.Args, p...)

	err := helper.ExecuteCommand(c)
	if err != nil {
		f.Log.Error("Remove packages", err.Error())
		return err
	}

	return nil
}

func (f *FedoraHelper) installPackages(p []string) error {
	f.Log.Info("Installing packages", strings.Join(p, ", "))

	c := helper.Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "install", "-y"},
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

func (f *FedoraHelper) checkInstalledPackage(p string) bool {
	c := helper.Cmd{
		Bin:  "dnf",
		Args: []string{"list", "installed", p},
	}

	err := helper.ExecuteCommand(c)
	if err != nil {
		return false
	}

	return true
}
