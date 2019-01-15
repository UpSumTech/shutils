package kubecmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func KubectlDebugCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug [no options!]",
		Short: `Complex or unusual kubectl debug command examples`,
		Long:  `Complex or unusual kubectl debug command examples`,
		Example: `
### Example commands for debugging kubernetes objects
shutils kubectl debug
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# Get kubernetes events sorted by creation time
kubectl get events  --sort-by='.metadata.creationTimestamp' -o json

# Get kubernetes events sorted by creation time for a specific deployment
kubectl get events  --sort-by='.metadata.creationTimestamp' -o json | jq '.items | sort_by(.lastTimestamp) | .[] | select(.involvedObject.name | contains("<deployment_name>"))'

# List all env vars for a specific deployment
kubectl set env deployment/webapp --list
			`)
		},
	}

	return cmd
}
