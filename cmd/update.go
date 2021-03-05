package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates any executable defined as an URL attribute in .hooks.yaml.",
		Long:  "Rebuilds the hooks as defined in the .hooks.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("UUpdating executables..")
			if isErrorBool(updateExecutables(), "[ERROR]") {
			}
			fmt.Println("DONE")
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateExecutables() (err error) {
	fmt.Println("TODO:// implement")
	err = nil
	return
}
