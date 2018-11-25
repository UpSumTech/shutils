package dbcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands for databases`
	parseLongDesc  = `Prints examples of commands for databases`
	parseExample   = `
	### Example commands for databases
	shutils db`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "db [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
export PGPASSWORD=<password>
psql -h <postgres_host> -p 5439 -U <user> -d <db_name> -c "ALTER TABLE <table_name> ADD COLUMN <col_name> NUMERIC(10,6) ENCODING runlength DEFAULT 0.0 NOT NULL"
psql -h <postgres_host> -p 5439 -U <user> -d <db_name> -a -w -f up_impression_agg_facts.sql

# ssh tunnel into a database server through a proxy server
ssh -f -L <high_localhost_port>:<database_server_host>:<database_server_port> user@proxy_server -N
pgcli -h localhost -p high_localhost_port -U database_user database_name
			`)
		},
	}

	return cmd
}
