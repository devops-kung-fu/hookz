package cmd

import (
	"github.com/devops-kung-fu/hookz/lib"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:     "initialize",
		Aliases: []string{"init"},
		Short:   "Initializes the hooks as defined in the .hookz.yaml file.",
		Long:    "Initializes the hooks as defined in the .hookz.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			color.Style{color.FgLightBlue, color.OpBold}.Println("Initializing Hooks")
			config, err := lib.ReadConfig(version)
			if lib.IsErrorBool(err, "[ERROR]") {
				return
			}
			if lib.IsErrorBool(lib.WriteHooks(config, verbose), "[ERROR]") {
				return
			}
			color.Style{color.FgLightGreen}.Println("\nDone!")
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "If true, output from commands is displayed when the hook executes.")
}
