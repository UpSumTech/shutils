package memcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands to operate on memory`
	parseLongDesc  = `Prints examples of commands to operate on memory`
	parseExample   = `
	### Available commands for operating on memory of a machine
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
# Lists all sysctl controlled parameters
sysctl -a

# The values of this parameter vm.overcommit_memory can be 0, 1, 2.
# 0 - estimate if we have enough RAM, 1 - always allow, 2 - say no if system doesnt have memory
# Another similar parameter is - vm.overcommit_ratio
sysctl -w vm.overcommit_memory=2
			`)
		},
	}

	return cmd
}
