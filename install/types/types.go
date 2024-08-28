package types

type Environment struct {
	OS   OS
	Cwd  string
	User struct {
		Username string
		Gid      string
	}
}

type OS struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PrettyName string `json:"pretty_name"`
	Version    string `json:"version"`
	VersionId  string `json:"version_id"`
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
	DotFilesRepo GitPackage `yaml:"dotfiles_repo"`
}

type GitPackage struct {
	Url string `yaml:"url"`
	Tag string `yaml:"tag"`
}

type Packages struct {
	Repo      []string              `yaml:"repo"`
	Base      []string              `yaml:"base"`
	Hyprland  []string              `yaml:"hyprland"`
	Sway      []string              `yaml:"sway"`
	Nvidia    []string              `yaml:"nvidia"`
	Sddm      []string              `yaml:"sddm"`
	Bluetooth []string              `yaml:"bluetooth"`
	Extras    []string              `yaml:"extras"`
	Aur       []string              `yaml:"aur"`
	AurExtra  []string              `yaml:"aur_extras"`
	Remove    []string              `yaml:"remove"`
	Fonts     []string              `yaml:"fonts"`
	Git       map[string]GitPackage `yaml:"git"`
}

type DistroConfig struct {
	Services struct {
		Bluetooth []string `yaml:"bluetooth"`
		Sddm      []string `yaml:"sddm"`
	} `yaml:"services"`
	Groups []string `yaml:"groups"`
	Udev   []struct {
		Name string `yaml:"name"`
		Rule string `yaml:"rule"`
		File string `yaml:"file"`
	} `yaml:"udev"`
}
