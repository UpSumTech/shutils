package main

import (
	"github.com/spf13/cobra"
	"github.com/sumanmukherjee03/shutils/cmd/kubecmds"
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

	rootCmd.AddCommand(kubecmds.Init())
	rootCmd.Execute()
}
