package cmd

import (
	"fmt"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	config         lib.Configuration
	interactiveCmd = &cobra.Command{
		Use:   "interactive",
		Short: "An interactive CLI used to generate a .hookz.yaml file",
		Long:  "An interactive CLI used to generate a .hookz.yaml file",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := lib.ReadConfig(version)
			if err != nil {
				config = lib.Configuration{
					Version: version,
				}
			} else {
				config = c
			}
			startInteractive()
		},
	}
)

func init() {
	initCmd.AddCommand(interactiveCmd)
}

func startInteractive() {
	config.Version = version
	showMenu()
}

func showMenu() {
	prompt := promptui.Select{
		Label: "Select Action",
		Items: []string{"Show Configuration", "Add Hook",
			"Save", "Exit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Save":
		saveConfiguration()
	default:
		fmt.Println("Too far away.")
	}
}

func saveConfiguration() {
	prompt := promptui.Prompt{
		Label:     "Save Configuration",
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "y":
		fmt.Print("TODO://Save Configuration")
	default:
		showMenu()
	}
}
