package main

import (
	"bufio"
	"os"
	"setup/helper"
	"setup/helper/distro"
	"setup/log"
	"setup/packages"
	"slices"
	"strings"
)

func main() {
	log.ClearLogs()
	l := log.NewStdOutLog()

	env, err := helper.GetEnvironment()
	if err != nil {
		os.Exit(1)
	}

	conf, err := helper.GetConfig(env)
	if err != nil {
		os.Exit(1)
	}

	helper.ClearScreen()
	l.Info("Preparing to install required packages")

	p := packages.NewPkg(conf, env)
	err = p.SetupPackages()
	if err != nil {
		os.Exit(1)
	}

	helper.ClearScreen()
	l.Info("Preparing to install flatpak packages")

	fp := packages.NewFlatpakHelper(conf, env)
	err = fp.InstallPackages()
	if err != nil {
		os.Exit(1)
	}

	dh, err := distro.NewDistroHelper(conf, env)
	if err != nil {
		os.Exit(1)
	}

	helper.ClearScreen()
	l.Info("Preparing to setup shell")

	err = dh.SetupDistro()
	if err != nil {
		os.Exit(1)
	}

	if os.Getenv("TEST") == "true" {
		os.Exit(0)
	}

	print("\n\n")
	print("Installation complete. You must reboot the machine to finish setup.\nDo you want to reboot now? (Y/n) ")

	in, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	in = strings.ReplaceAll(in, "\n", "")

	if slices.Contains([]string{"Y", "y", ""}, in) {
		println("Rebooting now")

		err = helper.Run("sudo", "reboot")
		if err != nil {
			panic(err)
		}
	}
}
