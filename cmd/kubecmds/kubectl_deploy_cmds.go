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
# kubectl scale up or scale down deployments
kubectl scale deployment nginx --replicas=2

# k8s to rollback deployment
kubectl rollout undo deployment nginx

# Get all the images being run currently for all your deployments
kubectl get deployment -o custom-columns=:.metadata.name --no-headers | xargs -n 1 -I % /bin/bash -c "echo %; kubectl get deployment % -o jsonpath='{ ..image }'; echo"

# Get a particular property from all the deployments
kubectl get deployment -o custom-columns=:.metadata.name --no-headers | xargs -n 1 -I % /bin/bash -c "echo -n % \-\ ; kubectl get deployment % -o jsonpath='{ ..periodSeconds }'; echo"

# Get all deployments, change some property and redeploy by staggering it
kubectl get deployments -o custom-columns=:.metadata.name --no-headers | xargs -n 1 -I % /bin/bash -c "kubectl get deployment % -o json | sed -e 's#DB_HOST#mysql.example.com#db.example.com#g' | kubectl apply -f -; sleep 120"

# Set quickfix image for a deployment to put out a fire
kubectl set image deployment/webapp app=dockerhub.com/example.com/app:quickfix --record

# Set env var image for a deployment quickly to put out a fire
kubectl set env deployment/webapp DB_HOST=db.example.com

# To manually trigger a job from a cron job
kubectl create job --from=cronjob/<cron-job-name> <new-manual-job-name>
			`)
		},
	}

	return cmd
}
