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
	version = "2.4.2"
	//Afs stores a global OS Filesystem that is used throughout hookz
	Afs   = &afero.Afero{Fs: afero.NewOsFs()}
	debug bool
	//Verbose determines if the execution of hookz should output verbose information
	Verbose bool
	//VerboseOutput is set to true if you wish to see debug information in the hooks as they execute in bash
	VerboseOutput bool
	rootCmd       = &cobra.Command{
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
				color.Style{color.FgWhite, color.OpBold}.Println("█▀█ █▄█ █▄█ █ █ █▄")
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
		util.IfErrorLog(err)

		if !b {
			e := errors.New("hookz must be run in a local .git repository")
			util.PrintErr(e)
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
	rootCmd.PersistentFlags().BoolVar(&VerboseOutput, "verbose-output", false, "show verbose hook output while executing hooks")
}
