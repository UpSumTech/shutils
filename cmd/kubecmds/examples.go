package kubecmds

import (
	"fmt"

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
		Use:     "kubectl",
		Short:   parseShortDesc,
		Long:    parseLongDesc,
		Example: parseExample,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# k8s get the images for pods quickly
kubectl get pods -l labelName=labelValue -o jsonpath="{..image}" | tr ' ' '\n'

# k8s get the names of the pods
kubectl get pod -l labelName=labelValue -o jsonpath="{..metadata.name}"

# k8s to rollback deployment
kubectl rollout undo deployment deploymentName

# run a kubernetes pod with a specific service account
kubectl run -i -t --rm test-pod --image=ubuntu:16.04 --restart=Never --serviceaccount=specificServiceAccount --image-pull-policy=Always --env FOO=$FOO --env BAR=$BAR
			`)
		},
	}

	return cmd
}
