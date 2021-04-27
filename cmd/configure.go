package cmd

import (
	"fmt"
	"log"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	config       lib.Configuration
	configureCmd = &cobra.Command{
		Use:   "configure",
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
			color.Style{color.FgWhite, color.OpBold}.Println("Configuring .hooks.yaml...\n")
			startInteractive()
		},
	}
)

func init() {
	rootCmd.AddCommand(configureCmd)
}

func startInteractive() {
	config.Version = version
	showMenu()
}

func showMenu() {
	prompt := promptui.Select{
		Label: "Select Action",
		Items: []string{"Show Configuration", "Add Hook",
			"Delete Hook", "Exit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Show Configuration":
		showConfiguration()
	case "Add Hook":
		hook := addHook()
		if hook.Type != "" {
			// actions := addActions()
		}
	case "Exit":
		return
	}
}

func addHook() (hook lib.Hook) {

	prompt := promptui.Select{
		Label: "Hook Type",
		Items: []string{
			"applypatch-msg",
			"commit-msg",
			"fsmonitor-watchman",
			"post-commit",
			"post-update",
			"pre-applypatch",
			"pre-commit",
			"pre-update",
			"prepare-commit-msg",
			"pre-push",
			"pre-rebase",
			"pre-receive",
			"update",
			"cancel"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	if result == "cancel" {
		return
	} else {
		hook.Type = result
	}
	return hook
}

// func saveConfiguration() {
// 	prompt := promptui.Prompt{
// 		Label:     "Save Configuration",
// 		IsConfirm: true,
// 	}

// 	result, err := prompt.Run()

// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		return
// 	}

// 	switch result {
// 	case "y":
// 		fmt.Print("TODO://Save Configuration")
// 	default:
// 		showMenu()
// 	}
// }

func showConfiguration() {
	yaml, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal("Failed to generate yaml", err)
	}
	fmt.Printf("%s\n", string(yaml))
}
