package kubecmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func KubectlRunCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run [no options!]",
		Short: `Complex or unusual kubectl run command examples`,
		Long:  `Complex or unusual kubectl run command examples`,
		Example: `
### Example commands for kubectl run
shutils kubectl run
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# run a kubernetes pod with a specific service account
kubectl run -i -t --rm test-pod --image=ubuntu:16.04 --restart=Never --serviceaccount=specificServiceAccount --image-pull-policy=Always --env FOO=$FOO --env BAR=$BAR
			`)
		},
	}

	return cmd
}
