package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/devops-kung-fu/common/util"

	"github.com/devops-kung-fu/hookz/lib"
)

func noConfig() {
	fmt.Println(".hookz.yaml file not found")
	fmt.Println("\nTo create a sample configuration run:")
	fmt.Println("        hookz init config")
	fmt.Println("\nRun 'hookz --help' for usage.")
	fmt.Println()
	os.Exit(1)
}

func badYaml() {
	util.PrintErr(errors.New("configuration in .hookz.yaml is not valid YAML syntax"))
	os.Exit(1)
}

// CheckConfig ensures that there is a .hookz.yaml file locally and the version is supported by the current version of hookz
func CheckConfig() (config lib.Configuration) {
	config, err := lib.ReadConfig(Afs, version)
	if err != nil && err.Error() == "NO_CONFIG" {
		noConfig()
	} else if err != nil && err.Error() == "BAD_YAML" {
		badYaml()
	}
	return
}

// InstallSources installs all go repositories that are found in the Sources section of the .hookz.yaml file.
func InstallSources(sources []lib.Source) (err error) {
	if len(sources) > 0 && Verbose {
		util.DoIf(Verbose, func() {
			util.PrintInfo("Installing sources...")
		})
	}
	for _, s := range sources {
		util.DoIf(Verbose, func() {
			util.PrintTabbedf("Installing Source: %s\n", s.Source)
		})
		err = lib.InstallSource(s)
		if err != nil {
			log.Printf("There was a problem installing source %s, error: %s", s.Source, err)
			return
		}
	}
	return
}
