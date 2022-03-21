package memcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands to operate on memory`
	parseLongDesc  = `Prints examples of commands to operate on memory`
	parseExample   = `
	### Available commands for memcmds
	shutils mem`
)

// Init instantiates the disk commands
func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "mem [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# To get the basic idea of how processes are using memory you can use ps.
# If the RSS of a process is growing it is generally indicative of some sort of a memory issue.
# VSZ/VIRT is virtual memory and represents the total memory mapped by a proc. This number is not of much interest.
# RSS/RES is physical memory and represents the total physical memory used by a proc. This number is of more interest. Although remember that this includes memory pages shared with other procs.
ps -eo pid,tid,class,rtprio,stat,vsz,rss,comm

# Recursively search the logs to find out if anything encountered an OOM error.
# You would generally see a OOM score of the process and the process name and pid.
# Based on that ypou could reduce the amount of memory requested by the proc, disallow proc to overcommit memory etc.
grep -i -r 'out of memory' /var/log/

# Get current memory stats
# There is info on application used memory, buffers and caches. Cached data is usually that's in hard disk but frequently being used and so is stored in cache.
# To the apps this is usually treated as free memory
free -h

top
ps aux | grep <proc_name>
cat /proc/<pid>/oom_score # If this is a very high number then there is a chance the proc is leaking memory and should be looked into.

# In linux systems the kernel allows requesting more memory than there is in the system for better memory utilization.
# However, it is possible to disallow memory overcommit.
# sysctl is the command that is used to modify the values of various kernel parameters
# This command lists all sysctl controlled parameters
sysctl -a

# The values of this parameter vm.overcommit_memory can be 0, 1, 2.
#   0 - estimate if we have enough RAM
#   1 - always allow
#   2 - say no if system doesnt have memory
# Another similar parameter is - vm.overcommit_ratio
# Remember to change the overcommit ratio as well when you change the overcommit memory
sysctl -w vm.overcommit_memory=2
sysctl -w vm.overcommit_ratio=100

# However, if you do not write these change to sysctl.conf these changes will not be persisted across system reboot
# To do that you can make the following changes
echo 'vm.overcommit_memory=2' >> /etc/sysctl.conf
echo 'vm.overcommit_ratio=100' >> /etc/sysctl.conf
sysctl -p # This is to reload the settings of sysctl.conf

# You could look at the memory consumption of a pid using pmaps utility
pmap -X <pid>
pmap --extended <pid>

# Linux also shows the details of the memory consumption of a process in the /proc/<pid>/smaps file
cat /proc/1/smaps
cat /proc/1/smaps | grep -i pss |  awk '{Total+=$2} END {print Total/1024" MB"}'

# If you want to collect coredumps and share that with the developers you can use the ulimit utility to allow the system to take coredumps
# In ubuntu generally the coredumps go to /var/crash/* . The coredumps are taken by a utility called 'apport'. In RHEL the related tools are abrt/abrt-addon-ccpp/abrt-cli etc.
ulimit -c unlimited
sysctl kernel.core_pattern # Check this to see where the coredumps are getting generated
cat /proc/sys/kernel/core_pattern # OR you can directly check the file for the kernel setting

sysctl -w kernel.core_pattern=<pattern>
systemctl restart apport # To restore the default settings of apport

# Also, you can use an external tool like valgrind to get memory leak errors.
# Here for example we are testing the command with arguments for sleep. And this should show no errors when you inspect the log file for the report
valgrind --leak-check=full --log-file=/tmp/mem-leak-sleep.log sleep 10
			`)
		},
	}

	return cmd
}
