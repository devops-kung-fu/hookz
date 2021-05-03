//Package lib Functionality for the Hookz CLI
package lib

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
)

func (d Deps) ReadConfig(version string) (config Configuration, err error) {

	filename, _ := filepath.Abs(".hookz.yaml")
	//yamlFile, readErr := ioutil.ReadFile(filename)
	yamlFile, readErr := d.Afero().ReadFile(filename)
	if readErr != nil {
		config, err = promptCreateConfig(version)
		if err != nil {
			return
		}
	} else {
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			return
		}
	}

	err = checkVersion(config, version)

	return
}

func checkVersion(config Configuration, version string) (err error) {
	if config.Version == "" {
		err = errors.New("no configuration version value found in .hookz.yaml")
		return
	}
	ver := strings.Split(config.Version, ".")
	verMatch := strings.Split(version, ".")
	if fmt.Sprintf("%v.%v", ver[0], ver[1]) != fmt.Sprintf("%v.%v", verMatch[0], verMatch[1]) {
		err = fmt.Errorf("version mismatch: Expected v%v.%v - Check your .hookz.yaml configuration", verMatch[0], verMatch[1])
	}
	return
}

func promptCreateConfig(version string) (config Configuration, err error) {
	fmt.Println("\nHookz was unable to find a .hookz.yaml file. Would you like")
	fmt.Println("to create a starter configuration?")

	prompt := promptui.Prompt{
		Label:     "Create starter .hookz.yaml?",
		IsConfirm: true,
	}

	result, promptErr := prompt.Run()

	if promptErr != nil {
		goto CONFIG_ERR
	}

	if result == "y" {
		command := "echo"
		config = Configuration{
			Version: version,
			Hooks: []Hook{
				{
					Type: "pre-commit",
					Actions: []Action{
						{
							Name: "Hello Hookz!",
							Exec: &command,
							Args: []string{"-e", "Hello Hookz!"},
						},
					},
				},
			},
		}

		file, merr := yaml.Marshal(config)
		if merr != nil {
			err = merr
			return
		}

		err = ioutil.WriteFile(".hooks.yaml", file, 0644)
		if err != nil {
			return
		}

		return
	}

CONFIG_ERR:
	log.Println("[ERROR]: Hookz cannot run without a .hookz.yaml file. Create one and try again")
	os.Exit(1)

	return
}
