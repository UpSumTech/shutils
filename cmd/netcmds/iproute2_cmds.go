package netcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Iproute2Cmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iproute2 [no options!]",
		Short: `Iproute2 command examples`,
		Long:  `Iproute2 command examples`,
		Example: `
### Example commands for iproute2
shutils net iproute2
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
######### iproute2 commands #########

ip addr show docker0 # Show the address of docker0 interface
ip addr show up # Show up and running interfaces

ip addr add 192.20.10.1/24 dev eth0 # Add a specific address to an interface
ip addr delete 192.20.10.1/24 dev eth0 # Remove an address from an interface
ip addr flush dev eth0 # Flush all addresses from an interface

ip link show dev eth0 # Shows the link (interface) for eth0
ip link set dev eth0 up # Turns up an interface
ip link set dev eth0 down # Turns down an interface
ip link set dev eth0 mtu 1200 # Control the mtu of an interface if your ethernet supports it

ip route show # Display the routes on a machine
ip route show to match 192.168.0.1/24 # Display routes matching subnet and all larger subnets

ip route add 0.0.0.0/0 via 192.0.2.1 # To add a default route
ip route add 192.0.2.128/24 via 192.0.2.1 # Add a new route via a gateway
ip route change 192.168.2.0/24 via 10.0.0.1 # To change a route to use a different gateway
ip route show cached # Show the cached routes

ip route add unreachable 192.0.2.128/24 # Returns "unreachable" for ICMP requests to the client
ip route add prohibit 192.0.2.128/24 # Returns "prohibited" for ICMP requests to the client

ip netns list # List the network namespaces
ip netns exec red /bin/bash # Start a bash shell in the red network namespace
ip netns pids red # List all processes in the red network namespace
ip netns identify 7000 # Identify the network namespace of the PID

ip monitor # Monitor network events like link or addr changes or routing table changes etc
ip monitor route # Monitor route table changes

ip netconf show # View sysctl config on the machine
ip netconf show dev docker0 # View sysctl config on the machine for a specific device
			`)
		},
	}

	return cmd
}
