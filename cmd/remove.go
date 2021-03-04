package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Removes the hooks as defined in the .hooks.yaml file.",
		Long:  "Removes the hooks as defined in the .hooks.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Removing git hooks...")
			if isErrorBool(removeHooks(), "[ERROR]") {
				return
			}
			fmt.Println("DONE")
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

func removeHooks() error {
	var config, err = readConfig()
	if err != nil {
		return err
	}

	for _, hook := range config.Hooks {
		filename, _ := filepath.Abs(".git/hooks/" + hook.Type)
		var err = os.Remove(filename)
		if err != nil {
			return err
		}

		fmt.Println("[*]" + " Deleted " + hook.Type)
	}
	return nil
}
