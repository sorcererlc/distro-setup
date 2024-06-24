#!/usr/bin/env bash

setup_fedora()
{
    sudo dnf update -y
    sudo dnf install -y \
        zsh \
        clang \
        bat \
        tldr \
        progress \
        htop
}

setup_debian()
{
    sudo apt update
    sudo apt install -y zsh clang
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
    "Debian GNU/Linux")
    setup_debian
    ;;
    "Ubuntu")
    setup_debian
    ;;
    *) echo "$OS is not yet supported. Feel free to make a pull request and add support for your distro.";;
esac

echo "Changing to zsh shell for $USER"
chsh -s /usr/bin/zsh

touch ~/.zsh_history
cat >> ~/.zshrc<< EOF
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
source /home/linuxbrew/.linuxbrew/share/antigen/antigen.zsh

THEME_DIR=$(brew --prefix oh-my-posh)/themes
OMP_THEME="quick-term"
export ZSH_AUTOSUGGEST_STRATEGY=(history completion)

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

export PATH="$PATH:$HOME/.local/bin"
EOF

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"

brew tap cantino/mcfly

brew install antigen
brew install jandedobbeleer/oh-my-posh/oh-my-posh
brew install cantino/mcfly/mcfly
brew install pipx

pipx install pls

echo "Setup complete. Please restart your session to load changes."
