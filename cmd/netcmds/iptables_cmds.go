package netcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func IptablesCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iptables [no options!]",
		Short: `Iptables command examples`,
		Long:  `Iptables command examples`,
		Example: `
### Example commands for iptables
shutils net iptables
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
######### iptables or iproute commands #########

# Firewalling with iptables
apt-get install iptables-persistent && netfilter-persistent save # For persisting iptable changes across reboots on ubuntu
service iptables save # For persisting iptables on rhel

iptables -S # List all active iptable rules
iptables -S FORWARD # List all forward rules
iptables -L # List rules by chain
iptables -L -n -v # List all firewall settings
iptables -L FORWARD # List rules for the FORWARD chain
iptables -L -t nat # List the routing policies for NAT table

# Displaying rules of all the different tables of iproute
iptables -t filter -vL # Show rules of the filter table
iptables -t nat -vL # Show rules of the nat table
iptables -t mangle -vL # Show rules of the mangle table
iptables -t raw -vL # Show rules of the raw table
iptables -t security -vL # Show rules of the security table

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
			`)
		},
	}

	return cmd
}
