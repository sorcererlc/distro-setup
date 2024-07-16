#!/usr/bin/env bash

echo "Installing and updating repositories"
sudo dnf install -y \
  "https://mirrors.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm" \
  "https://mirrors.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm"
sudo dnf copr enable -y codifryed/CoolerControl
sudo dnf update -y

flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo

echo "Installing NVIDIA driver"
sudo dnf install -y \
  akmod-nvidia
  # xorg-x11-drv-nvidia-cuda

# echo "Installing KDE Plasma"
# sudo dnf swap @gnome-desktop @kde-desktop
# sudo dnf install -y \
#   sddm \
  # plasma-workspace-x11

# echo "Switching desktop manager to sddm"
# sudo systemctl disable gdm
# sudo systemctl enable sddm

echo "Installing software"
sudo dnf install -y \
  steam \
  guake \
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
  neovim

flatpak install -y \
  com.vivaldi.Vivaldi \
  com.brave.Browser \
  com.spotify.Client \
  io.github.tdesktop_x64.TDesktop \
  com.discordapp.Discord \
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
  org.equeim.Tremotesf \
  org.freedesktop.Platform.VulkanLayer.MangoHud \
  org.jdownloader.JDownloader \
  com.obsproject.Studio \
  org.strawberrymusicplayer.strawberry \
  org.audacityteam.Audacity \
  io.gitlab.news_flash.NewsFlash \
  com.slack.Slack

echo "Setting up firewall"
sudo ufw reset                # Delete all existing rules
sudo ufw limit 22/tcp         # SSH
sudo ufw allow 4950/udp       # Warframe
sudo ufw allow 4955/tcp       # Warframe
sudo ufw enable

echo "Detecting hardware sensors"
sudo sensors-detect

echo "Setting up service configuration"
sudo cp -r /home/sorcerer/.services/etc/* /etc

echo "Installing auto-cpufreq"
# cd /home/sorcerer/Programs
# git clone https://github.com/AdnanHodzic/auto-cpufreq.git
# cd auto-cpufreq
cd /home/sorcerer/Programs/auto-cpufreq
git pull
sudo ./auto-cpufreq-installer

echo "Enabling and starting services"
sudo systemctl enable --now sshd
sudo systemctl enable --now input-remapper
sudo systemctl enable --now coolercontrold
sudo systemctl enable --now spacenavd
sudo systemctl enable --now auto-cpufreq

echo "Creating mount directories"
sudo mkdir -p /mnt/media
sudo mkdir -p /mnt/data
#sudo mkdir -p /mnt/windows
#sudo mkdir -p /mnt/winstorage

echo "Adding mount points to /etc/fstab"
sudo echo $'\n\n# NAS shares' >> /etc/fstab
sudo echo "//memoryalpha.home.local/media  /mnt/media  cifs  credentials=/home/sorcerer/.smbcredentials,uid=1000,gid=1000,file_mode=0775,dir_mode=0775,_netdev,iocharset=utf8,noperm  0 0" >> /etc/fstab
sudo echo "//memoryalpha.home.local/data   /mnt/data   cifs  credentials=/home/sorcerer/.smbcredentials,uid=1000,gid=1000,file_mode=0775,dir_mode=0775,_netdev,iocharset=utf8,noperm  0 0" >> /etc/fstab

echo "Mounting partitions"
sudo mount -a
