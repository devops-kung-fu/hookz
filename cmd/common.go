package cmd

import (
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

func CheckConfig() (config lib.Configuration) {
	config, err := lib.ReadConfig(Afs, version)
	if err != nil && err.Error() == "NO_CONFIG" {
		noConfig()
	}
	return
}

func InstallSources(sources []lib.Source) (err error) {
	util.DoIf(Verbose, func() {
		util.PrintInfo("Installing Sources...")
	})
	for _, s := range sources {
		util.DoIf(Verbose, func() {
			util.PrintTabbedf("Installing Source: %s", s.Source)
		})
		err = lib.InstallSource(s)
		if err != nil {
			log.Printf("There was a problem installing source %s, error: %s", s.Source, err)
			return
		}
	}
	return
}
