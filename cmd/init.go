package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/segmentio/ksuid"
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
			if lib.IsErrorBool(writeHooks(), "[ERROR]") {
				return
			}
			fmt.Println("\nDONE!")
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "If true, output from commands is displayed when the hook executes.")
}

//TODO: This is WAY too complex... should just use:
//  f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

func createFile(name string) error {

	var _, err = os.Stat(name)

	if os.IsNotExist(err) {
		var file, err = os.Create(name)
		if err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}

func createScriptFile(content string) (name string, err error) {
	var _, statErr = os.Stat(name)
	k, idErr := ksuid.NewRandom()
	name, _ = filepath.Abs(fmt.Sprintf(".git/hooks/%s", k.String()))
	if idErr != nil {
		fmt.Printf("Error generating KSUID: %v\n", err)
		return
	}

	if os.IsNotExist(statErr) {
		hookzFile, hookzFileErr := filepath.Abs(fmt.Sprintf(".git/hooks/%s.hookz", k.String()))
		createFile(hookzFile)
		if err != nil {
			err = hookzFileErr
			return
		}

		var file, createErr = os.Create(name)
		if err != nil {
			err = createErr
			return
		}

		err = os.Chmod(name, 0777)
		if err != nil {
			return
		}

		file, err = os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		_, err = file.WriteString(content)
		if err != nil {
			return
		}

		defer file.Close()
	}

	return
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
	echo -e "$blackText$green PASS $reset Hookz: $name ($type)"
else
	echo -e "$blackText$red FAIL $reset Hookz: $name ($type)"
	exit $commandexit
fi
`
	for _, hook := range config.Hooks {
		header := `#!/bin/bash

# This file was generated by Hookz
# For more information, check out https://github.com/devops-kung-fu/hookz

reset='\033[0m'        	# Text Reset
red='\033[41m'        	# Red Background
green='\033[42m'       	# Green Background
blackText='\033[0;30m' 	# Black Text
`
		filename, _ := filepath.Abs(fmt.Sprintf(".git/hooks/%s", hook.Type))
		hookzFile, _ := filepath.Abs(fmt.Sprintf(".git/hooks/%s.hookz", hook.Type))

		createFile(filename)
		createFile(hookzFile)

		err = os.Chmod(filename, 0777)
		if err != nil {
			return err
		}

		var file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		_, err = file.WriteString(header)
		if err != nil {
			return err
		}

		defer file.Close()
		fmt.Println(fmt.Sprintf("\n[*] Writing %s ", hook.Type))
		for _, action := range hook.Actions {

			var argsString string
			for _, arg := range action.Args {
				argsString = fmt.Sprintf("%s %s", argsString, arg)
			}

			if action.Exec == nil && action.URL != nil {
				filename, _ := lib.DownloadURL(*action.URL)
				action.Exec = &filename
			}

			if action.Exec == nil && action.Script != nil {
				scriptFileName, err := createScriptFile(*action.Script)
				if err != nil {
					return err
				}
				action.Exec = &scriptFileName
			}

			fmt.Println(fmt.Sprintf("    	Adding %s action: %s", hook.Type, action.Name))
			_, err = file.WriteString(fmt.Sprintf("name='%s'\ntype='%s'\n", action.Name, hook.Type))
			if err != nil {
				return err
			}

			if action.Exec != nil {
				if Verbose {
					_, err = file.WriteString(fmt.Sprintf("echo '[*] Executing %s: %s'\n", hook.Type, action.Name))
					if err != nil {
						return err
					}
					_, err = file.WriteString(fmt.Sprintf("%s%s\n%s", *action.Exec, argsString, exitCodeBlock))
				} else {
					_, err = file.WriteString(fmt.Sprintf("%s%s &> /dev/null\n %s", *action.Exec, argsString, exitCodeBlock))
				}
				if err != nil {
					return err
				}
			}
		}

		err = file.Sync()
		if err != nil {
			return err
		}
		fmt.Println("[*] Successfully wrote " + hook.Type)

	}
	return nil
}
