// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	rootCmd = &cobra.Command{
		Use:     "hookz",
		Short:   `Manages commit hooks inside a local git repository`,
		Version: "0.0.1",
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

}

func readConfig() (config Configuration, err error) {

	filename, _ := filepath.Abs(".hooks.yaml")
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		if err != nil {
			return
		}
	}

	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return
	}
	return
}

func isError(err error, pre string) error {
	if err != nil {
		log.Printf("%v: %v", pre, err)
	}
	return err
}

func isErrorBool(err error, pre string) (b bool) {
	if err != nil {
		log.Printf("%v: %v", pre, err)
		b = true
	}
	return
}

type Configuration struct {
	Hooks []struct {
		Name string   `json:"name"`
		Type string   `json:"type"`
		URL  *string  `json:"url,omitempty"`
		Args []string `json:"args"`
		Exec *string  `json:"exec,omitempty"`
	}
}
