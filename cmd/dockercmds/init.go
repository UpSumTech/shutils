package dockercmds

import (
	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of docker commands`
	parseLongDesc  = `Prints examples of complex docker commands that are unusual`
	parseExample   = `
	### Available commands for docker
	shutils docker`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "docker [sub]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
	}

	cmd.AddCommand(DockerCliCmds())
	cmd.AddCommand(DockerDebugCmds())
	return cmd
}
