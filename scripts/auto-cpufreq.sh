#!/usr/bin/env bash

chmod +x ./auto-cpufreq/auto-cpufreq-installer
echo "i" | sudo ./auto-cpufreq/auto-cpufreq-installer
rm -rf auto-cpufreq

