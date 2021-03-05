package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:     "initialize",
		Aliases: []string{"init"},
		Short:   "Initializes the hooks as defined in the .hookz.yaml file.",
		Long:    "Initializes the hooks as defined in the .hookz.yaml file.",
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
	initCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "If true, output from commands is displayed when the hook executes.")
}

func createFile(hookName string, hookType string) error {
	filename, _ := filepath.Abs(fmt.Sprintf(".git/hooks/%s", hookType))
	var _, err = os.Stat(filename)

	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		if err != nil {
			return err
		}

		defer file.Close()

		file, err = os.OpenFile(filename, os.O_RDWR, 0644)
		if err != nil {
			return err
		}

		header := `#!/bin/bash

reset='\033[0m'        # Text Reset
red='\033[41m'         # Red
green='\033[42m'       # Green
		
		`
		_, err = file.WriteString(fmt.Sprintf("%s\ntype='%s'\n", header, hookType))
		if err != nil {
			return err
		}

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

	exitCodeBlock := `
commandexit=$?

if [ $commandexit -eq 0 ]
then
	echo -e "$green[PASS]$reset $name ($type)"
else
	echo -e "$red[FAIL]$reset $name ($type)"
	exit $commandexit
fi
`
	for _, hook := range config.Hooks {
		createFile(hook.Name, hook.Type)
	}

	for _, hook := range config.Hooks {
		filename, _ := filepath.Abs(fmt.Sprintf(".git/hooks/%s", hook.Type))

		var file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("\nname='%s'\n", hook.Name))
		if err != nil {
			return err
		}

		if hook.URL != nil {
			//Download the binary, put it in the hooks folder and chmod 0777, and set the hook.Exec
			temp := "echo TODO://"
			hook.Exec = &temp
		}

		if hook.Exec != nil {
			if Verbose {
				_, err = file.WriteString(fmt.Sprintf("%s\n%s", *hook.Exec, exitCodeBlock))
			} else {
				_, err = file.WriteString(fmt.Sprintf("%s &> /dev/null\n %s", *hook.Exec, exitCodeBlock))
			}
			if err != nil {
				return err
			}
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
