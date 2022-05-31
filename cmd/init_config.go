package cmd

import (
	"fmt"
	"os"

	"github.com/devops-kung-fu/common/util"
	"github.com/gookit/color"
	"github.com/spf13/cobra"

	"github.com/devops-kung-fu/hookz/lib"
)

var (
	initConfigCmd = &cobra.Command{
		Use:   "config",
		Short: "Creates a starter .hookz.yaml file.",
		Long:  "Creates a starter .hookz.yaml file.",
		PreRun: func(cmd *cobra.Command, args []string) {
			if lib.HasExistingHookzYaml(Afs) {
				color.Style{color.FgRed, color.OpBold}.Println("Existing .hookz.yaml file detected!")
				fmt.Println("\nThis file must be deleted before running this command.")
				fmt.Println()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			util.DoIf(Verbose, func() {
				util.PrintInfo("Creating Sample Config")
			})
			_, err := lib.CreateConfig(Afs, version)
			if util.IsErrorBool(err) {
				return
			}
			util.DoIf(Verbose, func() {
				util.PrintSuccess("Done")
			})
		},
	}
)

func init() {
	initCmd.AddCommand(initConfigCmd)
}
