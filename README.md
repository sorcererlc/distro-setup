# Linux distro setup utility

### Purpose

This tool is meant to be used on a minimal Linux to install desired packages and perform initial configuration automatically. It was born out of my need to replicate a desktop setup on multiple machines.

The included YAML files will set up Hyprland or Sway as a window manager with Waybar and some other utilities to make it a fully functional lightweight desktop system.

### Important note

As of now this tool assumes that it runs on a fresh Linux install and will not check for existing packages and configuration. This will most likely lead to broken configuration files if run on an already setup system! This might change in the future but is not a priority right now.

### Why go instead of bash scripts?

While the result is the same and, indeed, everything this tool does is done via shell commands, I find working with configuration files and performing more complex programming tasks easier with a proper programming language compared to bash scripts. Also, I'm a big proponent of compiled vs. interpreted languages in general :)

### Installation

```
git clone --depth 1 https://github.com/sorcererlc/distro-setup.git
cd distro-setup
make install
```

If you would like to see all the commands the tool will run without making any changes to your system you can run
```
make test
```
