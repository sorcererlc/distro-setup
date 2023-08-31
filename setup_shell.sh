#!/usr/bin/env bash

sudo apt update
sudo apt install -y zsh clang
chsh -s /usr/bin/zsh

touch ~/.zshrc
echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"' >> ~/.zshrc
echo 'source /home/linuxbrew/.linuxbrew/share/antigen/antigen.zsh' >> ~/.zshrc
echo 'export THEME_DIR=$(brew --prefix oh-my-posh)/themes' >> ~/.zshrc
echo 'export OMP_THEME="quick-term"' >> ~/.zshrc
echo 'eval "$(oh-my-posh init zsh --config ${THEME_DIR}/${OMP_THEME}.omp.json)"' >> ~/.zshrc
echo 'eval "$(mcfly init zsh)"' >> ~/.zshrc
echo 'alias ls="pls -a -d perms -d user -d group -d size -d mtime -d git"' >> ~/.zshrc
echo 'export ZSH_AUTOSUGGEST_STRATEGY=(history completion)' >> ~/.zshrc
echo 'antigen bundle zsh-users/zsh-completions' >> ~/.zshrc
echo 'antigen bundle zsh-users/zsh-autosuggestions' >> ~/.zshrc
echo 'antigen bundle zsh-users/zsh-syntax-highlighting' >> ~/.zshrc
echo 'antigen apply' >> ~/.zshrc
echo 'HISTFILE=~/.zsh_history' >> ~/.zshrc
echo 'HISTSIZE=10000' >> ~/.zshrc
echo 'SAVEHIST=10000' >> ~/.zshrc
echo 'setopt appendhistory' >> ~/.zshrc

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"

brew tap cantino/mcfly

brew install antigen
brew install jandedobbeleer/oh-my-posh/oh-my-posh
brew install cantino/mcfly/mcfly
brew install pipx

pipx install pls
pipx ensurepath

echo "Setup complete. Please restart your session to load changes."
