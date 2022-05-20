package cmd

import (
	"fmt"

	"github.com/devops-kung-fu/common/util"
	"github.com/gookit/color"
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
			color.Style{color.FgLightBlue, color.OpBold}.Println("Removing Hooks")
			fmt.Println()
			if util.IsErrorBool(lib.RemoveHooks(Afs, Verbose), "[ERROR]") {
				return
			}
			color.Style{color.FgLightGreen}.Println("Done!")
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}
