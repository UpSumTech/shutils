package diskcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands to operate on the disk`
	parseLongDesc  = `Prints examples of commands to operate on the disk`
	parseExample   = `
	### Available commands for operating on the disk
	shutils disk`
)

// Init instantiates the disk commands
func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "disk [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
ls -lah # Display all files in current dir in human readable format
ls -ltr # Display files newest last
du -sh ~/**/* | sort -rn # Display dirs in sorted sizes
df -h # Display free disk space in human readable format
cat /proc/partitions # Display partitions
lsblk -f # Display block storage devices as a tree view. Should list the same partitions above.
mount | column -t # Get all mounted filesystems
cat /etc/fstab # Get the static file system info
free -m # Get free memory statistics for the system
cat /etc/fstab # Get the information on mounted volumes
file -s /dev/xvda # Check whether the mounted device has any data or not

# Process to obtain device name where ebs volume is attached
vol_id="<ebs_volume_id>"
vol_id_on_disk="$(echo "$vol_id" | tr -d '-')"
device_id="$(ls /dev/disk/by-id/*-Amazon_Elastic_Block_Store_$vol_id_on_disk | head -1)"
device_name="/dev/$(readlink "$device_id" | tr / '\n' | tail -1)"

blkid -o value -s TYPE $device_name || mkfs -t ext4 "$device_name" # Check if volume is already formatted or not. If not format it to ext4
mkdir -p "<dir_you_want_to_mount_volume_to>" && mount "$dev_name" "<dir_you_want_to_mount_volume_to>" # mount ebs volume in a ec2 machine
			`)
		},
	}

	return cmd
}
