#!/usr/bin/env bash

PACMAN_CONF="/etc/pacman.conf"

# Remove comments '#' from specific lines
LINES=(
    "Color"
    "CheckSpace"
    "VerbosePkgLists"
    "ParallelDownloads"
)

# Uncomment specified lines if they are commented out
for LINE in "${LINES[@]}"; do
    if grep -q "^#$LINE" "$PACMAN_CONF"; then
        sudo sed -i "s/^#$LINE/$LINE/" "$PACMAN_CONF"
    fi
done

# Add "ILoveCandy" below ParallelDownloads if it doesn't exist
if grep -q "^ParallelDownloads" "$PACMAN_CONF" && ! grep -q "^ILoveCandy" "$PACMAN_CONF"; then
    sudo sed -i "/^ParallelDownloads/a ILoveCandy" "$PACMAN_CONF"
fi

# updating pacman.conf
sudo pacman -Sy

