package distro

import "setup/helper"

func (f *DistroHelper) setupShell() error {
	f.Log.Info("Setting up shell")

	err := helper.Run("chsh", "-s", "$(which zsh)")
	if err != nil {
		f.Log.Error("Change shell to zsh", err.Error())
		return err
	}

	err = helper.Run("ln", "-s", f.Env.Cwd+"/config/home/zshrc", "$HOME/.zshrc")
	if err != nil {
		f.Log.Error("Link .zshrc", err.Error())
		return err
	}

	err = helper.Run("curl", "https://ohmyposh.dev/install.sh", "--create-dirs", "-o", "./ohmyposh/install.sh", "&&", "sh", "./ohmyposh/install.sh")
	if err != nil {
		f.Log.Error("Install OhMyPosh", err.Error())
		return err
	}

	err = helper.Run("rm", "-rf", "ohmyposh")
	if err != nil {
		f.Log.Error("Cleanup OhMyPosh", err.Error())
		return err
	}

	return nil
}
