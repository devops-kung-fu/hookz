package cmd

import (
	"fmt"
	"os"

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
			existingHookz := lib.HasExistingHookzYaml(Afs)
			if existingHookz {
				color.Style{color.FgRed, color.OpBold}.Println("Existing .hookz.yaml file detected!")
				fmt.Println("\nThis file must be deleted before running this command.")
				fmt.Println()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			color.Style{color.FgLightBlue, color.OpBold}.Println("Creating Sample Config")
			fmt.Println()
			_, err := lib.CreateConfig(Afs, version)
			if lib.IsErrorBool(err, "[ERROR]") {
				return
			}
			color.Style{color.FgLightGreen}.Println("Done!")
		},
	}
)

func init() {
	initCmd.AddCommand(initConfigCmd)
}
