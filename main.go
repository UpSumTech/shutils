package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootShortDesc = "Shutils is a simple utilty to print use cases of various sysadmin CLI tools"
	rootLongDesc  = `Shutils is a flexible tool built in golang.
	It prints cheat sheets for various sysadmin CLI tools.`
	cfgFile string
	Dryrun  bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:              "shutils [sub]",
		Short:            rootShortDesc,
		Long:             rootLongDesc,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("TODO: example usage to be changed afterwards")
		},
	}

	rootCmd.Execute()
}
