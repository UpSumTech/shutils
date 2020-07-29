package dockercmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func DockerDebugCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug [no options!]",
		Short: `docker debug commands`,
		Long:  `docker debug commands`,
		Example: `
### Example commands for debugging docker containers
shutils docker debug
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# Get the host pid of the running container
docker inspect <container_id> -f "{{.State.Pid}}"

# Get docker event stream
docker events --since '2020-07-28T22:30:00' --until '2020-07-28T23:00:00'

# Get all the env vars inside the container process namespace
cat /proc/<container_pid_on_host>/environ | sed -E 's#([A-Z_0-9]*)=([\s]*)#\n\1=\2#g'; echo
cat /proc/<container_pid_on_host>/cmdline | strings | xargs # regen the exact cmd that was used to run the container

# Get the inodes of the different namespaces the container is in
ls -lah /proc/<container_pid_on_host>/ns

# Get container process status
cat /proc/<container_pid_on_host>/status

# Outbound connections in container network namespace
nsenter -t <container_pid_on_host> -n netstat -pte -W --numeric-ports

# Inbound connections in container network namespace
nsenter -t <container_pid_on_host> -n netstat -ptuwl --numeric-ports

# Get ip address of docker container
nsenter -t <container_pid_on_host> -n ip addr show

# Open shell and go to container's working dir by entering container namespaces - pid, ipc, network, UTS
nsenter -t <container_pid_on_host> -n -i -u -p -w

# tcpdump inside the containers network namespace for udp traffic only to see DNS resolutions
nsenter -t <container_pid_on_host> -n tcpdump -i eth0 udp port 53

# tcpdump inside the containers network namespace for all traffic
nsenter -t <container_pid_on_host> -n tcpdump -i eth0 -A -s0

# tcpdump inside the containers network namespace for all incoming/outgoing traffic to mysql instance
nsenter -t <container_pid_on_host> tcpdump -i eth0 -A any host <mysql_host_or_ip> and port 3306
			`)
		},
	}

	return cmd
}
