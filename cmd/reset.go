package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"

	"github.com/devops-kung-fu/hookz/lib"
)

var (
	resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Rebuilds the hooks as defined in the .hookz.yaml file.",
		Long:  "Rebuilds the hooks as defined in the .hookz.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			color.Style{color.FgLightBlue, color.OpBold}.Println("Reset Hookz")
			fmt.Println()

			if lib.IsErrorBool(lib.RemoveHooks(Afs, Verbose), "[ERROR]") {
				return
			}
			config, err := lib.ReadConfig(Afs, version)
			if err != nil && err.Error() == "NO_CONFIG" {
				os.Exit(1)
			}
			if lib.IsErrorBool(err, "[ERROR]") {
				return
			}
			err = lib.InstallSources(config.Sources)
			if err != nil {
				log.Println("There was a problem installing sources")
				log.Println(err)
			}
			if lib.IsErrorBool(lib.WriteHooks(Afs, config, Verbose, debug), "[ERROR]") {
				return
			}
			color.Style{color.FgLightGreen}.Println("Done!")
		},
	}
)

func init() {
	rootCmd.AddCommand(resetCmd)
	resetCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "If true, output from commands is displayed when the hook executes.")
}
