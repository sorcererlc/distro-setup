#!/usr/bin/env bash

setup_fedora()
{
    ./fedora.sh
}

setup_arch()
{
    ./arch.sh
}

setup_debian()
{
    ./debian.sh
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

echo "Detecting hardware sensors"
sudo sensors-detect

echo "Setting up service configuration"
sudo cp -r ./services/etc/* /etc

echo "Enabling and starting services"
sudo systemctl enable --now sshd
sudo systemctl enable --now input-remapper
sudo systemctl enable --now coolercontrold
sudo systemctl enable --now spacenavd
sudo systemctl enable --now auto-cpufreq

./flatpak.sh
./shares.sh
./udev.sh
./shell.sh

read -p "Setup complete. You must reboot for the changes to apply. Do you want to reboot now? (y/n)" proceed
if [ "$proceed" != "y" ]; then
    echo "Reboot skipped."
    exit 0
fi

sudo reboot