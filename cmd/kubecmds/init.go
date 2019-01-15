package kubecmds

import (
	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of kubectl command`
	parseLongDesc  = `Prints examples of complex kubectl commands that are unusual`
	parseExample   = `
	### Available commands for kubecmds
	shutils kubectl`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "kubectl [sub]",
		Short:   parseShortDesc,
		Long:    parseLongDesc,
		Example: parseExample,
	}

	cmd.AddCommand(KubectlPodCmds())
	cmd.AddCommand(KubectlRunCmds())
	cmd.AddCommand(KubectlDeployCmds())
	cmd.AddCommand(KubectlDebugCmds())
	return cmd
}
