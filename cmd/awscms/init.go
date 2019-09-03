package awscmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `awscli useful commands`
	parseLongDesc  = `awscli useful commands`
	parseExample   = `
	### awscli useful commands
	shutils awscmds`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "awscmds [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
######### Simple, built-in filtering #########

# List instances and attributes
aws ec2 describe-instances --query 'Reservations[].Instances[].[InstanceId, PublicIpAddress, PrivateIpAddress]' --output text
aws rds describe-db-instances --query "DBInstances[].[DBInstanceIdentifier, DeletionProtection]" --output text

# Same as above but filter by attribute, for example engine=mysql
aws rds describe-db-instances --filters Name=engine,Values=mysql --query "DBInstances[].[DBInstanceIdentifier, EngineVersion]"

######### More advanced filtering using jq #########

# Select only A-Alias records
aws route53 list-resource-record-sets --hosted-zone-id $ZONE |jq '.ResourceRecordSets[]| select(.AliasTarget)'

			`)
		},
	}

	return cmd
}
