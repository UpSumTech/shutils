#! /usr/bin/env bash

sudo apt-get install moreutils # Just install moreutils and thank me later
sudo apt-get install iputils-arping # package to send ARP requests at the ethernet level
sudo apt-get install tcptraceroute # useful to trace route with tcp instead of icmp
sudo apt-get install procinfo # useful to read proc files for socket information
sudo apt-get install sockstat # useful to read socket information
sudo apt-get install nmap # useful for scanning and trouble shooting large networks
sudo apt-get install multitail # useful for tailing multiple logs in one terminal window

cat /etc/*-release # To know what flavour of linux you are running
uname -m # System architecture
arch # System architecture
lsb_release -a # List distro info
lsmod # Lsit all the kernel modules
dpkg -l # List all packages in debian systems
apt list --installed | grep ssl # List installed packages in ubuntu
lscpu # List cpu info
lspci -mm # List all PCI buses in the system in machine readable format

rpm -qa # List of all rpm packages on a RHEL system
rpm -qf /bin/echo # List the package the file originated from

# To get a quick summary of what this binary might be doing. If you suspect something you have never seen before
whatis nc
