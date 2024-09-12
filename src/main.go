package main

import (
	"bufio"
	"fmt"
	"os"
	"setup/helper"
	"setup/helper/distro"
	"setup/log"
	"setup/packages"
	"strings"
)

var Reset = "\033[0m"
var Green = "\033[32m"

func menuItem(i string, s string) {
	fmt.Printf("[%s%s%s] %s\n", Green, i, Reset, s)
}

func menu() string {
	menuItem("1", "Update distro and initial setup. Run this first!")
	menuItem("2", "Install base packages")
	menuItem("3", "Install extra packages")
	menuItem("4", "Install NVIDIA driver")
	menuItem("5", "Install SDDM")
	menuItem("6", "Install Hyprland")
	menuItem("7", "Install bluetooth")
	menuItem("8", "Install flatpak packages")
	menuItem("9", "Setup shell")
	menuItem("r", "Reboot")
	menuItem("q", "Quit")

	in, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	in = strings.ReplaceAll(in, "\n", "")

	return in
}

func fail(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	log.ClearLogs()
	l := log.NewStdOutLog()

	env, err := helper.GetEnvironment()
	fail(err)

	conf, err := helper.GetConfig(env)
	fail(err)

	for {
		helper.ClearScreen()

		p, err := packages.NewPkg(conf, env)

		c := menu()
		switch c {
		case "1":
			err = p.SetupDistro()
			fail(err)
			break
		case "2":
			err = p.InstallBasePackages()
			fail(err)
			break
		case "3":
			err = p.InstallExtraPackages()
			fail(err)
			break
		case "4":
			err = p.InstallNvidia()
			fail(err)
			break
		case "5":
			err = p.InstallSddm()
			fail(err)
			break
		case "6":
			err = p.InstallHyprland()
			fail(err)
			break
		case "7":
			err = p.InstallBluetooth()
			fail(err)
			break
		case "8":
			fp := packages.NewFlatpakHelper(conf, env)
			err = fp.InstallPackages()
			fail(err)
			break
		case "9":
			dh, err := distro.NewDistroHelper(conf, env)
			fail(err)

			helper.ClearScreen()
			l.Info("Preparing to setup shell")

			err = dh.SetupDistro()
			fail(err)
			break
		case "r":
			err = helper.Run("sudo", "reboot")
			fail(err)
			break
		case "q", "Q":
			println("Bye")
			os.Exit(0)
		}
	}
}
