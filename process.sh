#! /usr/bin/env bash

# Get error codes from processes
errno -l # List of error codes
errno ENOENT # description of error code from process

ps -auxefw # All processes with process tree
ps axww | grep ssh # Look for process
ps aux | grep ss[h] # Look for processes without the grep pid
pgrep -a -u root -f ssh # Look for ssh processes owned by root
ps -fp $(pgrep -u root ssh) # Print full info for processes named ssh and owned by root
ps --sort -rss -eo pid,pcpu,pmem,rss,vsize,size,cmd | head -n 10 # Track process by high memory usage
ps --sort -pcpu -eo pid,pcpu,pmem,rss,vsize,size,cmd | head -n 10 # Track process by high cpu usage
ps -A -o 'pid,ppid,stat,time,command' | awk '{if($2 == /Z/) print $0}' # Track zombie processes

fuser -av /proc/meminfo # Find what process is using this file. Try this with the top command
fuser -av 22/tcp # Find what processes are using this socket

pmap -xp $(pgrep -u root ssh) # Get the memory footprint of the ssh processes owned by root

strace top # Trace system calls for top

bg # After ctrl+z puts process in background
jobs -l # List all background jobs
fg % 2 # Bring a job to foreground

kill -s HUP PID # Send SIGHUP to process. This might be handled and process can be restarted or config reloaded.
kill -s TERM PID # Send SIGTERM to process. This can be handled and process might gracefully shut down.
kill -s KILL PID # Send SIGKILL to process. This cant be caught and might result in messed up system state.
killall -1 unicorn # Send hangup to all unicorn processes
pkill -TERM -u nobody # Kill all processes owned by user nobody. Passenger generally could be running as user nobody.
fuser -k -TERM -m 3000/tcp # Kill all application processes listening to port 3000. Useful to stop rails applications in general.

# Run top in the background and get a list of all open files with lsof
ps auww | grep to[p] | awk '{print $2}' | xargs -n 1 -I % lsof -p

lsof -i 4 -a -p PID # List open sockets for PID
lsof -p PID # List all open files for PID
