package cmd

import (
	"fmt"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates any executable defined as an URL attribute in .hooks.yaml.",
		Long:  "Rebuilds the hooks as defined in the .hooks.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Updating executables...")
			if lib.IsErrorBool(updateExecutables(), "[ERROR]") {
			}
			fmt.Println("\nDONE!")
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateExecutables() (err error) {
	var config Configuration
	config, err = readConfig()
	if err != nil {
		return err
	}
	for _, hook := range config.Hooks {
		for _, action := range hook.Actions {
			if action.URL != nil {
				_, _ = lib.DownloadURL(*action.URL)
			}
		}
	}
	return
}
