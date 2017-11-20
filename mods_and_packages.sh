#! /usr/bin/env bash

cat /etc/*-release # To know what flavour of linux you are running

sudo apt-get install moreutils # Just install moreutils and thank me later
sudo apt-get install iputils-arping # package to send ARP requests at the ethernet level
sudo apt-get install tcptraceroute # useful to trace route with tcp instead of icmp
sudo apt-get install procinfo # useful to read proc files for socket information
sudo apt-get install sockstat # useful to read socket information
sudo apt-get install nmap # useful for scanning and trouble shooting large networks
