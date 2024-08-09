#!/usr/bin/env bash

setup_fedora()
{
    sudo dnf upgrage -y
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
        pyton-pipx \
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

echo "Changing shell to zsh for $USER"
chsh -s $(which zsh)

echo "Installing Advanced Copy"
curl https://raw.githubusercontent.com/jarun/advcpmv/master/install.sh --create-dirs -o ./advcpmv/install.sh && (cd advcpmv && sh install.sh)
sudo mv ./advcpmv/advcp /usr/local/bin/cpg
sudo mv ./advcpmv/advmv /usr/local/bin/mvg
rm -rf ./advcpmv

echo "Creating .zsh_history and .zshrc"
touch $HOME/.zsh_history
cp .zshrc $HOME

echo "Installing Oh My Posh"
curl -s https://ohmyposh.dev/install.sh | bash -s

echo "Installing antigen"
mkdir -p $HOME/.shell
curl -L git.io/antigen > $HOME/.shell/antigen.zsh

echo "Installing mcfly"
curl -LSfs https://raw.githubusercontent.com/cantino/mcfly/master/ci/install.sh | sudo sh -s -- --git cantino/mcfly

pipx install pls

read -p "Setup complete. You must restart your session for the changes to apply. Do you want to logout now? (y/n)" proceed
if [ "$proceed" != "y" ]; then
    echo "Logout skipped."
    exit 0
fi

logout
