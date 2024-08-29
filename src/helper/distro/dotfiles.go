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
	args = append(args, p.Url, "../dotfiles")

	err := helper.Run(args...)
	if err != nil {
		f.Log.Error("Clone dotfiles repo", err.Error())
		return err
	}

	err = helper.Run("./scripts/dotfiles.sh")
	if err != nil {
		f.Log.Error("Setup dotfiles", err.Error())
		return err
	}

	return nil
}
