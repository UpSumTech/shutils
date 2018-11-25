package usercmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands to handle user or group related operations`
	parseLongDesc  = `Prints examples of commands to handle user or group related operations`
	parseExample   = `
	### Available commands for user or group related operations
	shutils user`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "user [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# Scenario : lets say you added a temp user in a m/c to test stuff out and then want to cleanly remove it
useradd -m -G developer -g developer --s /bin/bash -c developer,,,, developer # This adds a user with a shell
cat /etc/group | grep sudo || groupadd sudo # Add the sudo group if it already isnt there
usermod -a -G sudo developer # Modify the user and add him to the sudo group
deluser developer --remove-all-files # delete the user with all his files
ps U developer # Kill all processes belonging to a deleted user
slay -clean developer # Kills all remaining processes from the above step
groupdel sudo # To delete a group

echo 'developer ALL=(ALL) NOPASSWD:/usr/bin/top' >> /etc/sudoers # To only allow a user to execute a single command
			`)
		},
	}

	return cmd
}
