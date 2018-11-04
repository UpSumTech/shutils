#! /usr/bin/env bash

echo `expr 1 + 2` # math with int
cal 8 2017 # display calendar
date +"%Y/%m/%d:%H:%M:%S" # display date in a specific format
uuidgen # Generates a uuid

# To look at aggregated logs from a bunch of servers
multitail --merge-all -cS apache -cS log4j -e 'error' -l 'ssh -t user@server1 "tail -f /var/log/nginx.log"' -cS apache -cS log4j -e 'error' -l 'ssh -t user@server2 "tail -f /var/log/nginx.log"' --no-mergeall
