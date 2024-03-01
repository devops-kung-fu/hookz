package cmd

import (
	"os"

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
			util.DoIf(Verbose, func() {
				util.PrintInfo("Resetting Hooks")
			})
			if util.IsErrorBool(lib.RemoveHooks(Afs, Verbose)) {
				return
			}
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
	rootCmd.AddCommand(resetCmd)
}
