#! /usr/bin/env bash

uptime # How long has this machine been running
who # who is currently logged into the system
w # List processes belonging to logged in users

last # List last logged in users
last reboot # List last time the system restarted
last root # List last time the system was accessed as root

wall 'Taking down system for maintenance' # Send a warning to other users to wrap up their work before maintenance window begins
lsb_release -a # List distro info
lsmod # Lsit all the kernel modules
dpkg -l # List all packages
lscpu # List cpu info
lspci -mm # Lsit all PCI buses in the system in machine readable format

# Scan for binaries that have suid bit set for user and group
find / -xdev -type f -perm /u+s,g+s -print # Useful for finding binaries you dont recognize

# Scenario : Running lsof shows deleted uid 999 hanging onto processes
# find user with UID 999
awk -v uid=999 -F ":" '$3==uid {print $1}' /etc/passwd
ps -U 999 # find processes owned by user 999
