// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	//Verbose identifies if extended output should be configured during init and reset
	Version = "2.1.1"
	Verbose bool
	rootCmd = &cobra.Command{
		Use:     "hookz",
		Short:   `Manages commit hooks inside a local git repository`,
		Version: Version,
	}
)

// Execute creates the command tree and handles any error condition returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	color.Style{color.FgWhite, color.OpBold}.Println("Hookz")
	fmt.Println("https://github.com/devops-kung-fu/hookz")
	fmt.Printf("Version: %s\n", Version)
	fmt.Println("")
}

func readConfig() (config Configuration, err error) {

	filename, _ := filepath.Abs(".hookz.yaml")
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		if err != nil {
			return
		}
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return
	}

	if config.Version == "" {
		err = errors.New("no configuration version value found in .hookz.yaml")
		return
	}

	// Check version
	ver := strings.Split(config.Version, ".")
	verMatch := strings.Split(Version, ".")
	if fmt.Sprintf("%v.%v", ver[0], ver[1]) != fmt.Sprintf("%v.%v", verMatch[0], verMatch[1]) {
		err = fmt.Errorf("version mismatch: Expected v%v.%v - Check your .hookz.yaml configuration", verMatch[0], verMatch[1])
	}
	return
}
