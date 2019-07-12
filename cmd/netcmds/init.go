package netcmds

import (
	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of networking command`
	parseLongDesc  = `Prints examples of complex networking commands`
	parseExample   = `
	### Available commands for netcmds
	shutils net`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "net [sub]",
		Short:   parseShortDesc,
		Long:    parseLongDesc,
		Example: parseExample,
	}

	cmd.AddCommand(IptablesCmds())
	cmd.AddCommand(Iproute2Cmds())
	cmd.AddCommand(MiscCmds())
	return cmd
}
