package helper

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Options struct {
		WindowManager  string `yaml:"window_manager"`
		GlobalMangoHud bool   `yaml:"global_mango_hud"`
	} `yaml:"options"`
	Packages struct {
		Base      bool `yaml:"base"`
		Nvidia    bool `yaml:"nvidia"`
		Sddm      bool `yaml:"sddm"`
		Bluetooth bool `yaml:"bluetooth"`
		Extras    bool `yaml:"extras"`
		Dotfiles  bool `yaml:"dotfiles"`
	} `yaml:"main"`
	Flatpak struct {
		Packages struct {
			Base   bool `yaml:"base"`
			Devel  bool `yaml:"devel"`
			Extras bool `yaml:"extras"`
			Misc   bool `yaml:"misc"`
		} `yaml:"packages"`
	} `yaml:"flatpak"`
}

func LoadCongig() (*Config, error) {
	conf := Config{}

	c, err := os.Getwd()
	if err != nil {
		println("Error reading CWD")
		return nil, err
	}

	fs := c + "/config.yml"
	println("Loading config file " + c)

	f, err := os.ReadFile(fs)
	if err != nil {
		println("Error loading config file" + err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		println("Error decoding config file" + err.Error())
		return nil, err
	}

	return &conf, nil
}

func SetupDistro(id string) error {
	switch id {
	case "fedora":
		break
	case "arch":
		break
	}

	return nil
}

func SetupDotfiles(id string) error {

	return nil
}
