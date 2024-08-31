#!/usr/bin/env bash

SERVICE_DIR="/etc/systemd/system/getty@tty1.service.d"
CONF_FILE="$SERVICE_DIR/autologin.conf"

sudo mkdir -p $SERVICE_DIR

echo "[Service]" | sudo tee $CONF_FILE
echo "ExecStart=" | sudo tee -a $CONF_FILE
echo "ExecStart=-/sbin/agetty -o '-p -f -- \\u' --noclear --autologin username %I \$TERM" | sudo tee -a $CONF_FILE
