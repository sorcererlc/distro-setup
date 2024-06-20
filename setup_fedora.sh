#!/usr/bin/env bash

echo "Updating repositories..."
sudo dnf install https://mirrors.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm https://mirrors.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm
sudo dnf copr enable -y codifryed/CoolerControl
sudo dnf update -y

echo "Installing KDE Plasma..."
sudo dnf install -y @kde-desktop-environment

echo "Installing NVIDIA driver"
sudo dnf install -y akmod-nvidia xorg-x11-drv-nvidia-cuda

echo "Installing software"
flatpak install -y \
  com.vivaldi.Vivaldi \
  com.brave.Browser \
  com.spotify.Client \
  com.valvesoftware.Steam \
  org.telegram.desktop \
  com.discordapp.Discord \
  org.openrgb.OpenRGB \
  tv.plex.PlexDesktop \
  com.vscodium.codium \
  com.github.Matoking.protontricks \
  com.heroicgameslauncher.hgl \
  net.lutris.Lutris \
  org.prismlauncher.PrismLauncher \
  com.protonvpn.www \
  org.kicad.KiCad \
  rest.insomnia.Insomnia \
  org.openttd.OpenTTD \
  io.openrct2.OpenRCT2 \
  org.nomacs.ImageLounge \
  com.sublimetext.three \
  io.github.TransmissionRemoteGtk \
  org.freedesktop.Platform.VulkanLayer.MangoHud \
  org.jdownloader.JDownloader

sudo dnf install -y \
  input-remapper \
  lm_sensors \
  dnf-plugins-core \
  coolercontrol \
  celluloid

echo "Detecting hardware sensors"
sudo sensors-detect

echo "Enabling and starting services"
sudo systemctl enable --now input-remapper
sudo systemctl enable --now coolercontrold

echo "Creating mount directories"
sudo mkdir -p /mnt/storage
sudo mkdir -p /mnt/sata
sudo mkdir -p /mnt/windows
sudo mkdir -p /mnt/media
sudo mkdir -p /mnt/data
