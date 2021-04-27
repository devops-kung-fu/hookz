package cmd

import (
	"fmt"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Rebuilds the hooks as defined in the .hooks.yaml file.",
		Long:  "Rebuilds the hooks as defined in the .hooks.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			color.Style{color.FgGray, color.OpBold}.Println("Reset Hooks")
			fmt.Println()
			fmt.Println("[*] Removing existing hooks...")
			config, err := lib.ReadConfig(version)
			if lib.IsErrorBool(err, "[ERROR]") {
				return
			}
			if lib.IsErrorBool(lib.RemoveHooks(), "[ERROR]") {
				return
			}
			if lib.IsErrorBool(lib.WriteHooks(config, verbose), "[ERROR]") {
				return
			}
			fmt.Println("\nDONE!")
		},
	}
)

func init() {
	rootCmd.AddCommand(resetCmd)
	resetCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "If true, output from commands is displayed when the hook executes.")
}
