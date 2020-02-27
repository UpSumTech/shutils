package rabbitmqcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands for debugging rabbitmq`
	parseLongDesc  = `Prints examples of commands for debugging rabbitmq`
	parseExample   = `
	### Example commands for debugging rabbitmq
	shutils rabbitmq`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "rabbitmq [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
rabbitmqctl list_vhost <vhost_name>
rabbitmqctl add_vhost <vhost_name>
			`)
		},
	}

	return cmd
}
