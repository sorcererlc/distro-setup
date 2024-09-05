package distro

import (
	"os"
	"setup/helper"
	"setup/log"
	"setup/types"
	"strings"

	"gopkg.in/yaml.v3"
)

type DistroHelper struct {
	Conf         *types.Config
	Env          *types.Environment
	DistroConfig *types.DistroConfig
	Log          *log.Log
}

func NewDistroHelper(c *types.Config, e *types.Environment) (*DistroHelper, error) {
	f := DistroHelper{
		Conf: c,
		Env:  e,
		Log:  log.NewLog("distro-setup.log"),
	}

	err := f.LoadConfig()
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func (f *DistroHelper) LoadConfig() error {
	c := types.DistroConfig{}

	fs, err := os.ReadFile("./distro_config/distro.yml")
	if err != nil {
		f.Log.Error("Load distro config file", err.Error())
		return err
	}

	err = yaml.Unmarshal(fs, &c)

	f.DistroConfig = &c

	fs, err = os.ReadFile("./")

	return nil
}

func (f *DistroHelper) SetupDistro() error {
	for _, g := range f.DistroConfig.Groups {
		err := f.createGroup(g)
		if err != nil {
			return err
		}

		err = f.addUserToGroup(g, f.Env.User.Username)
		if err != nil {
			return err
		}
	}

	if f.Conf.Packages.Bluetooth {
		for _, s := range f.DistroConfig.Services.Bluetooth {
			err := f.enableService(s)
			if err != nil {
				return err
			}
		}
	}

	if f.Conf.Packages.Sddm {
		for _, s := range f.DistroConfig.Services.Sddm {
			err := f.enableService(s)
			if err != nil {
				return err
			}
		}

		err := helper.Run("sudo", "systemctl", "set-default", "graphical.target")
		if err != nil {
			f.Log.Error("Set default graphical target", err.Error())
			return err
		}
	}

	err := f.detectSensors()
	if err != nil {
		return err
	}

	if f.Conf.Options.Firewall {
		err = f.setupFirewall()
		if err != nil {
			return err
		}
	}

	for _, r := range f.DistroConfig.Udev {
		err := f.setupUdevRule(r.Name, r.Rule, r.File)
		if err != nil {
			return err
		}
	}

	err = f.reloadUdevRules()
	if err != nil {
		return err
	}

	if f.Conf.Options.NetworkShares {
		err = f.setupNetworkShares()
		if err != nil {
			return err
		}
	}

	err = f.setupShell()
	if err != nil {
		return err
	}

	if f.Conf.Packages.Dotfiles {
		err := f.setupDotfiles(f.Conf.DotFilesRepo)
		if err != nil {
			return err
		}
	}

	if f.Conf.Packages.NVim {
		err := f.setupNeoVim(f.Conf.NVimRepo)
		if err != nil {
			return err
		}
	}

	if f.Conf.Options.AutoLogin {
		err := f.setupAutoLogin()
		if err != nil {
			return err
		}

		if !f.Conf.Packages.Sddm {
			err := f.setupWindowManagerAutostart()
			if err != nil {
				return err
			}
		}
	}

	if len(f.DistroConfig.Shell.Base) > 0 {
		err := f.runShellCommands(f.DistroConfig.Shell.Base)
		if err != nil {
			return err
		}
	}

	if f.Conf.Options.WindowManager == "hyprland" {
		err := f.runShellCommands(f.DistroConfig.Shell.Hyprland)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *DistroHelper) writeFile(fs string, s string, a bool, su bool) error {
	args := []string{"echo", s}

	if su {
		args = append(args, "|", "sudo", "tee")
		if a {
			args = append(args, "-a")
		}
	} else {
		if a {
			args = append(args, ">>")
		} else {
			args = append(args, ">")
		}
	}

	args = append(args, fs)

	err := helper.Run(args...)
	if err != nil {
		return err
	}

	return nil
}

func (f *DistroHelper) setupAutoLogin() error {
	if f.Conf.Packages.Sddm {
		f.Log.Info("Setting up auto SDDM login")

		err := helper.Run("sudo", "mkdir", "-p", "/etc/sddm.conf.d")
		if err != nil {
			f.Log.Error("Create /etc/sddm.conf.d", err.Error())
			return err
		}

		fs := "[Autologin]\nRelogin=false\nUser=" + f.Env.User.Username + "\nSession=" + f.Conf.Options.WindowManager

		err = os.WriteFile("autologin.conf", []byte(fs), 0644)
		if err != nil {
			f.Log.Error("Write autologin file", err.Error())
			return err
		}

		err = helper.Run("sudo", "mv", "autologin.conf", "/etc/sddm.conf.d/")
		if err != nil {
			f.Log.Error("Move autologin.conf to /etc/sddm.conf.d/", err.Error())
			return err
		}

		return nil
	}

	f.Log.Info("Setting up TTY autologin")

	err := helper.Run("./scripts/autologin.sh")
	if err != nil {
		f.Log.Error("Write autologin file", err.Error())
		return err
	}

	return nil
}

func (f *DistroHelper) setupWindowManagerAutostart() error {
	f.Log.Info("Setting up " + f.Conf.Options.WindowManager + " autostart")

	fs := "if [ -z \"$WAYLAND_DISPLAY\" ] && [ \"$XDG_VTNR\" -eq 1 ]; then\n  exec " + f.Conf.Options.WindowManager + "\nfi"

	err := os.WriteFile(f.Env.User.HomeDir+"/.zprofile", []byte(fs), 0744)
	if err != nil {
		f.Log.Error("Write .zprofile", err.Error())
		return err
	}

	return nil
}

func (f *DistroHelper) runShellCommands(cs []string) error {
	for _, c := range cs {
		args := strings.Split(c, " ")
		err := helper.Run(args...)
		if err != nil {
			f.Log.Error("Run command", c, err.Error())
			return err
		}
	}

	return nil
}
