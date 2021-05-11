package cmd

import (
	"fmt"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates any executable defined as an URL attribute in .hooks.yaml.",
		Long:  "Rebuilds the hooks as defined in the .hooks.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			color.Style{color.FgLightBlue, color.OpBold}.Println("Update Executables")
			fmt.Println()
			config, err := lib.NewOsFs().ReadConfig(version)
			if lib.IsErrorBool(err, "[ERROR]") {
				return
			}
			if lib.IsErrorBool(lib.NewOsFs().UpdateExecutables(config), "[ERROR]") {
				return
			}
			color.Style{color.FgLightGreen}.Println("Done!")
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
}
