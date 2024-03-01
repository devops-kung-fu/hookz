package cmd

import (
	"os"

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
			config, err := CheckConfig()
			if err != nil {
				if err != nil && err.Error() == "NO_CONFIG" {
					NoConfig()
				} else {
					util.PrintErr(err)
				}
				os.Exit(1)
			}
			_ = InstallSources(config.Sources)
			updateCount, _ := lib.UpdateExecutables(Afs, config)
			util.DoIf(updateCount == 0, func() {
				util.PrintInfo("Nothing to Update!")
			})
			util.DoIf(Verbose, func() {
				util.PrintSuccess("Done")
			})
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
}
