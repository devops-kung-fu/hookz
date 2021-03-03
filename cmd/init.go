package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes the hooks as defined in the .hookz.yaml file.",
		Long:  "Initializes the hooks as defined in the .hookz.yaml file.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Initializing git hooks...")
			if isErrorBool(writeHooks(), "[ERROR]") {
				return
			}
			fmt.Println("DONE")
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

func createFile(filename string) error {
	var _, err = os.Stat(filename)

	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		err = os.Chmod(filename, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeHooks() error {
	var config, err = readConfig()
	if err != nil {
		return err
	}

	for _, hook := range config.Hooks {
		filename, _ := filepath.Abs(".git/hooks/" + hook.Type)
		createFile(filename)

		var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString("#!/bin/sh\necho " + hook.Type + "\necho " + hook.Name)
		if err != nil {
			return err
		}

		err = file.Sync()
		if err != nil {
			return err
		} else {
			fmt.Println("[*] Successfully wrote " + hook.Type)
		}
	}
	return nil
}
