// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	version = "2.2.0"
	verbose bool
	rootCmd = &cobra.Command{
		Use:     "hookz",
		Short:   `Manages commit hooks inside a local git repository`,
		Version: version,
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
	fmt.Printf("Version: %s\n", version)
	fmt.Println("")
}
