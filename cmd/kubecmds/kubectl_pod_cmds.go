package kubecmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func KubectlPodCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pod [no options!]",
		Short: `Complex or unusual kubectl pod command examples`,
		Long:  `Complex or unusual kubectl pod command examples`,
		Example: `
### Example commands for kubectl pods
shutils kubectl pod
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# k8s get the images for pods quickly
kubectl get pods -l labelName=labelValue -o jsonpath="{..image}" | tr ' ' '\n'

# k8s get the names of the pods
kubectl get pod -l labelName=labelValue -o jsonpath="{..metadata.name}"
			`)
		},
	}

	return cmd
}
