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

// Init - initiates db commands
func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "db [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
#######################################################
################### Postgres commands #################
#######################################################
export PGPASSWORD=<password>
psql -h <postgres_host> -p 5439 -U <user> -d <db_name> -c "ALTER TABLE <table_name> ADD COLUMN <col_name> NUMERIC(10,6) ENCODING runlength DEFAULT 0.0 NOT NULL"
psql -h <postgres_host> -p 5439 -U <user> -d <db_name> -a -w -f up_impression_agg_facts.sql

# ssh tunnel into a database server through a proxy server
ssh -f -L <high_localhost_port>:<database_server_host>:<database_server_port> user@proxy_server -N
pgcli -h localhost -p high_localhost_port -U database_user database_name

# Get all DB parameters in postgres
select name, setting, boot_val, reset_val, unit from pg_settings order by name;

# To see postgres extensions installed
show shared_preload_libraries;

# To show RDS plugins
SHOW rds.extensions;

# To show shared libraries
SHOW shared_preload_libraries;

# To see replication lag between master and a replica
SELECT extract(epoch from now() - pg_last_xact_replay_timestamp()) AS slave_lag

# To display version/time etc general information in postgres
select version();
select CURRENT_DATE;
select CURRENT_TIME;
select CURRENT_TIME - interval '1 hour';

select count(*) from <table_name> where created_on > (now()::timestamp without time zone - interval '10 minute');
select source,name,setting,boot_val,reset_val from pg_settings where name != 'rds.extensions' and boot_val != reset_val order by name;

# check to see if a database is running aws aurora
select AURORA_VERSION();

# from the WAL log publishing side what is the current replay_lsn
select * from pg_stat_replication;

# Lookup all replication slots
select * from pg_replication_slots;
# Lookup the current checkpoint
select pg_current_wal_lsn();
# Lookup the restart_lsn, and confirmed_flush_lsn for the slots
# lsn is generally like this current > last confirmed flush > last restart checkpoint
select slot_name, restart_lsn, confirmed_flush_lsn from pg_replication_slots;

# get replication slot lag in a human readable form
select slot_name, pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_lsn(),restart_lsn)) as replicationSlotLag, active from pg_replication_slots ;

# drop a db replication slot if necessary
select pg_drop_replication_slot('<slotname>');

# To show current activity in a postgres database
select * from pg_stat_activity;
select * from pg_stat_activity where usename like '<username>%';

# To show how much IO happens on the different indexes in postgres
select * from pg_statio_all_indexes where schemaname = 'public';

# To show different kinds of grants in a postgres database
select * from information_schema.role_table_grants where grantee = 'your_user';
select * from information_schema.role_column_grants where grantee = 'your_user';
select * from information_schema.role_routine_grants where grantee = 'your_user';

# Show all postgres users
select * from pg_user;

# Show all postgres groups
select * from pg_group;

# Show all postgres databases
select * from pg_database;

# Show postgres namespace named
select * from pg_namespace where nspname = '<your pg namespace>';

# To kill a postgres RDS backend proc
select pg_terminate_backend(<pid>);






#######################################################
###################### MYSQL ##########################
#######################################################
# In mysql innodb to show metrics and stats
use information_schema;
select name, subsystem, count, type, comment from INNODB_METRICS where status = 'enabled';

# In mysql to show global timeout variables
show global variables like '%timeout%';
show global variables like 'performance_schema';

# In mysql innodb to show details on columns in a table
use information_schema;
select column_name, column_default, is_nullable, data_type, character_maximum_length,column_type, column_key, extra from COLUMNS where table_name = "<table_name>" and table_schema = "<db_name>";

# In mysql get constraints for databases
use information_schema;
select * from table_constraints where table_name = "<table_name>" and table_schema = "<db_name>";

# In mysql get stats for table
use information_schema;
select * from statistics where table_name = "<table_name>" and table_schema = "<db_name>";

# In mysql to get all constraints on table in db referencing other tables
use information_schema;
select table_name, column_name, constraint_name, referenced_table_name, referenced_column_name from key_column_usage where table_schema = "<db_name>" and table_name = "<table_name>" and referenced_column_name is not null;

# In mysql get all constraints referencing a particular table in db
use information_schema;
select table_name, column_name, constraint_name, referenced_table_name, referenced_column_name from key_column_usage where table_schema = "<db_name>" and referenced_table_name = "<table_name>" and referenced_column_name is not null;

# In mysql get tables affected by on delete cascade
use information_schema;
select table_name from referential_constraints where constraint_schema = '<db_name>' and referenced_table_name = '<table_name_where_clause_applied>' and delete_rule = 'cascade';

# Disable all table constraints
alter table mytable nocheck constraint all

# Enable all table constraints
alter table mytable with check check constraint all

# Disable single constraint
alter table mytable nocheck constraint myconstraint

# Enable single constraint
alter table mytable with check check constraint myconstraint

# To run a single command for mysql using the cli
mycli -h localhost -u <user> -p<password> -e "show processlist;"
mycli -h localhost -u <root_user> -p<root_password> -e "show full processlist;"

# To see processes that are not administrative
select * from processlist where User not in ('rdsadmin');

# To see processes that are for a specific db
select * from processlist where db = '<db_name>';

# To see transactions that are creating locks
use information_schema;
select * from innodb_trx;
select * from innodb_locks;
select * from innodb_lock_waits;
select * from innodb_locks where lock_table = <db_name>.<table_name>;
select innodb_locks.* from innodb_locks join innodb_lock_waits on (innodb_locks.lock_trx_id = innodb_lock_waits.blocking_trx_id);
select trx_id, trx_requested_lock_id, trx_mysql_thread_id, trx_query from innodb_trx where trx_state = 'lock wait';
# To see which transactions are waiting and which transactions are blocking the waiting transactions
select r.trx_id waiting_trx_id, r.trx_mysql_thread_id waiting_thread, r.trx_query waiting_query, b.trx_id blocking_trx_id, b.trx_mysql_thread_id blocking_thread, b.trx_query blocking_query from innodb_lock_waits w inner join innodb_trx b on b.trx_id = w.blocking_trx_id inner join innodb_trx r on r.trx_id = w.requesting_trx_id;

# To troubleshoot replication lag and compare master log file and replicas log files - especially file processed by IO_THREAD and file processed by SQL_THREAD
# Run this in the master instance
show master status;
# Run this in the slave instance
show slave status;

# In mysql to see slave status on a replica
show slave status;
# In mysql rds master to add a replication user
grant replication slave, replication client on *.* to <replication_user>@'%';
# In mysql if slave (say an on-prem DB which is the slave) is configured properly to replicate through it's my.cnf, then add the settings to replicate
CHANGE MASTER TO MASTER_HOST='<rds-endpoint>',MASTER_USER='<replication-user-in-master>',MASTER_PASSWORD='<passwd>',MASTER_LOG_FILE='<bin_log_file_name>',MASTER_LOG_POS=<bin_log_file_position>
# In mysql (on the on-prem DB) to start replication in a replica
start slave;

# In RDS slave to stop replication
CALL mysql.rds_stop_replication;

# In RDS to kill a session that is running on a particular thread id
CALL mysql.rds_kill(<thread_id_from_processlist>);

# In RDS to kill a query that is running on a particular thread id
CALL mysql.rds_kill_query(<thread_id_from_processlist>);
			`)
		},
	}

	return cmd
}
