#!/usr/bin/env bash

if ! grep -qE '^MODULES=.*nvidia. *nvidia_modeset.*nvidia_uvm.*nvidia_drm' /etc/mkinitcpio.conf; then
  sudo sed -Ei 's/^(MODULES=\([^\)]*)\)/\1 nvidia nvidia_modeset nvidia_uvm nvidia_drm)/' /etc/mkinitcpio.conf
fi

if [ ! -f /etc/default/grub ]; then
  if ! sudo grep -q "nvidia-drm.modeset=1" /etc/default/grub; then
    sudo sed -i 's/\(GRUB_CMDLINE_LINUX_DEFAULT=".*\)"/\1 nvidia-drm.modeset=1"/' /etc/default/grub
    sudo grub-mkconfig -o /boot/grub/grub.cfg
  fi
fi

if [ ! -f /etc/modprobe.d/nvidia.conf ]; then
  echo -e "options nvidia_drm modeset=1 fbdev=1" | sudo tee /etc/modprobe.d/nvidia.conf
fi
# echo "blacklist nouveau" | sudo tee "/etc/modprobe.d/nouveau.conf"
# echo "install nouveau /bin/true" | sudo tee -a "/etc/modprobe.d/blacklist.conf"
