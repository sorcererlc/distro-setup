package types

type Environment struct {
	OS  OS
	Cwd string
}

type OS struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PrettyName string `json:"pretty_name"`
	Version    string `json:"version"`
	VersionId  string `json:"version_id"`
}

type Packages struct {
	Repo      []string `yaml:"repo"`
	Base      []string `yaml:"base"`
	Hyprland  []string `yaml:"hyprland"`
	Sway      []string `yaml:"sway"`
	Nvidia    []string `yaml:"nvidia"`
	Sddm      []string `yaml:"sddm"`
	Bluetooth []string `yaml:"bluetooth"`
	Extras    []string `yaml:"extras"`
	Aur       []string `yaml:"aur,omitempty"`
	AurExtra  []string `yaml:"aur_extras,omitempty"`
	Remove    []string `yaml:"remove"`
	Fonts     []string `yaml:"fonts"`
}

type Config struct {
	Options struct {
		WindowManager  string `yaml:"window_manager"`
		GlobalMangoHud bool   `yaml:"global_mango_hud"`
	} `yaml:"options"`
	Packages struct {
		Nvidia    bool `yaml:"nvidia"`
		Sddm      bool `yaml:"sddm"`
		Bluetooth bool `yaml:"bluetooth"`
		Extras    bool `yaml:"extras"`
		Dotfiles  bool `yaml:"dotfiles"`
	} `yaml:"packages"`
	Flatpak struct {
		Packages struct {
			Base   bool `yaml:"base"`
			Devel  bool `yaml:"devel"`
			Extras bool `yaml:"extras"`
			Misc   bool `yaml:"misc"`
		} `yaml:"packages"`
	} `yaml:"flatpak"`
}
