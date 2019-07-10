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

############# Inside psql console ##############
# Get all DB parameters in postgres
select name, setting, boot_val, reset_val, unit from pg_settings order by name;

# To see postgres extensions installed
show shared_preload_libraries;

# To show RDS plugins
SHOW rds.extensions;

# To show version in postgres
select version();

# To show current activity in a postgres database
select * from pg_stat_activity;

# To show different kinds of grants in a postgres database
select * from information_schema.role_table_grants where grantee = 'your_user';
select * from information_schema.role_column_grants where grantee = 'your_user';
select * from information_schema.role_routine_grants where grantee = 'your_user';

# In mysql innodb to show metrics and stats
use information_schema;
select name, subsystem, count, type, comment from INNODB_METRICS where status = 'enabled';

# In mysql innodb to show details on columns in a table
use information_schema;
select column_name, column_default, is_nullable, data_type, character_maximum_length,column_type, column_key, extra from COLUMNS where table_name = "<table_name>" and table_schema = "<db_name>";

# In mysql get constraints for databases
use information_schema;
select * from table_constraints where table_name = "<table_name>" and table_schema = "<db_name>";

# In mysql get stats for table
use information_schema;
select * from statistics where table_name = "<table_name>" and table_schema = "<db_name>";

# In RDS to kill a session that is running on a particular thread id
CALL mysql.rds_kill(<thread_id_from_processlist>);

# In RDS to kill a query that is running on a particular thread id
CALL mysql.rds_kill_query(<thread_id_from_processlist>);
			`)
		},
	}

	return cmd
}
