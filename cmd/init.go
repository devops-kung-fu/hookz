package cmd

import (
	"fmt"
	"os"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:     "initialize",
		Aliases: []string{"init"},
		Short:   "Initializes the hooks as defined in the .hookz.yaml file.",
		Long:    "Initializes the hooks as defined in the .hookz.yaml file.",
		PreRun: func(cmd *cobra.Command, args []string) {
			existingHookz := lib.HasExistingHookz(lib.NewOsFs())
			if existingHookz {
				fmt.Println("Existing hooks detected")
				fmt.Println("\nDid you mean to reset?")
				fmt.Println("        hookz reset [--verbose] [--debug]")
				fmt.Println("\nRun 'hookz --help' for usage.")
				fmt.Println()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			fs := lib.NewOsFs()
			color.Style{color.FgLightBlue, color.OpBold}.Println("Initializing Hooks")
			fmt.Println()
			config, err := lib.ReadConfig(fs, version)
			if err != nil && err.Error() == "NO_CONFIG" {
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
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "If true, output from commands is displayed when the hook executes.")
}
