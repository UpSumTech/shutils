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
useradd -m -G developer -g developer --s /bin/bash -c developer,,,, developer # This adds a user with a shell

cat /etc/group | grep sudo || groupadd sudo # Add the sudo group if it already isnt there

usermod -a -G sudo developer # Modify the user and add him to the sudo group

deluser developer --remove-all-files # delete the user with all his files

ps U developer # Kill all processes belonging to a deleted user

slay -clean developer # Kills all remaining processes from the above step

groupdel sudo # To delete a group

echo 'developer ALL=(ALL) NOPASSWD:/usr/bin/top' >> /etc/sudoers # To only allow a user to execute a single command

awk -F':' '{ print $1}' /etc/passwd # To list all users in a system
getent passwd # To list all users in /etc/passwd

# To get all login users
uid_min=$(grep "^UID_MIN" /etc/login.defs)
awk -F':' -v "limit=${uid_min##UID_MIN}" '{ if ( $3 >= limit ) print $1}' /etc/passwd
			`)
		},
	}

	return cmd
}
