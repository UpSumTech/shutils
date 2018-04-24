#! /usr/bin/env bash

ls -lah # Display all files in current dir in human readable format
ls -ltr # Display files newest last
du -sh ~/**/* | sort -rn # Display dirs in sorted sizes
df -h # Display free disk space in human readable format
cat /proc/partitions # Display partitions
lsblk -f # Display block storage devices as a tree view. Should list the same partitions above.
mount | column -t # Get all mounted filesystems
cat /etc/fstab # Get the static file system info

free -m # Get free memory statistics for the system
