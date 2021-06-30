package cmd

import (
	"fmt"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:     "remove",
		Aliases: []string{"delete"},
		Short:   "Removes the hooks as defined in the .hooks.yaml file and any generated scripts.",
		Long:    "Removes the hooks as defined in the .hooks.yaml file and any generated scripts.",
		Run: func(cmd *cobra.Command, args []string) {
			color.Style{color.FgLightBlue, color.OpBold}.Println("Removing Hooks")
			fmt.Println()
			if lib.IsErrorBool(lib.RemoveHooks(lib.NewOsFs(), verbose), "[ERROR]") {
				return
			}
			color.Style{color.FgLightGreen}.Println("Done!")
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}
