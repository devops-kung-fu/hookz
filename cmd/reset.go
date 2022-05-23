package cmd

import (
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
			if util.IsErrorBool(lib.RemoveHooks(Afs, Verbose)) {
				return
			}
			config := CheckConfig()
			_ = InstallSources(config.Sources)
			if util.IsErrorBool(lib.WriteHooks(Afs, config, Verbose, VerboseOutput)) {
				return
			}
			util.PrintSuccess("Done")
		},
	}
)

func init() {
	rootCmd.AddCommand(resetCmd)
}
