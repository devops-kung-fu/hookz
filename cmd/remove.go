package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Removes the hooks as defined in the .hooks.yaml file.",
		Long:  "Removes the hooks as defined in the .hooks.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Removing git hooks...")
			if isErrorBool(removeHooks(), "[ERROR]") {
				return
			}
			fmt.Println("DONE")
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}
