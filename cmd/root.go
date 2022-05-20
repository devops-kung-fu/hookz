// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/devops-kung-fu/common/util"
	"github.com/gookit/color"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	version = "2.4.0"
	Afs     = &afero.Afero{Fs: afero.NewOsFs()}
	debug   bool
	Verbose bool
	rootCmd = &cobra.Command{
		Use:     "hookz",
		Short:   `Manages commit hooks inside a local git repository`,
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !debug {
				log.SetOutput(ioutil.Discard)
			}
			util.DoIf(Verbose, func() {
				fmt.Println()
				color.Style{color.FgWhite, color.OpBold}.Println("█ █ █▀█ █▀█ █▄▀ ▀█")
				color.Style{color.FgWhite, color.OpBold}.Println("█▀█ █▄█ █▄█ █░█ █▄")
				fmt.Println()
				fmt.Println("DKFM - DevOps Kung Fu Mafia")
				fmt.Println("https://github.com/devops-kung-fu/hookz")
				fmt.Printf("Version: %s\n", version)
				fmt.Println()
			})
		},
	}
)

// Execute creates the command tree and handles any error condition returned
func Execute() {
	cobra.OnInitialize(func() {
		b, err := Afs.DirExists(".git")
		util.IfErrorLog(err, "[ERROR]")

		if !b {
			e := errors.New("hookz must be run in a local .git repository")
			util.IfErrorLog(e, "ERROR")
			os.Exit(1)
		}
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show debug output")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", true, "show verbose output")
}
