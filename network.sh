#! /usr/bin/env bash

telnet github.com 22 # try to see if port is accepting connections on remote machine
curl icanhazip.com # find external ip of machine easily

dig rainandrhyme.com # get DNS records for a domain
dig @8.8.8.8 rainandrhyme.com # get DNS records using a google server
dig @localhost rainandrhyme.com # get DNS records using the local DNS server you are running something like dnsmasq
getent hosts rainandrhyme.com # check if you have a DNS entry in your hosts file

arp -a # ARP of router
arping -I eth0 10.23.11.101 # Ping the device at the ethernet layer
ping google.ca # Check if request is even going out
traceroute google.ca # Trace using icmp
tcptraceroute google.ca # Trace using tcp instead of icmp

ifconfig -a # Display all interfaces
ifconfig en0 # Display selected interface

netstat -r -f inet # Route table for inet address family
netstat -tup # List currently active connection to the system
netstat -tupl # List listening ports
netstat -anp --udp --tcp | grep LISTEN # List listening ports for tcp and udp connections

ifdata -e eth0; echo $? # Checks existence of interface and prints the exit status
ifdata -pa eth0 # Network address of interface
ifdata -pn eth0 # Netmask of interface
ifdata -pN eth0 # Network address of interface
ifdata -pb eth0 # Broadcast of interface
ifdata -p eth0 # Prints details of the interface
ifdata -si eth0 # Stats of interface for incoming requests
ifdata -so eth0 # Stats of interface for outgoing requests

lsof -i 4 -n -P # List of all open sockets for ipv4
lsof -P -iTCP -sTCP:LISTEN # List all TCP listening sockets

socklist # List of all open sockets
sockstat -p 22 # Get socket info for port 22

iptables -L -n -v # List all firewall settings
iptables -Z # Clear out the counter for packets and bytes in the INPUT, OUTPUT and FORWARD chains
iptables -L -t nat # Check the NAT status
cat /proc/sys/net/ipv4/ip_forward # Check if IP forwarding is on. Useful for NAT instances

nmap -T4 -F 198.10.100.0/24 # Scanning a large network for open ports
nmap -T4 -Pn -F 198.10.100.21 # Checks with ping if host is up. Host could be behind a firewall
nmap -T4 -Pn -F 198.10.100.21 --traceroute # trace the path to host along with scanning open ports
nmap -Pn -p 22 198.10.100.21 # Scan port 22 for the given host

# ssh tunnel into a remote server to use a service on a blocked port running on that server
ssh -f -L <high_localhost_port>:localhost:<servers_blocked_port> user@proxy_server -N
nc -z localhost <high_localhost_port> # To verify that the tunnel is working

# quickly reasoning about CIDRs
ipcalc 172.16.1.0/24 -s 15 15 # gives you detailed info to partition a network with 2 subnets of size 15 each
echo "ibase=A;obase=2;248" | bc # quickly does conversions for you to understand network and host bits faster
echo "ibase=2;obase=A;11111000" | bc # quickly convert binary to decimal. Again, easy to convert CIDR to decimal

# quick proxy server listening on port 9999 and forwarding all requests to sectools.org
mkfifo response_pipe && nc -l 9999  0<response_pipe | nc sectools.org 80 1>response_pipe
