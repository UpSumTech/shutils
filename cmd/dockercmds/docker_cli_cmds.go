package dockercmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func DockerCliCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli [no options!]",
		Short: `docker cli commands`,
		Long:  `docker cli commands`,
		Example: `
### Example commands for docker cli
shutils docker cli
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# Get the running containers
docker ps --filter status=running

# Get container id for deployment inside k8s node
docker ps --filter status=running | grep -v POD | grep -e <container_name>_<deployment_name> | cut -d " " -f1
			`)
		},
	}

	return cmd
}
