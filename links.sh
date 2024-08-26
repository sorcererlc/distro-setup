#!/usr/bin/env bash

SOURCE=/mnt$HOME
DEST=$HOME

for dir in $SOURCE/.var/app/*/
do
  dir=${dir%*/}
  ln -s ${dir%*/} $DEST/.var/app/${dir##*/}
done

#for dir in $SOURCE/.config/*/
#do
#  dir=${dir%*/}
#  ln -s ${dir%*/} $DEST/.config/${dir##*/}
#done

#for dir in $SOURCE/.local/share/*/
#do
#  dir=${dir%*/}
#  ln -s ${dir%*/} $DEST/.local/share/${dir##*/}
#done
