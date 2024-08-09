#!/usr/bin/env bash

echo "Creating mount directories"
sudo mkdir -p /mnt/media
sudo mkdir -p /mnt/data
#sudo mkdir -p /mnt/windows
#sudo mkdir -p /mnt/winstorage

sudo chown -R $USER:$USER /mnt/*

echo "Adding mount points to /etc/fstab"
echo $'\n\n# NAS shares' | sudo tee -a /etc/fstab
echo "//memoryalpha.home.local/media  /mnt/media  cifs  credentials=$HOME/.smbcredentials,uid=1000,gid=1000,file_mode=0775,dir_mode=0775,_netdev,iocharset=utf8,noperm  0 0" | sudo tee -a /etc/fstab
echo "//memoryalpha.home.local/data   /mnt/data   cifs  credentials=$HOME/.smbcredentials,uid=1000,gid=1000,file_mode=0775,dir_mode=0775,_netdev,iocharset=utf8,noperm  0 0" | sudo tee -a /etc/fstab

echo "Mounting partitions"
sudo mount -a