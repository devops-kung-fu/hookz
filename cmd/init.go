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
			existingHookz := lib.NewOsFs().HasExistingHookz()
			if existingHookz {
				fmt.Println("Existing hookz files detected")
				fmt.Println("\nDid you mean to reset?")
				fmt.Println("        hookz reset [--verbose] [--debug]")
				fmt.Println("\nRun 'hookz --help' for usage.")
				fmt.Println()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			f := lib.NewOsFs()
			color.Style{color.FgLightBlue, color.OpBold}.Println("Initializing Hooks")
			fmt.Println()
			config, err := f.ReadConfig(version)
			if lib.IsErrorBool(err, "[ERROR]") {
				return
			}
			if lib.IsErrorBool(f.WriteHooks(config, verbose, debug), "[ERROR]") {
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
