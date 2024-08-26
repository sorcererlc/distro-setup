package main

import (
	"setup/helper"
)

func main() {
	conf, err := helper.LoadCongig()
	if err != nil {
		panic(err)
	}

	os, err := helper.GetOS()
	if err != nil {
		panic(err)
	}

	pkg, distPkg, err := helper.LoadPackages(os.Id)
	if err != nil {
		panic(err)
	}

	// Install extra repositories
	err = helper.InstallRepos(os.Id, os.VersionId, distPkg.Repo)
	if err != nil {
		panic(err)
	}

	// Install base packages
	err = helper.InstallPackages(os.Id, append(pkg.Base, distPkg.Base...))
	if err != nil {
		panic(err)
	}

	// Install and configure Nvidia packages if enabled
	if conf.Packages.Nvidia {
		err = helper.InstallPackages(os.Id, distPkg.Nvidia)
		if err != nil {
			panic(err)
		}
	}

	// Install window manager packages
	println("Installing " + conf.Options.WindowManager + " window manager")
	switch conf.Options.WindowManager {
	case "hyprland":
		err = helper.InstallPackages(os.Id, append(pkg.Hyprland, distPkg.Hyprland...))
		break
	case "sway":
		err = helper.InstallPackages(os.Id, append(pkg.Sway, distPkg.Hyprland...))
		break
	default:
		panic("Packages for specified window manager are not present in config files")
	}
	if err != nil {
		panic(err)
	}

	// Install bluetooth packages
	if conf.Packages.Bluetooth {
		err = helper.InstallPackages(os.Id, append(pkg.Bluetooth, distPkg.Bluetooth...))
		if err != nil {
			panic(err)
		}
	}

	// Install SDDM
	if conf.Packages.Sddm {
		err = helper.InstallPackages(os.Id, append(pkg.Sddm, distPkg.Sddm...))
		if err != nil {
			panic(err)
		}
	}

	// Install extras
	if conf.Packages.Extras {
		err = helper.InstallPackages(os.Id, append(pkg.Extras, distPkg.Extras...))
		if err != nil {
			panic(err)
		}
	}

	fp, err := helper.NewFlatpakHelper(conf)
	if err != nil {
		panic(err)
	}

	fp.InstallRepos()
	fp.InstallPackages()

	// Configure various distro settings (theming, fonts, symlinks, etc)
	err = helper.SetupDistro(os.Id)
	if err != nil {
		panic(err)
	}
}
