package cmd

import (
	"fmt"
	"os"

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
