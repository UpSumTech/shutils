package main

import (
	"github.com/UpSumTech/shutils/cmd/dbcmds"
	"github.com/UpSumTech/shutils/cmd/diskcmds"
	"github.com/UpSumTech/shutils/cmd/dockercmds"
	"github.com/UpSumTech/shutils/cmd/ec2cmds"
	"github.com/UpSumTech/shutils/cmd/filecmds"
	"github.com/UpSumTech/shutils/cmd/kubecmds"
	"github.com/UpSumTech/shutils/cmd/memcmds"
	"github.com/UpSumTech/shutils/cmd/misccmds"
	"github.com/UpSumTech/shutils/cmd/netcmds"
	"github.com/UpSumTech/shutils/cmd/pkgcmds"
	"github.com/UpSumTech/shutils/cmd/proccmds"
	"github.com/UpSumTech/shutils/cmd/rabbitmqcmds"
	"github.com/UpSumTech/shutils/cmd/seccmds"
	"github.com/UpSumTech/shutils/cmd/syshealthcmds"
	"github.com/UpSumTech/shutils/cmd/usercmds"
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
