package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Rebuilds the hooks as defined in the .hooks.yaml file.",
		Long:  "Rebuilds the hooks as defined in the .hooks.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			hookzHeader()
			fmt.Println("Resetting git hooks\n")
			fmt.Println("[*] Removing existing hooks...")
			if isErrorBool(removeHooks(), "[ERROR]") {
				return
			}
			if isErrorBool(writeHooks(), "[ERROR]") {
				return
			}
			fmt.Println("\nDONE!")
		},
	}
)

func init() {
	rootCmd.AddCommand(resetCmd)
	resetCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "If true, output from commands is displayed when the hook executes.")
}
