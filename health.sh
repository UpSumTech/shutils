#! /usr/bin/env bash

#### Scenario : Find the load average of a Linux system ####
cat /proc/loadavg # Check for the load average of a system
w # Shows who is logged into a system
uptime # See the system uptime

#### Scenario : VM is running low on virtual memory ####
#### This will however not persist across reboots
df -h # Look at the current disk usage
dd if=/dev/zero of=/mnt/swapfile bs=1M count=1025 # Create a swap file of 1024Mb to increase virtual memory of the system
swapon -s # List the current swap partitions
swapon /mnt/swapfile # Activate the new swap file
