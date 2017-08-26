#! /usr/bin/env bash

arp -a # ARP of router
arping -I eth0 10.23.11.101 # Ping the device at the ethernet layer
ping google.ca # Check if request is even going out
traceroute google.ca # Trace using icmp
tcptraceroute google.ca # Trace using tcp instead of icmp

ifconfig -a # Display all interfaces
ifconfig en0 # Display selected interface

netstat -r -f inet # Route table for inet address family

ifdata -e eth0; echo $? # Checks existence of interface and prints the exit status
ifdata -pa eth0 # Network address of interface
ifdata -pn eth0 # Netmask of interface
ifdata -pN eth0 # Network address of interface
ifdata -pb eth0 # Broadcast of interface
ifdata -p eth0 # Prints details of the interface
ifdata -si eth0 # Stats of interface for incoming requests
ifdata -so eth0 # Stats of interface for outgoing requests

lsof -i 4 -n -P # List of all open sockets for ipv4
socklist # List of all open sockets
