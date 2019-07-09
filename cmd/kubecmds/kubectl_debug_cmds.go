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

# To check the state of a cluster and all of its components
kubectl get componentstatus

# Drain pods from a node that is unhealthy
kubectl drain node.example.com --ignore-daemonsets --delete-local-data

# Quick sample deployment to test scheduling on a node
kubectl run hello-world --replicas=1 --image=gcr.io/google-samples/node-hello:1.0  --port=8080 --overrides='{ "apiVersion": "apps/v1beta1", "spec": { "template": { "spec": { "nodeSelector": { "kubernetes.io/hostname": "nodename.example.com" } } } } }'

# Quickly drain node on kube cluster
kubectl drain <NODE_NAME> --ignore-daemonsets --delete-local-data

# To wipe out config for a cluster and start afresh
# Especially required if certs or auth has changed for the k8s cluster
kubectl config unset users.my.cluster.name.co-basic-auth
kubectl config delete-context my.cluster.name

# To get the pods with a label and custom output columns
# This is especially useful when iterating over multiple pods at once with something like xargs
kubectl get pod -l app=<app-label> --no-headers -o custom-columns=:.metadata.name,:.spec.nodeName

# To test access to k8s objects by impersonating a different user and group
kubectl get pod nginx-hjd72gsj-xj6pc --as=new_dev@bench.co --as-group=FrontEnd
			`)
		},
	}

	return cmd
}
