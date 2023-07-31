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
nc -zv <service-fqdn> 8080 # To verify that remote service is accepting connections when telnet is not available

# test connectivity without telnet or nc installed
timeout 1 bash -c '</dev/tcp/google.com/443 && echo Port is open || echo Port is closed' || echo Connection timeout

# Find my own public ip
curl icanhazip.com # find external ip of machine easily
dig myip.opendns.com # find external ip via DNS query, helps if curl/wget not working

######### DNS troubleshooting ################
dig rainandrhyme.com # get DNS records for a domain
dig @8.8.8.8 rainandrhyme.com # get DNS records using a google server
dig +noall +answer @8.8.8.8 rainandrhyme.com # only get the answer to the DNS resolution and not the query parts
dig +trace @8.8.8.8 rainandrhyme.com # trace recursively how the DNS is getting resolved
dig @localhost rainandrhyme.com # get DNS records using the local DNS server you are running something like dnsmasq
dig +nocmd +noall +answer A foo.bar.com
getent hosts rainandrhyme.com # check if you have a DNS entry in your hosts file

# normal dns query with doggo
doggo github.com
doggo MX github.com @9.9.9.9
doggo MX github.com @tcp://1.1.1.1:53
# If you want to perform DOT (dns over tls) then use the proper port, ie 853 and mention that in the transport type
doggo MX github.com @tls://1.1.1.1:853
# This is an example for DOH(dns over https) dns query for cloudflare DOH server
doggo archive.org @https://cloudflare-dns.com/dns-query
doggo archive.org @https://cloudflare-dns.com/dns-query --json
doggo archive.org @https://cloudflare-dns.com/dns-query --json --debug
doggo duckduckgo.com --time

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
netstat -ptuwl --numeric-ports # Get all the inbound connections
netstat -pte -W --numeric-ports # Get all the outbound connections

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

# the -q flag is used for quiet output so that you are not flooded with too much data
ngrep '' udp # print udp packets
ngrep -q 'HTTP' 'udp' # print packets with header matching the string HTTP sent with UDP
ngrep -d eth0 port 8080 # print packets for port 8080 on eth0
ngrep -d any not port 22 # print packets for all traffic except ssh traffic
ngrep -d eth0 "example-.*.com" port 8080 # search for a string in the packets
ngrep -d any -i "user|PASS" port 8080 # search for a case insensitive regex in the packets
ngrep -d any -W byline "<search-string>" port 8080 # use the byline option for more readable input
ngrep -q 'HTTP' host 10.20.11.15 # print packets with header matching the string HTTP sent to/from the host
ngrep -q 'HTTP' src host 10.20.21.55 # print packets with header matching the string HTTP sent from the source host
ngrep -q 'HTTP' dest host 10.20.19.51 # print packets with header matching the string HTTP sent to the destination host
ngrep -q 'HTTP' dest host 10.20.19.51 # print packets with header matching the string HTTP sent to the destination host

# If you want to capture or read a pcap file then you can use these commands
# The -t option to ncap captures timestamps. -O is to write, -I is to read.
ngrep -O network_capture.pcap -qt 'HTTP'
ngrep -I network_capture.pcap -qt 'HTTP'

# Dont run nmap in aws infra, you will get banned very quickly.
nmap -T4 -F 198.10.100.0/24 # Scanning a large network for open ports
nmap -T4 -Pn -F 198.10.100.21 # Checks with ping if host is up. Host could be behind a firewall
nmap -T4 -Pn -F 198.10.100.21 --traceroute # trace the path to host along with scanning open ports
nmap -Pn -p 22 198.10.100.21 # Scan port 22 for the given host

# nping is a pretty useful utility to test network connectivity from machines
# For most of the commands though you will need root privileges
sudo nping --icmp -c 2  google.com # Can you ping the domain and get back a response for ICMP
nping --tcp-connect -c 2 -p 80 google.com # can you connect with tcp to a specific port
sudo nping --tcp -c 2 --flags S -p 80 google.com # you can also send a tcp request with a specific flag like SYN
sudo nping --udp -c 2 -p 40125 1.1.1.1 # You can try and connect with udp to a server and port
# You can echo client to nmap.org. They maintain a echo server application.
# You can use this to check if the is a NAT between you and the internet.
sudo nping -c 1 --echo-client "public" scanme.nmap.org

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

# To easily send and receive files using nc
tar cf - * | nc <ip_of_receiving_host> <port> # Tar and send files to a specific port on another host from a machine
nc -l -p <port> | tar x # Untar by listening to a port on the target machine

# To send and receive LVM files over the network
dd if=/dev/mapper/foo bs=4M | nc <ip_of_receiving_host> <port> # Send block storage files to a specific port on another host from a machine
nc -l <port> | dd of=/dev/mapper/foo bs=4M # Receive block storage file on the target machine

# quickly reasoning about CIDRs
ipcalc 172.16.1.0/24 -s 15 15 # gives you detailed info to partition a network with 2 subnets of size 15 each
echo "ibase=A;obase=2;248" | bc # quickly does conversions for you to understand network and host bits faster. This one is for decimal to binary.
echo "ibase=2;obase=A;11111000" | bc # quickly convert binary to decimal. Again, easy to convert CIDR to decimal.
prips 10.20.10.0/24 # Print out all the ip addresses in a network range
# This command below is very useful to quickly scan the network in a CIDR range
prips 130.229.16.0/20 | parallel --timeout 2 -j0 'ping -c 1 {} >/dev/null && echo {}' 2>/dev/null

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
