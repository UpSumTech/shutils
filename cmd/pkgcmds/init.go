package pkgcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands to list packages and info about them on a distro`
	parseLongDesc  = `Prints examples of commands to list packages and info about them on a distro`
	parseExample   = `
	### Example commands for listing the package and distro info etc
	shutils pkg`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "pkg [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
sudo apt-get install moreutils # Just install moreutils and thank me later
sudo apt-get install iputils-arping # package to send ARP requests at the ethernet level
sudo apt-get install tcptraceroute # useful to trace route with tcp instead of icmp
sudo apt-get install procinfo # useful to read proc files for socket information
sudo apt-get install sockstat # useful to read socket information
sudo apt-get install nmap # useful for scanning and trouble shooting large networks
sudo apt-get install multitail # useful for tailing multiple logs in one terminal window

cat /etc/*-release # To know what flavour of linux you are running
uname -m # System architecture
arch # System architecture
lsb_release -a # List distro info
lsmod # Lsit all the kernel modules
dpkg -l # List all packages in debian systems
apt list --installed | grep ssl # List installed packages in ubuntu

# The commands below can help manage/update gpg keys used to verify package signatures in a debian system
# Get the list of gpg keys which are trusted and managed by the system from here
ls -lah /usr/share/keyrings
# If you want to add a new gpg key for package verification, then this is the command
curl -fsSL <gpg-key-url> | sudo gpg -o /usr/share/keyrings/<key-name>.gpg --dearmor
# The gpg key file should match the filename which is used to verify the signature of the package installed
cat /etc/apt/sources.list.d/<repo>
sudo apt-get update

# List the trusted keys in the keyring. But this command could be deprecated soon.
sudo apt-key list

lscpu # List cpu info
lspci -mm # List all PCI buses in the system in machine readable format

rpm -qa # List of all rpm packages on a RHEL system
rpm -qf /bin/echo # List the package the file originated from

# To get a quick summary of what this binary might be doing. If you suspect something you have never seen before
whatis nc

# To find out what package installed a file
dpkg -S $(command -v ab)
			`)
		},
	}

	return cmd
}
