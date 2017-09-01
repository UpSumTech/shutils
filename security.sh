#! /usr/bin/env bash

uptime # How long has this machine been running
lsb_release -a # List distro info
lsmod # Lsit all the kernel modules
dpkg -l # List all packages
lscpu # List cpu info
lspci -mm # Lsit all PCI buses in the system in machine readable format

