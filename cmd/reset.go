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
			f := lib.NewOsFs()
			color.Style{color.FgLightBlue, color.OpBold}.Println("Reset Hooks")
			fmt.Println()
			fmt.Println("[*] Removing existing hooks...")

			if lib.IsErrorBool(f.RemoveHooks(), "[ERROR]") {
				return
			}
			config, err := f.ReadConfig(version)
			if lib.IsErrorBool(err, "[ERROR]") {
				return
			}
			if lib.IsErrorBool(f.WriteHooks(config, debug), "[ERROR]") {
				return
			}
			color.Style{color.FgLightGreen}.Println("\nDone!")
		},
	}
)

func init() {
	rootCmd.AddCommand(resetCmd)
	resetCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "If true, output from commands is displayed when the hook executes.")
}
