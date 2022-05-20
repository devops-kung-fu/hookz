package cmd

import (
	"fmt"
	"log"

	"github.com/devops-kung-fu/common/util"
	"github.com/spf13/cobra"

	"github.com/devops-kung-fu/hookz/lib"
)

var (
	resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Rebuilds the hooks as defined in the .hookz.yaml file.",
		Long:  "Rebuilds the hooks as defined in the .hookz.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("Resetting Hooks")
			fmt.Println()

			if util.IsErrorBool(lib.RemoveHooks(Afs, Verbose), "[ERROR]") {
				return
			}
			config := CheckConfig()
			err := lib.InstallSources(config.Sources)
			if err != nil {
				util.PrintErr("There was a problem installing sources", err)
				log.Println(err)
			}
			if util.IsErrorBool(lib.WriteHooks(Afs, config, Verbose, VerboseOutput), "[ERROR]") {
				return
			}
			util.PrintSuccess("Done")
		},
	}
)

func init() {
	rootCmd.AddCommand(resetCmd)
}
