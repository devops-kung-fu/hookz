package cmd

import (
	"fmt"
	"os"

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
			fs := lib.NewOsFs()
			color.Style{color.FgLightBlue, color.OpBold}.Println("Reset Hooks")
			fmt.Println()

			if lib.IsErrorBool(lib.RemoveHooks(fs, verbose), "[ERROR]") {
				return
			}
			config, err := lib.ReadConfig(fs, version)
			if err.Error() == "NO_CONFIG" {
				os.Exit(1)
			}
			if lib.IsErrorBool(err, "[ERROR]") {
				return
			}
			if lib.IsErrorBool(lib.WriteHooks(fs, config, verbose, debug), "[ERROR]") {
				return
			}
			color.Style{color.FgLightGreen}.Println("Done!")
		},
	}
)

func init() {
	rootCmd.AddCommand(resetCmd)
	resetCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "If true, output from commands is displayed when the hook executes.")
}
