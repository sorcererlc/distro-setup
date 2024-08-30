package distro

import (
	"setup/helper"
	"setup/types"
)

func (f *DistroHelper) setupNeoVim(p types.GitPackage) error {
	f.Log.Info("Setting up NeoVim")

	err := helper.Run("rm", "-rf", f.Env.User.HomeDir+"/.config/nvim")
	if err != nil {
		f.Log.Error("Remove existing NeoVim config directory", err.Error())
		return err
	}

	args := []string{"git", "clone", "--depth", "1"}
	if p.Tag != "" {
		args = append(args, "--branch", p.Tag)
	}
	args = append(args, p.Url, f.Env.User.HomeDir+"/.config/nvim")

	err = helper.Run(args...)
	if err != nil {
		f.Log.Error("Clone NeoVim config repo", err.Error())
		return err
	}

	return nil
}
