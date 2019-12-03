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

# kubectl run arbitrary pod and execute some task quickly
kubectl run test-pod --record --wait=true --image-pull-policy=Always --image=ubuntu:16.04 --restart=Never --env="DB_HOST=mysql.example.com" --env="DB_PORT=3306" --command -- /bin/bash -c 'while true; do sleep 3; done'; sleep 10; kubectl get pod test-pod; kubectl cp $HOME/.secrets test-pod:/root/.secrets; kubectl cp $(pwd) test-pod:/root/example-repo; kubectl exec test-pod -- /bin/bash -c "cd /root/example-repo; ./bin/run.sh"; kubectl delete pod test-pod

# If you want to annotate a pod with a specific iam role that can be assumed with kube2iam
kubectl annotate pods test-pod "iam.amazonaws.com/role"="arn:aws:iam::<accountid>:role/<some_iam_role>"
			`)
		},
	}

	return cmd
}
