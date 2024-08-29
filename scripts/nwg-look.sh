#!/usr/bin/env bash

DIR=$PWD

cd nwg-look
make build > $DIR/logs/nwg-look.log
sudo make install >> $DIR/logs/nwg-look.log
