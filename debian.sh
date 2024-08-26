#!/usr/bin/env bash

echo "Installing software"
sudo apt-get update
sudo apt-get install -y \
  zsh \
  bat \
  tldr \
  progress \
  htop \
  exa \
  unzip

echo "Installing Advanced Copy"
curl https://raw.githubusercontent.com/jarun/advcpmv/master/install.sh --create-dirs -o ./advcpmv/install.sh && (cd advcpmv && sh install.sh)
sudo mv ./advcpmv/advcp /usr/local/bin/cpg
sudo mv ./advcpmv/advmv /usr/local/bin/mvg
rm -rf ./advcpmv