package syshealthcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands to find out about the health of the system`
	parseLongDesc  = `Prints examples of commands to find out about the health of the system`
	parseExample   = `
	### Available commands for finding out about the health of the system
	shutils syshealth`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "syshealth [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
#### Scenario : Find the load average of a Linux system ####
cat /proc/loadavg # Check for the load average of a system
w # Shows who is logged into a system
uptime # See the system uptime

#### Scenario : VM is running low on virtual memory ####
#### This will however not persist across reboots
cat /proc/meminfo
vmstat
lsblk # List all block devices
df -hT # Look at the current disk usage
dd if=/dev/zero of=/mnt/swapfile bs=1M count=1025 # Create a swap file of 1024Mb to increase virtual memory of the system
swapon -s # List the current swap partitions
swapon /mnt/swapfile # Activate the new swap file

#### Journalctl usage for viewing logs ################
journalctl # For all logs
journalctl -r | less # Look at logs in reverse
journalctl -u ssh -f # To follow ssh logs
journalctl --since "2019-01-30 12:30:00" --until "2019-01-30 14:30:00" -u ssh # Show ssh logs for that period

######## Taking a heapdump for a running java service ########
export pid_of_java_process="$(jcmd -l | grep -i <name_of_jar_file>.jar | awk '{print $1}')"
jmap -dump:live,file=heapdump.hprof <pid_of_java_process> # Dump live objects only
jmap -dump:format=b,file=heapdump.hprof <pid_of_java_process> # Dump for 32 bit jvm
jmap -J-d64 -dump:format=b,file=heapdump.hprof <pid_of_java_process> # Dump for 64 bit jvm
jmap -J-d64 -dump:live,format=b,file=heapdump.hprof <pid_of_java_process> # Dump live objects for a 64 bit jvm
			`)
		},
	}

	return cmd
}
