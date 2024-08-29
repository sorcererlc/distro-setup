package distro

import (
	"setup/helper"
	"setup/types"
)

func (f *DistroHelper) setupDotfiles(p types.GitPackage) error {
	f.Log.Info("Setting up dotfiles")

	args := []string{"git", "clone", "--depth", "1"}
	if p.Tag != "" {
		args = append(args, "--branch", p.Tag)
	}
	args = append(args, p.Url)

	err := helper.Run(args...)
	if err != nil {
		f.Log.Error("Clone dotfiles repo", err.Error())
		return err
	}

	_ = helper.Run("cd", "dotfiles")

	err = helper.Run("make", "install")
	if err != nil {
		f.Log.Error("Setup dotfiles", err.Error())
		return err
	}

	return nil
}
