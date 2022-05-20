package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/devops-kung-fu/common/util"
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
				fmt.Println("        hookz reset [--verbose] [--debug] [--verbose-output")
				fmt.Println("\nRun 'hookz --help' for usage.")
				fmt.Println()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			util.DoIf(Verbose, func() {
				util.PrintInfo("Creating hooks")
			})
			config := CheckConfig()
			if len(config.Sources) > 0 && Verbose {
				util.DoIf(Verbose, func() {
					util.PrintInfo("Installing sources")
				})
			}
			err := lib.InstallSources(config.Sources)
			if err != nil {
				log.Println("There was a problem installing sources")
				log.Println(err)
				return
			}
			if util.IsErrorBool(lib.WriteHooks(Afs, config, Verbose, VerboseOutput), "[ERROR]") {
				return
			}
			util.DoIf(Verbose, func() {
				util.PrintSuccess("Done")
			})
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}
