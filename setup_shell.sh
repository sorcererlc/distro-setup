#!/usr/bin/env bash

setup_fedora()
{
    sudo dnf install -y \
        zsh \
        bat \
        tldr \
        progress \
        htop \
        pipx \
        unzip
}

setup_arch()
{
    sudo pacman -Syu
    sudo pacman -S \
        zsh \
        bat \
        tldr \
        progress \
        htop \
        unzip
}

setup_debian()
{
    sudo apt update
    sudo apt install -y \
        zsh \
        bat \
        tldr \
        progress \
        htop \
        pipx \
        unzip
}

if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$NAME
    VER=$VERSION_ID
fi

echo "Detected $OS $VER"

case $OS in
    "Fedora Linux")
    setup_fedora
    ;;
    "Arch Linux")
    setup_arch
    ;;
    "Debian GNU/Linux")
    setup_debian
    ;;
    "Ubuntu")
    setup_debian
    ;;
    *) echo "$OS is not yet supported. Feel free to make a pull request and add support for your distro.";;
esac

echo "Changing to zsh shell for $USER"
chsh -s $(which zsh)

echo "Creating .zsh_history and .zshrc"

touch $HOME/.zsh_history
cat >> $HOME/.zshrc<< EOF
source $HOME/.shell/antigen.zsh

THEME_DIR=$HOME/.cache/oh-my-posh/themes
OMP_THEME="powerlevel10k_classic"
export ZSH_AUTOSUGGEST_STRATEGY=(history completion)

export PATH="$PATH:$HOME/.local/bin"

eval "$(oh-my-posh init zsh --config $THEME_DIR/$OMP_THEME.omp.json)"
eval "$(mcfly init zsh)"

alias ls="pls -a -d perms -d user -d group -d size -d mtime -d git"
alias cp="cp -i"
alias cat="bat"
alias codium="flatpak run com.vscodium.codium "

antigen bundle zsh-users/zsh-completions
antigen bundle zsh-users/zsh-autosuggestions
antigen bundle zsh-users/zsh-syntax-highlighting
antigen apply

HISTFILE=~/.zsh_history
HISTSIZE=10000
SAVEHIST=10000
setopt appendhistory

# Key Bindings
bindkey '^[[H' beginning-of-line
bindkey '^[[F' end-of-line
bindkey '^[[3~' delete-char
EOF

mkdir -p ~/.shell

echo "Installing Oh My Posh"
curl -s https://ohmyposh.dev/install.sh | bash -s

echo "Installing antigen"
curl -L git.io/antigen > ~/.shell/antigen.zsh

echo "Installing mcfly"
curl -LSfs https://raw.githubusercontent.com/cantino/mcfly/master/ci/install.sh | sudo sh -s -- --git cantino/mcfly

pipx install pls

echo "Setup complete. Please restart your session to load changes."
