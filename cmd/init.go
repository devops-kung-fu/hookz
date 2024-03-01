package cmd

import (
	"fmt"
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
			config, err := CheckConfig(Afs)
			if err != nil {
				if err != nil && err.Error() == "NO_CONFIG" {
					NoConfig()
				} else {
					util.PrintErr(err)
				}
				os.Exit(1)
			}
			_ = InstallSources(config.Sources)
			if util.IsErrorBool(lib.WriteHooks(Afs, config, Verbose, VerboseOutput)) {
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
