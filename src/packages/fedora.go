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

	err = f.installAdvCpMv()
	if err != nil {
		return err
	}

	err = f.installAutoCpuFreq(pkg.Git["auto-cpufreq"])
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

	for _, r := range f.CoprRepos.Copr {
		err = helper.Run("sudo", "dnf", "copr", "enable", "-y", r)
		if err != nil {
			f.Log.Error("Enable Copr repositories", err.Error())
			return err
		}
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

	for _, pk := range p {
		err := helper.Run("sudo", "dnf", "install", "-y", pk)
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

func (f *FedoraHelper) checkInstalledPackage(p string) bool {
	err := helper.Run("dnf", "list", "installed", p)
	if err != nil {
		return false
	}

	return true
}

func (f *FedoraHelper) runGitCommand(c string) error {
	args := strings.Split(c, " ")

	return helper.Run(args...)
}

func (f *FedoraHelper) setupNwgLook(p types.GitPackage) error {
	helper.ClearScreen()
	f.Log.Info("Cloning nwg-look repo")

	_, cpe := os.Stat("/usr/bin/nwg-look")
	if cpe == nil {
		f.Log.Info("nwg-look is already installed, skipping")
		return nil
	}

	_, re := os.Stat("nwg-look")
	if re == nil {
		err := helper.Run("rm", "-rf", "nwg-look")
		if err != nil {
			f.Log.Error("Remove nwg-look repo", err.Error())
			return err
		}
	}

	err := helper.Run("git", "clone", "--recursive", "--depth", "1", "--branch", p.Tag, p.Url)
	if err != nil {
		f.Log.Error("Clone nwg-look repo", err.Error())
		return err
	}

	helper.ClearScreen()
	f.Log.Info("Installing nwg-look")

	for _, c := range p.Commands {
		err := f.runGitCommand(c)
		if err != nil {
			f.Log.Error(c, err.Error())
			return err
		}
	}

	return nil
}

func (f *FedoraHelper) installAutoCpuFreq(p types.GitPackage) error {
	helper.ClearScreen()
	f.Log.Info("Cloning auto-cpufreq repo")

	_, cpe := os.Stat("/usr/local/bin/auto-cpufreq")
	if cpe == nil {
		f.Log.Info("auto-cpufreq is already installed, skipping")
		return nil
	}

	_, re := os.Stat("auto-cpufreq")
	if re == nil {
		err := helper.Run("rm", "-rf", "auto-cpufreq")
		if err != nil {
			f.Log.Error("Remove auto-cpufreq repo", err.Error())
			return err
		}
	}

	err := helper.Run("git", "clone", "--recursive", "--depth", "1", p.Url)
	if err != nil {
		f.Log.Error("Clone auto-cpufreq repo", err.Error())
		return err
	}

	f.Log.Info("Installing auto-cpufreq")

	for _, c := range p.Commands {
		err := f.runGitCommand(c)
		if err != nil {
			f.Log.Error(c, err.Error())
			return err
		}
	}

	return nil
}

func (f *FedoraHelper) installAdvCpMv() error {
	helper.ClearScreen()
	f.Log.Info("Installing advcpmv")

	_, cpe := os.Stat("/usr/local/bin/cpg")
	_, mve := os.Stat("/usr/local/bin/mvg")
	if cpe == nil && mve == nil {
		f.Log.Info("advcpmv is already installed, skipping")
		return nil
	}

	err := helper.Run("curl", "-O", "https://raw.githubusercontent.com/jarun/advcpmv/master/install.sh")
	if err != nil {
		f.Log.Error("Remove advcpmv", err.Error())
		return err
	}

	err = helper.Run("sh", "install.sh")
	if err != nil {
		f.Log.Error("Build advcpmv", err.Error())
		return err
	}

	helper.ClearScreen()
	f.Log.Info("Finished building advcpmv. Copying binaries to /usr/local/bin...")

	err = helper.Run("sudo", "mv", "advcp", "/usr/local/bin/cpg")
	err = helper.Run("sudo", "mv", "advmv", "/usr/local/bin/mvg")
	if err != nil {
		f.Log.Error("Install advcpmv binaries", err.Error())
		return err
	}

	err = helper.Run("rm", "-f", "./install.sh")
	if err != nil {
		f.Log.Error("Cleanup advcpmv", err.Error())
		return err
	}

	return nil
}
