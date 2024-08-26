#!/usr/bin/env bash

echo "Installing and updating repositories"
sudo dnf install -y \
  "https://mirrors.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm" \
  "https://mirrors.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm"
sudo dnf copr enable -y codifryed/CoolerControl patrickl/pipewire-wineasio
sudo dnf upgrade -y

echo "Installing NVIDIA driver"
sudo dnf install -y \
  akmod-nvidia \
  libva \
  libva-nvidia-driver
# sudo dnf install -y xorg-x11-drv-nvidia-cuda

echo "Installing software"
sudo dnf install -y \
  git \
  nano \
  vi \
  flatpak \
  steam \
  alacritty \
  flatseal \
  input-remapper \
  lm_sensors \
  dnf-plugins-core \
  coolercontrol \
  celluloid \
  vlc \
  virt-manager \
  kate \
  darktable \
  nomacs \
  spacenavd \
  openrgb \
  ufw \
  firejail \
  remmina \
  mc \
  winetricks \
  android-tools \
  neovim \
  owncloud-client \
  zoxide \
  pipewire-wineasio \
  qalculate-gtk \
  zsh \
  bat \
  tldr \
  progress \
  htop \
  exa \
  unzip \
  goverlay

echo "Setting up firewall"
sudo ufw reset                # Delete all existing rules
sudo ufw limit 22/tcp         # SSH
sudo ufw allow 4950/udp       # Warframe
sudo ufw allow 4955/tcp       # Warframe
sudo ufw enable

echo "Installing auto-cpufreq"
git clone https://github.com/AdnanHodzic/auto-cpufreq.git
sudo ./auto-cpufreq/auto-cpufreq-installer
rm -rf ./auto-cpufreq

echo "Installing Advanced Copy"
curl https://raw.githubusercontent.com/jarun/advcpmv/master/install.sh --create-dirs -o ./advcpmv/install.sh && (cd advcpmv && sh install.sh)
sudo mv ./advcpmv/advcp /usr/local/bin/cpg
sudo mv ./advcpmv/advmv /usr/local/bin/mvg
rm -rf ./advcpmv

# mkdir -p $HOME/repo
# cd $HOME/repo
# git clone https://github.com/JaKooLit/Fedora-Hyprland.git
# cd Fedora-Hyprland
# chmod +x install.sh
# ./install.sh
