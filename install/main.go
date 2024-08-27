package main

import (
	"os"
	"setup/helper"
	"setup/log"
	"setup/packages"
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
	fp.InstallRepos()
	fp.InstallPackages()
}
