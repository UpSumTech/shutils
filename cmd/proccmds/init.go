package proccmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands for debugging processes`
	parseLongDesc  = `Prints examples of commands for debugging processes`
	parseExample   = `
	### Example commands for debugging processes
	shutils proc`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "proc [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
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

pstree -aclp <pid> # Get an uncompressed version of process tree. Useful for process managers like supervisor

fuser -av /proc/meminfo # Find what process is using this file. Try this with the top command
fuser -av 22/tcp # Find what processes are using this socket
fuser -cuk 22/tcp # Find and kill all processes using this socket

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
lsof -p PID | grep 'cwd' # To find what the current working dir of the command is
find /proc/<PID>/fd -type l | xargs ls -lah # List the open file descriptors of a process
lsof -R | grep <open file descriptor of PIPE> | awk '{print system("ps "$2)}' # This tells you which process is listening to the other end of the pipe
lsof -P -iTCP -sTCP:LISTEN | grep <open file descriptor of socket> # This tells you what socket the process is listening to

cat /proc/<pid>/environ | strings # what env vars were set when the process started
cat /proc/<pid>/cmdline | strings | xargs # regen the exact cmd that was used to run the process

# detach top running with absolute values and delayed by 5 seconds from login session and collect the logs
nohup top -e -s 5 >> top.log 2>&1 & # nohup doesnt background the process by default
echo $! > top.pid # Capture the pid of the top command to be able to kill later
kill -9 $(cat top.pid) && rm top.pid # Kill nohup process after done

ps U user # Kill all processes belonging to a user
slay -clean user # Kills all remaining processes from the above step

cat /etc/services | grep -i tmux # To see the port and protocol that tmux is using

# For debugging system and library calls in programs
strace ls -i foo # Trace the system calls for the command
ltrace ls -i foo # Trace the library calls for the command

# To run a process in a detached screen session
screen -S <screen-session-name> -dm bash -c "htop"
# To kill a screen session programatically
screen -XS "<screen-session-name>" quit
sudo su -l root -c "screen -S <screen-session-name> -dm bash -c \"tcpdump -vvv -i any -s0 -A -w nginx.dump\""

service --status-all # List all services on an ubuntu box

# To find the path of coredumps
cat /proc/sys/kernel/core_pattern
# To generate a coredump file for testing try this
ulimit -c unlimited
kill -s SIGSEGV $$
			`)
		},
	}

	return cmd
}
