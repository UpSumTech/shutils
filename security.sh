#! /usr/bin/env bash

runlevel # Get the current run level of the system
uptime # How long has this machine been running
who # who is currently logged into the system
w # List processes belonging to logged in users

last # List last logged in users
last reboot # List last time the system restarted
last root # List last time the system was accessed as root

wall 'Taking down system for maintenance' # Send a warning to other users to wrap up their work before maintenance window begins

cat /etc/*-release # To know what flavour of linux you are running
uname -m # System architecture
arch # System architecture
lsb_release -a # List distro info
lsmod # Lsit all the kernel modules
dpkg -l # List all packages in debian systems
lscpu # List cpu info
lspci -mm # List all PCI buses in the system in machine readable format

rpm -qa # List of all rpm packages on a RHEL system
rpm -qf /bin/echo # List the package the file originated from

whatis nc # To get a quick summary of what this binary might be doing. If you suspect something you have never seen before.

# Scan for binaries that have suid bit set for user and group
find / -xdev -type f -perm /u+s,g+s -print # Useful for finding binaries you dont recognize

# Scenario : Running lsof shows deleted uid 999 hanging onto processes
# find user with UID 999
awk -v uid=999 -F ":" '$3==uid {print $1}' /etc/passwd
ps -U 999 # find processes owned by user 999

ls -i # List inode or physical addresses of files. Sometimes useful to find files with weird chars in names
zdump PST EST IST # list current time in different time zones

# To look at aggregated logs from a bunch of servers
multitail --merge-all -cS apache -cS log4j -e 'error' -l 'ssh -t user@server1 "tail -f /var/log/nginx.log"' -cS apache -cS log4j -e 'error' -l 'ssh -t user@server2 "tail -f /var/log/nginx.log"' --no-mergeall
