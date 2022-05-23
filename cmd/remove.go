package cmd

import (
	"github.com/devops-kung-fu/common/util"
	"github.com/spf13/cobra"

	"github.com/devops-kung-fu/hookz/lib"
)

var (
	removeCmd = &cobra.Command{
		Use:     "remove",
		Aliases: []string{"delete"},
		Short:   "Removes the hooks as defined in the .hookz.yaml file and any generated scripts.",
		Long:    "Removes the hooks as defined in the .hookz.yaml file and any generated scripts.",
		Run: func(cmd *cobra.Command, args []string) {
			util.DoIf(Verbose, func() {
				util.PrintInfo("Removing hooks and supporting files...")
			})
			if util.IsErrorBool(lib.RemoveHooks(Afs, Verbose)) {
				return
			}
			util.DoIf(Verbose, func() {
				util.PrintSuccess("Done")
			})
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}
