package cmd

import (
	"github.com/devops-kung-fu/common/util"
	"github.com/spf13/cobra"

	"github.com/devops-kung-fu/hookz/lib"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates any defined sources and/or executable defined as an URL attribute in .hookz.yaml.",
		Long:  "Updates any defined sources and/or executable defined as an URL attribute in .hookz.yaml.",
		Run: func(cmd *cobra.Command, args []string) {
			util.DoIf(Verbose, func() {
				util.PrintInfo("Updating sources and executables")
			})
			config := CheckConfig()
			_ = InstallSources(config.Sources)
			if util.IsErrorBool(lib.UpdateExecutables(Afs, config)) {
				return
			}
			util.DoIf(Verbose, func() {
				util.PrintSuccess("Done!")
			})
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
}
