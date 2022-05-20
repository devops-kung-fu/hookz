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
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates any executable defined as an URL attribute in .hookz.yaml.",
		Long:  "Rebuilds the hooks as defined in the .hookz.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			color.Style{color.FgLightBlue, color.OpBold}.Println("Update Executables")
			fmt.Println()
			config, err := lib.ReadConfig(Afs, version)
			if err != nil && err.Error() == "NO_CONFIG" {
				NoConfig()
				os.Exit(1)
			}
			if util.IsErrorBool(err, "[ERROR]") {
				return
			}
			if util.IsErrorBool(lib.UpdateExecutables(Afs, config), "[ERROR]") {
				return
			}
			color.Style{color.FgLightGreen}.Println("Done!")
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
}
