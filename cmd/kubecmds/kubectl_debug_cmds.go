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

# Get kubernetes events and messages and pods that are created by specific deployment
kubectl get events --sort-by='.metadata.creationTimestamp' -o json | jq '.items | sort_by(.lastTimestamp) | .[] | select(.involvedObject.name | contains("<deployment_name>"))' | jq -r "[.message, .metadata.name, .firstTimestamp]"

# List all env vars for a specific deployment
kubectl set env deployment/webapp --list

# To check the state of a cluster and all of its components
kubectl get componentstatus

# Drain pods from a node that is unhealthy
kubectl drain node.example.com --ignore-daemonsets --delete-emptydir-data

# Quick sample deployment to test scheduling on a node
kubectl run hello-world --replicas=1 --image=gcr.io/google-samples/node-hello:1.0  --port=8080 --overrides='{ "apiVersion": "apps/v1beta1", "spec": { "template": { "spec": { "nodeSelector": { "kubernetes.io/hostname": "nodename.example.com" } } } } }'

# Quickly drain node on kube cluster
kubectl drain <NODE_NAME> --ignore-daemonsets --delete-emptydir-data

# To wipe out config for a cluster and start afresh
# Especially required if certs or auth has changed for the k8s cluster
kubectl config unset users.my.cluster.name.co-basic-auth
kubectl config delete-context my.cluster.name

# To get the pods with a label and custom output columns
# This is especially useful when iterating over multiple pods at once with something like xargs
kubectl get pod -l app=<app-label> --no-headers -o custom-columns=:.metadata.name,:.spec.nodeName

# To test access to k8s objects by impersonating a different user and group
kubectl get pod nginx-pod-1 --as=dev@example.com --as-group=FrontEnd

# To find out which k8s services are using external load balancers
kubectl get services -o custom-columns=:.metadata.name --no-headers | grep -v kubernetes | xargs -n 1 -I % /bin/bash -c "echo -n %; echo -n ' '; kubectl get service % -o jsonpath='{ ..hostname }'; echo" | tee | column -t -s ' '

# Some useful helm commands
helm plugin install <git_repo_url>
helm plugin list
helm plugin uninstall <plugin_name_from_list>

helm diff upgrade my-release stable/postgresql --values values.yaml
helm diff release my-prod my-stage
helm diff revision my-release 2 3
helm diff rollback my-release 2

helm env --vars-only

helm get $(helm last)

helm nuke

helm repo add cert-manager git+https://github.com/jetstack/cert-manager@deploy/charts?ref=v0.6.2
helm search cert-manager
helm install cert-manager/cert-manager --version "0.6.6"
helm fetch cert-manager/cert-manager --version "0.6.6"
helm install . -f git+https://github.com/aslafy-z/helm-git@tests/fixtures/example-chart/values.yaml

kubeval my-invalid-rc.yaml
kubeval --strict additional-properties.yaml
cat my-invalid-rc.yaml | kubeval
kubeval fixtures/invalid.yaml -o json

helm kubeval charts/stable/nginx-ingress
helm kubeval . -v 1.9.0
helm kubeval charts/stable/nginx-ingress --set controller.image.tag=latest

helm schema-gen values.yaml

helmsman -f example.toml
helmsman --apply -f example.toml
helmsman --debug --apply -f example.toml
helmsman --debug --dry-run -f example.toml
helmsman --debug --dry-run --target artifactory -f example.toml

helm version
helm ls
helm create api-server
			`)
		},
	}

	return cmd
}
