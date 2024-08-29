package distro

import "setup/helper"

func (f *DistroHelper) setupShell() error {
	_ = helper.Run("clear")

	f.Log.Info("Setting up shell.")

	err := helper.Run("chsh", "-s", "/usr/bin/zsh")
	if err != nil {
		f.Log.Error("Change shell to zsh", err.Error())
		return err
	}

	err = helper.Run("rm", "-f", "$HOME/.zshrc")
	if err != nil {
		f.Log.Error("Remove ~/.zshrc. Please remove the file manually and try again.", err.Error())
		return err
	}

	err = helper.Run("ln", "-s", f.Env.Cwd+"/config/home/zshrc", "$HOME/.zshrc")
	if err != nil {
		f.Log.Error("Link .zshrc", err.Error())
		return err
	}

	err = helper.Run("curl", "-O", "https://ohmyposh.dev/install.sh")
	if err != nil {
		f.Log.Error("Download OhMyPosh install scripts", err.Error())
		return err
	}

	err = helper.Run("sh", "install.sh")
	if err != nil {
		f.Log.Error("Install OhMyPosh", err.Error())
		return err
	}

	err = helper.Run("rm", "-f", "install.sh")
	if err != nil {
		f.Log.Error("Cleanup OhMyPosh", err.Error())
		return err
	}

	return nil
}
