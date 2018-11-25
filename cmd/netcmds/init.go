package netcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of debugging network`
	parseLongDesc  = `Prints examples of debugging network`
	parseExample   = `
	### Example commands for debugging network
	shutils db`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "net [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
telnet github.com 22 # try to see if port is accepting connections on remote machine
curl icanhazip.com # find external ip of machine easily

dig rainandrhyme.com # get DNS records for a domain
dig @8.8.8.8 rainandrhyme.com # get DNS records using a google server
dig +noall +answer @8.8.8.8 rainandrhyme.com # only get the answer to the DNS resolution and not the query parts
dig @localhost rainandrhyme.com # get DNS records using the local DNS server you are running something like dnsmasq
getent hosts rainandrhyme.com # check if you have a DNS entry in your hosts file

arp -a # ARP of router
arping -I eth0 10.23.11.101 # Ping the device at the ethernet layer
ping google.ca # Check if request is even going out
traceroute google.ca # Trace using icmp
tcptraceroute google.ca # Trace using tcp instead of icmp

ifconfig -a # Display all interfaces
ifconfig en0 # Display selected interface

netstat -r inet # Route table for DARPA internet address family
netstat -tup # List currently active connection to the system
netstat -tupl # List listening ports
netstat -anp --udp --tcp | grep LISTEN # List listening ports for tcp and udp connections

ifdata -e eth0; echo $? # Checks existence of interface and prints the exit status
ifdata -pa eth0 # Network address of interface
ifdata -pn eth0 # Netmask of interface
ifdata -pN eth0 # Network address of gateway
ifdata -pb eth0 # Broadcast of interface
ifdata -p eth0 # Prints details of the interface
ifdata -si eth0 # Stats of interface for incoming requests
ifdata -so eth0 # Stats of interface for outgoing requests

lsof -i 4 -n -P # List of all open sockets for ipv4
lsof -P -iTCP -sTCP:LISTEN # List all TCP listening sockets

socklist # List of all open sockets
sockstat -p 22 # Get socket info for port 22
sockstat -cl -U 0 # Get all the connected sockets root is listening to

iptables -L -n -v # List all firewall settings
iptables -Z # Clear out the counter for packets and bytes in the INPUT, OUTPUT and FORWARD chains
iptables -L -t nat # Check the NAT status
cat /proc/sys/net/ipv4/ip_forward # Check if IP forwarding is on. Useful for NAT instances

nmap -T4 -F 198.10.100.0/24 # Scanning a large network for open ports
nmap -T4 -Pn -F 198.10.100.21 # Checks with ping if host is up. Host could be behind a firewall
nmap -T4 -Pn -F 198.10.100.21 --traceroute # trace the path to host along with scanning open ports
nmap -Pn -p 22 198.10.100.21 # Scan port 22 for the given host

# Exec ssh commands on a machine and get it's output
ssh -o ExitOnForwardFailure=yes foo.example.com -t "ps -elf"

# ssh through a jump host with a connection timeout of 10s
ssh -A -J jump.example.com -i ~/.ssh/priv_key_for_jumphost.pem ubuntu@dev.example.com -o ConnectTimeout=10

# ssh tunnel into a remote server to use a service on a blocked port running on that server
ssh -f -L <high_localhost_port>:localhost:<servers_blocked_port> user@proxy_server -N
nc -z localhost <high_localhost_port> # To verify that the tunnel is working

# To send and receive files using nc
tar cf - * | netcat <ip_of_receiving_host> <port> # Tar and send files to a specific port on another host from a machine
netcat -l -p <port> | tar x # Untar by listening to a port on the target machine

# To send and receive LVM files over the network
dd if=/dev/mapper/foo bs=4M | netcat <ip_of_receiving_host> <port> # Send block storage files to a specific port on another host from a machine
nc -l <port> | dd of=/dev/mapper/foo bs=4M # Receive block storage file on the target machine

# quickly reasoning about CIDRs
ipcalc 172.16.1.0/24 -s 15 15 # gives you detailed info to partition a network with 2 subnets of size 15 each
echo "ibase=A;obase=2;248" | bc # quickly does conversions for you to understand network and host bits faster
echo "ibase=2;obase=A;11111000" | bc # quickly convert binary to decimal. Again, easy to convert CIDR to decimal

# quick proxy server listening on port 9999 and forwarding all requests to sectools.org
mkfifo response_pipe && nc -l 9999  0<response_pipe | nc sectools.org 80 1>response_pipe

# Kill process to free up a port
fuser -n tcp 8080 # Identify a process using a port
fuser -k -n tcp 8080 # To free up the 8080 tcp port by killing the process running it


# Firewall in ubuntu
ufw allow 22 # Allow ssh port if firewall is there
ufw deny 22 # Deny ssh port if firewall is there


######### iproute commands #########

# Firewalling with iptables
apt-get install iptables-persistent && netfilter-persistent save # For persisting iptable changes across reboots on ubuntu
service iptables save # For persisting iptables on rhel

iptables -S # List all active iptable rules
iptables -S FORWARD # List all forward rules
iptables -L # List rules by chain
iptables -L FORWARD # List rules for the FORWARD chain
iptables -L -t nat # List the routing policies for NAT table

iptables -Z # Clear the counters for all the chains
iptables -Z INPUT # Clear the counters for the INPUT chain

iptables -D INPUT 3 # To delete the third rule from the INPUT chain based on the output you saw from the list command
iptables -D INPUT -m conntrack --ctstate INVALID -j DROP # To drop a specific rule (everything coming after '-D INPUT')


# Recreate firewall by flushing and recreating everything
iptables -P INPUT ACCEPT # Set the default policy to ACCEPT for the INPUT chain to allow access to the machine and prevent being locked out of ssh
iptables -P FORWARD ACCEPT # Set the default policy to ACCEPT for the FORWARD chain to allow access to the machine and prevent being locked out of ssh
iptables -P OUTPUT ACCEPT # Set the default policy to ACCEPT for the OUTPUT chain to allow access to the machine and prevent being locked out of ssh

iptables -t nat -F # Flush the nat table
iptables -t mangle -F # Flush the mangle table

iptables -F # Flush all chains
iptables -F INPUT # Flush the INPUT chain
iptables -X # Delete all non-default chains like DOCKER, DOCKER-ISOLATION etc

# Enable the loopback interface for the INPUT and OUTPUT chain
iptables -A INPUT -i lo -j ACCEPT
iptables -A OUTPUT -o lo -j ACCEPT

# Enable established connections
iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
iptables -A OUTPUT -m conntrack --ctstate ESTABLISHED -j ACCEPT

# Forward traffic from internal eth1 interface to external eth0 interface
iptables -A FORWARD -i eth1 -o eth0 -j ACCEPT

# Blocking traffic
iptables -A INPUT -m conntrack --ctstate INVALID -j DROP # Drop incoming invalid packets
iptables -A INPUT -s 20.30.40.8 -j DROP # Block an ip from connecting
iptables -A INPUT -i eth0 -s 20.30.40.8 -j DROP # Block an ip from connecting to a specific interface
iptables -A OUTPUT -p tcp --dport 25 -j REJECT # Block outgoing SMTP traffic

# Allowing ssh and rsync traffic
iptables -A INPUT -p tcp -s 20.30.40.15/24 --dport 22 -m conntrack --ctstate NEW,ESTABLISHED -j ACCEPT # Allow incoming ssh traffic from a CIDR
iptables -A OUTPUT -p tcp --sport 22 -m conntrack --ctstate ESTABLISHED -j ACCEPT # Allow outgoing ssh traffic
iptables -A INPUT -p tcp -s 20.30.40.15/24 --dport 873 -m conntrack --ctstate NEW,ESTABLISHED -j ACCEPT # Allow incoming rsync traffic from a specific CIDR
iptables -A OUTPUT -p tcp --sport 873 -m conntrack --ctstate ESTABLISHED -j ACCEPT # Allow outgoing traffic for rsync

# Combining the ssh and rsync rules into one
iptables -A INPUT -p tcp -s 20.30.40.15/24 -m multiport --dport 22,873 -m conntrack --ctstate NEW,ESTABLISHED -j ACCEPT # Allow incoming rsync traffic from a specific CIDR
iptables -A OUTPUT -p tcp -m multiport --sport 22,873 -m conntrack --ctstate ESTABLISHED -j ACCEPT # Allow outgoing traffic for rsync

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
