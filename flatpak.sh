#!/usr/bin/env bash

flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo

flatpak install -y \
  com.vivaldi.Vivaldi \
  com.brave.Browser \
  org.telegram.desktop \
  io.github.spacingbat3.webcord \
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
  org.jdownloader.JDownloader \
  com.obsproject.Studio \
  org.strawberrymusicplayer.strawberry \
  org.audacityteam.Audacity \
  io.gitlab.news_flash.NewsFlash \
  com.slack.Slack
