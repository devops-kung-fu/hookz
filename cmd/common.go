package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/devops-kung-fu/common/util"
	"github.com/spf13/afero"

	"github.com/devops-kung-fu/hookz/lib"
)

func NoConfig() {
	fmt.Println(".hookz.yaml file not found")
	fmt.Println("\nTo create a sample configuration run:")
	fmt.Println("        hookz init config")
	fmt.Println("\nRun 'hookz --help' for usage.")
	fmt.Println()
}

// TODO: add code coverage
// CheckConfig ensures that there is a .hookz.yaml file locally and the version is supported by the current version of hookz
func CheckConfig(afs *afero.Afero) (config lib.Configuration, err error) {
	config, err = lib.ReadConfig(Afs, version)
	var returnErr error
	if err != nil && err.Error() == "NO_CONFIG" {
		returnErr = errors.New("NO_CONFIG")
	} else if err != nil && err.Error() == "BAD_YAML" {
		returnErr = errors.New("configuration in .hookz.yaml is not valid YAML syntax")
	}
	return config, returnErr
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
