#!/usr/bin/env bash

echo "Installing and updating repositories"
sudo pacman -Syu

echo "Installing NVIDIA driver"
sudo pacman -Sy --needed \
  nvidia \
  nvidia-settings

# Set NVIDIA DRM params
sudo echo -e "options nvidia_drm modeset=1 fbdev=1" | sudo tee -a /etc/modprobe.d/nvidia.conf
# Load NVIDIA kernel modules
sudo sed -Ei 's/^(MODULES=\([^\)]*)\)/\1 nvidia nvidia_modeset nvidia_uvm nvidia_drm)/' /etc/mkinitcpio.conf
# Blacklist nouveau driver
#echo "blacklist nouveau" | sudo tee -a "$NOUVEAU"
#echo "install nouveau /bin/true" | sudo tee -a "/etc/modprobe.d/blacklist.conf"

echo "Installing paru"
git clone https://aur.archlinux.org/paru.git
cd paru
makepkg -si

echo "Installing software"
sudo pacman -Sy --needed \
  git \
  nano \
  vi \
  base-devel \
  dolphin \
  flatpak \
  steam \
  lm_sensors \
  celluloid \
  vlc \
  virt-manager \
  kate \
  darktable \
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
  zsh \
  bat \
  tldr \
  progress \
  htop \
  exa \
  unzip

paru -Sy \
  input-remapper-git \
  coolercontrol \
  flatseal \
  nomacs \
  spacenavd \
  auto-cpufreq

echo "Installing Advanced Copy"
curl https://raw.githubusercontent.com/jarun/advcpmv/master/install.sh --create-dirs -o ./advcpmv/install.sh && (cd advcpmv && sh install.sh)
sudo mv ./advcpmv/advcp /usr/local/bin/cpg
sudo mv ./advcpmv/advmv /usr/local/bin/mvg
rm -rf ./advcpmv