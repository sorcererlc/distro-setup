#!/usr/bin/env bash

echo "Updating repositories"
sudo dnf install -y "https://mirrors.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm" "https://mirrors.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm"
sudo dnf copr enable -y codifryed/CoolerControl
sudo dnf update -y

echo "Installing NVIDIA driver"
sudo dnf install -y akmod-nvidia xorg-x11-drv-nvidia-cuda

echo "Installing KDE Plasma"
sudo dnf install -y \
  @kde-desktop-environment \
  plasma-workspace-x11 \
  sddm

echo "Switching desktop manager to sddm"
sudo systemctl disable gdm
sudo systemctl enable sddm

echo "Installing software"
flatpak install -y \
  com.vivaldi.Vivaldi \
  com.brave.Browser \
  com.spotify.Client \
  com.valvesoftware.Steam \
  org.telegram.desktop \
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
  io.github.TransmissionRemoteGtk \
  org.freedesktop.Platform.VulkanLayer.MangoHud \
  org.jdownloader.JDownloader \
  com.obsproject.Studio \
  org.strawberrymusicplayer.strawberry \
  org.audacityteam.Audacity

sudo dnf install -y \
  input-remapper \
  lm_sensors \
  dnf-plugins-core \
  coolercontrol \
  celluloid \
  virt-manager \
  kate \
  darktable \
  nomacs \
  spacenavd \
  openrgb \
  ufw \
  firejail

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

echo "Enabling and starting services"
sudo systemctl enable --now input-remapper
sudo systemctl enable --now coolercontrold
sudo systemctl enable --now spacenavd

echo "Creating mount directories"
sudo mkdir -p /mnt/storage
sudo mkdir -p /mnt/sata
sudo mkdir -p /mnt/windows
sudo mkdir -p /mnt/media
sudo mkdir -p /mnt/data

echo "Adding mount points to /etc/fstab"
cat >> /etc/fstab<< EOF

# Local partitions
PARTUUID="550393b3-fbca-4aea-a1dd-13512a81a234"   /mnt/storage    ntfs    uid=1000,gid=1000,nofail 0 2
PARTUUID="f2cfc843-afed-4aaf-a729-e237f4ab2d97"   /mnt/sata       ntfs    uid=1000,gid=1000,nofail 0 2
PARTUUID="7046eacd-7ed6-4005-b237-74e87ae708b2"   /mnt/windows    ntfs    uid=1000,gid=1000,nofail 0 2

# NAS shares
//memoryalpha.home.local/media                    /mnt/media      cifs    credentials=/home/sorcerer/.smbcredentials,uid=1000,gid=1000,file_mode=0775,dir_mode=0775,_netdev,iocharset=utf8,noperm   0 0
//memoryalpha.home.local/data                     /mnt/data       cifs    credentials=/home/sorcerer/.smbcredentials,uid=1000,gid=1000,file_mode=0775,dir_mode=0775,_netdev,iocharset=utf8,noperm   0 0
//memoryalpha.home.local/media                    /home/sorcerer/.var/app/org.jdownloader.JDownloader/downloads       cifs    credentials=/home/sorcerer/.smbcredentials,uid=1000,gid=1000,file_mode=0775,dir_mode=0775,_netdev,iocharset=utf8,noperm   0 0
EOF

echo "Mounting partitions"
sudo mount -a

echo "Setting up shell"
chmod +x setup_shell.sh
./setup_shell.sh
