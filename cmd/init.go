package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/devops-kung-fu/common/util"
	"github.com/gookit/color"
	"github.com/spf13/cobra"

	"github.com/devops-kung-fu/hookz/lib"
)

var (
	initCmd = &cobra.Command{
		Use:     "initialize",
		Aliases: []string{"init"},
		Short:   "Initializes the hooks as defined in the .hookz.yaml file.",
		Long:    "Initializes the hooks as defined in the .hookz.yaml file.",
		PreRun: func(cmd *cobra.Command, args []string) {
			existingHookz := lib.HasExistingHookz(Afs)
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
			color.Style{color.FgLightBlue, color.OpBold}.Println("Initializing Hooks")
			fmt.Println()
			config, err := lib.ReadConfig(Afs, version)
			if err != nil && err.Error() == "NO_CONFIG" {
				NoConfig()
				os.Exit(1)
			}
			if util.IsErrorBool(err, "[ERROR]") {
				return
			}
			err = lib.InstallSources(config.Sources)
			if err != nil {
				log.Println("There was a problem installing sources")
				log.Println(err)
			}
			if util.IsErrorBool(lib.WriteHooks(Afs, config, Verbose, debug), "[ERROR]") {
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
