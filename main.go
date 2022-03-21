package main

import (
	"shutils/cmd/dbcmds"
	"shutils/cmd/diskcmds"
	"shutils/cmd/dockercmds"
	"shutils/cmd/ec2cmds"
	"shutils/cmd/filecmds"
	"shutils/cmd/kubecmds"
	"shutils/cmd/memcmds"
	"shutils/cmd/misccmds"
	"shutils/cmd/netcmds"
	"shutils/cmd/pkgcmds"
	"shutils/cmd/proccmds"
	"shutils/cmd/rabbitmqcmds"
	"shutils/cmd/seccmds"
	"shutils/cmd/syshealthcmds"
	"shutils/cmd/usercmds"

	"github.com/spf13/cobra"
)

var (
	rootShortDesc = "Shutils is a simple utilty to print use cases of various sysadmin CLI tools"
	rootLongDesc  = `Shutils is a flexible tool built in golang.
	It prints cheat sheets for various sysadmin CLI tools.`
)

func main() {
	var rootCmd = &cobra.Command{
		Use:              "shutils [sub]",
		Short:            rootShortDesc,
		Long:             rootLongDesc,
		TraverseChildren: true,
	}

	rootCmd.AddCommand(diskcmds.Init())
	rootCmd.AddCommand(kubecmds.Init())
	rootCmd.AddCommand(dockercmds.Init())
	rootCmd.AddCommand(filecmds.Init())
	rootCmd.AddCommand(usercmds.Init())
	rootCmd.AddCommand(syshealthcmds.Init())
	rootCmd.AddCommand(pkgcmds.Init())
	rootCmd.AddCommand(dbcmds.Init())
	rootCmd.AddCommand(netcmds.Init())
	rootCmd.AddCommand(proccmds.Init())
	rootCmd.AddCommand(seccmds.Init())
	rootCmd.AddCommand(ec2cmds.Init())
	rootCmd.AddCommand(misccmds.Init())
	rootCmd.AddCommand(rabbitmqcmds.Init())
	rootCmd.AddCommand(memcmds.Init())
	rootCmd.Execute()
}
