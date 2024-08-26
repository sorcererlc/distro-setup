#!/usr/bin/env bash

echo "Changing shell to zsh for $USER"
chsh -s $(which zsh)

echo "Linking .zshrc"
ln -s $PWD/config/home/zshrc $HOME/.zshrc

echo "Installing Oh My Posh"
curl -s https://ohmyposh.dev/install.sh | bash -s
