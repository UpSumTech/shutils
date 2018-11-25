package kubecmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func KubectlDeployCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy [no options!]",
		Short: `Complex or unusual kubectl deployment related command examples`,
		Long:  `Complex or unusual kubectl deployment related command examples`,
		Example: `
### Example commands for kubectl deployments
shutils kubectl deploy
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# k8s to rollback deployment
kubectl rollout undo deployment deploymentName
			`)
		},
	}

	return cmd
}
