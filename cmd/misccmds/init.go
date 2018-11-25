package misccmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Miscellaneous commands that belong nowhere`
	parseLongDesc  = `Miscellaneous commands that belong nowhere`
	parseExample   = `
	### Miscellaneous commands
	shutils misc`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "misc [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
echo $(expr 1 + 2) # math with int
cal 8 2017 # display calendar
date +"%Y/%m/%d:%H:%M:%S" # display date in a specific format
uuidgen # Generates a uuid

# To look at aggregated logs from a bunch of servers
multitail --merge-all -cS apache -cS log4j -e 'error' -l 'ssh -t user@server1 "tail -f /var/log/nginx.log"' -cS apache -cS log4j -e 'error' -l 'ssh -t user@server2 "tail -f /var/log/nginx.log"' --no-mergeall
			`)
		},
	}

	return cmd
}
