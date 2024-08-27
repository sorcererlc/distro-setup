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

	env, err := helper.GetEnvironment()
	if err != nil {
		os.Exit(1)
	}

	conf, err := helper.GetConfig(env)
	if err != nil {
		os.Exit(1)
	}

	p := packages.NewPkg(conf, env)
	err = p.SetupPackages()
	if err != nil {
		os.Exit(1)
	}

	fp := packages.NewFlatpakHelper(conf, env)
	err = fp.InstallRepos()
	if err != nil {
		os.Exit(1)
	}
	err = fp.InstallPackages()
	if err != nil {
		os.Exit(1)
	}

	dh, err := distro.NewDistroHelper(conf, env)
	if err != nil {
		os.Exit(1)
	}
	err = dh.SetupDistro()
	if err != nil {
		os.Exit(1)
	}

	print("\n\n")
	print("Installation complete. You must reboot the machine to finish setup.\nDo you want to reboot now? (Y/n) ")

	in, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	in = strings.ReplaceAll(in, "\n", "")

	if slices.Contains([]string{"Y", "y", ""}, in) {
		println("Rebooting now")

		if os.Getenv("TEST") != "true" {
			err = helper.Run("sudo", "reboot")
			if err != nil {
				panic(err)
			}
		}
	}
}
