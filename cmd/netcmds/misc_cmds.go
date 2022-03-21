package netcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func MiscCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "misc [no options!]",
		Short: `Prints examples of misc cmds debugging network`,
		Long:  `Prints examples of misc cmds debugging network`,
		Example: `
### Misc example commands for debugging network
shutils net misc
		`,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
telnet github.com 22 # try to see if port is accepting connections on remote machine

# test connectivity without telnet or nc installed
timeout 1 bash -c '</dev/tcp/google.com/443 && echo Port is open || echo Port is closed' || echo Connection timeout

curl icanhazip.com # find external ip of machine easily
dig myip.opendns.com # find external ip via DNS query, helps if curl/wget not installed

dig rainandrhyme.com # get DNS records for a domain
dig @8.8.8.8 rainandrhyme.com # get DNS records using a google server
dig +noall +answer @8.8.8.8 rainandrhyme.com # only get the answer to the DNS resolution and not the query parts
dig +trace @8.8.8.8 rainandrhyme.com # trace recursively how the DNS is getting resolved
dig @localhost rainandrhyme.com # get DNS records using the local DNS server you are running something like dnsmasq
dig +nocmd +noall +answer A foo.bar.com
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

# To get the server fingerprint or public keys for adding to the known_hosts file of a client
ssh-keyscan -H gitlab.com

# Removes all keys belonging to a hostname from known_hosts
ssh-keygen -R <hostname>

# Generate public keypair for a given private key
ssh-keygen -y -f private.pem

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

######### tcpdump commands #########
# Get all https traffic without converting addresses to names with hex output
tcpdump -nnSX port 443

# Get traffic going to or from a particular host
tcpdump host 1.1.1.1

# Filter traffic by source or destination
tcpdump src 1.1.1.1
tcpdump dst 1.0.0.1

# Find packets going to/from a particular network
tcpdump net 172.31.0.0/16

# Verbose and human readable timestamp on any interface and all packets
tcpdump -vvv -tttt -i any -s0 -A -w <file_name>.dump

# Human readable timestamp on eth0 interface and for a set of hosts on http port
tcpdump -tttt 'host (182.101.17.43 or 182.101.17.42) and port 80' -i eth0 -s0

# Human readable timestamp for a particular network and with detailed hex output and only 10 packets
tcpdump -tttt -s0 -c 10 -xx -XX net 172.31.0.0/16

# Only display outgoing data packets to the destination IP on port 80 - no SYN, FIN, ACK packets
tcpdump -vvvv -tttt -s0 -A -X 'port 80 and dst 172.15.10.207 and (((ip[2:2] - ((ip[0]&0xf)<<2)) - ((tcp[12]&0xf0)>>2)) != 0)''>>))))'

# Copied from stack overflow - displays headers and packets in a better format for easier visualization
tcpdump -A -s 10240 'tcp port 8080 and (((ip[2:2] - ((ip[0]&0xf)<<2)) - ((tcp[12]&0xf0)>>2)) != 0)' | egrep --line-buffered "^........(GET |HTTP\/|POST |HEAD )|^[A-Za-z0-9-]+: " | sed -r 's/^........(GET |HTTP\/|POST |HEAD )/\n\1/g'>>))))'

######### tcpdump commands #########
# Capture traffic on ssh port and print them to the console
tcpflow -p -c -i eth0 port 22

######### WIRESHARK filter commands ##########

# Find all packets with src and dest ips or CIDRs
(ip.dst==52.7.68.129 || ip.dst==52.203.198.115) && (ip.src==172.17.0.0/16)

# Find all packets with src and dest ips or CIDRs that are experiencing conn resets
(ip.dst==52.7.68.129 || ip.dst==52.203.198.115) && (ip.src==172.17.0.0/16) && (tcp.flags.reset == 1)

# Follow a tcp stream
tcp.stream == 28
			`)
		},
	}

	return cmd
}
