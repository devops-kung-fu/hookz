// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/gookit/color"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	version = "2.3.0"
	debug   bool
	verbose bool
	rootCmd = &cobra.Command{
		Use:     "hookz",
		Short:   `Manages commit hooks inside a local git repository`,
		Version: version,
	}
)

// Execute creates the command tree and handles any error condition returned
func Execute() {
	cobra.OnInitialize(func() {
		var fs = afero.NewOsFs()
		afs := &afero.Afero{Fs: fs}
		b, err := afs.DirExists(".git")
		lib.IfErrorLog(err, "[ERROR]")

		if !b {
			e := errors.New("Hookz must be run in a local .git repository")
			lib.IfErrorLog(e, "ERROR")
			os.Exit(1)
		}
	})
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

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Extended output as hookz executes.")
}
