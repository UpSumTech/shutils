package main

import (
	"github.com/spf13/cobra"
	"github.com/sumanmukherjee03/shutils/cmd/dbcmds"
	"github.com/sumanmukherjee03/shutils/cmd/diskcmds"
	"github.com/sumanmukherjee03/shutils/cmd/dockercmds"
	"github.com/sumanmukherjee03/shutils/cmd/ec2cmds"
	"github.com/sumanmukherjee03/shutils/cmd/filecmds"
	"github.com/sumanmukherjee03/shutils/cmd/kubecmds"
	"github.com/sumanmukherjee03/shutils/cmd/misccmds"
	"github.com/sumanmukherjee03/shutils/cmd/netcmds"
	"github.com/sumanmukherjee03/shutils/cmd/pkgcmds"
	"github.com/sumanmukherjee03/shutils/cmd/proccmds"
	"github.com/sumanmukherjee03/shutils/cmd/seccmds"
	"github.com/sumanmukherjee03/shutils/cmd/syshealthcmds"
	"github.com/sumanmukherjee03/shutils/cmd/usercmds"
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
	rootCmd.Execute()
}
