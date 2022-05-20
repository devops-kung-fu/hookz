//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/devops-kung-fu/common/util"
	"github.com/segmentio/ksuid"
	"github.com/spf13/afero"
)

type command struct {
	Name         string
	Type         string
	ShortCommand string
	FullCommand  string
	Debug        bool
}

//CreateFile creates a file for a provided FileSystem and file name
func CreateFile(afs *afero.Afero, name string) (err error) {

	file, err := afs.Fs.Create(name)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	return
}

//CreateScriptFile creates an executable script file with a random name given a string of content
func CreateScriptFile(afs *afero.Afero, content string) (name string, err error) {

	k, idErr := ksuid.NewRandom()
	name = k.String()
	if util.IsErrorBool(idErr, "ERROR") {
		err = idErr
		return
	}
	path, _ := os.Getwd()
	p := fmt.Sprintf("%s/%s", path, ".git/hooks")

	hookzFile := fmt.Sprintf("%s/%s.hookz", p, name)
	scriptName := fmt.Sprintf("%s/%s", p, name)

	err = CreateFile(afs, hookzFile)
	if err != nil {
		return
	}

	err = afs.WriteFile(scriptName, []byte(content), 0644)
	if err != nil {
		return
	}

	err = afs.Fs.Chmod(scriptName, 0777)
	if err != nil {
		return
	}

	return
}

func buildFullCommand(action Action, debug bool) string {
	var argsString, fullCommand string
	for _, arg := range action.Args {
		argsString = fmt.Sprintf("%s %s", argsString, arg)
	}
	if action.Exec != nil {
		if debug {
			fullCommand = fmt.Sprintf("%s%s", *action.Exec, argsString)
		} else {
			fullCommand = fmt.Sprintf("%s%s &> /dev/null", *action.Exec, argsString)
		}
	}
	return fullCommand
}

//WriteHooks writes all of the generated scripts to the .git/hooks directory
func WriteHooks(afs *afero.Afero, config Configuration, verbose bool, verboseOutput bool) (err error) {
	log.Println("Writing hooks")
	for _, hook := range config.Hooks {
		var commands []command
		util.DoIf(verbose, func() {
			util.PrintInfo(fmt.Sprintf("Writing %s", hook.Type))
		})

		for _, action := range hook.Actions {
			err = buildExec(afs, &action)
			if err != nil {
				return err
			}
			util.DoIf(verbose, func() {
				util.PrintTabbed(fmt.Sprintf("Adding %s action: %s", hook.Type, action.Name))
			})

			fullCommand := buildFullCommand(action, verboseOutput)

			commands = append(commands, command{
				Name:         action.Name,
				Type:         hook.Type,
				ShortCommand: *action.Exec,
				FullCommand:  fullCommand,
				Debug:        verboseOutput,
			})
		}
		err = writeTemplate(afs, commands, hook.Type)
		if err != nil {
			return
		}
		util.DoIf(verbose, func() {
			util.PrintSuccess(fmt.Sprintf("Successfully wrote %s", hook.Type))
		})
	}

	_ = WriteShasum(afs)
	return
}

func buildExec(afs *afero.Afero, action *Action) (err error) {
	if action.Exec == nil && action.URL != nil {
		filename, err := DownloadFile(afs, ".git/hooks", *action.URL)
		action.Exec = &filename
		if err != nil {
			return err
		}
	}
	if action.Exec == nil && action.Script != nil {
		scriptFileName, _ := CreateScriptFile(afs, *action.Script)
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		fullScriptFileName := fmt.Sprintf("%s/%s/%s", path, ".git/hooks", scriptFileName)
		action.Exec = &fullScriptFileName
	}
	return
}

func writeTemplate(afs *afero.Afero, commands []command, hookType string) (err error) {
	path, _ := os.Getwd()
	p := fmt.Sprintf("%s/%s", path, ".git/hooks")

	hookzFile := fmt.Sprintf("%s/%s.hookz", p, hookType)
	err = CreateFile(afs, hookzFile)
	if err != nil {
		return
	}

	filename := fmt.Sprintf("%s/%s", p, hookType)
	file, err := afs.Create(filename)
	if err != nil {
		return err
	}
	t := genTemplate(hookType)
	err = t.ExecuteTemplate(file, hookType, commands)
	if err != nil {
		return err
	}
	err = afs.Fs.Chmod(filename, 0777)
	if err != nil {
		return err
	}

	return
}

//HasExistingHookz determines if any .hookz touch files exist in the .git/hooks directory
func HasExistingHookz(afs *afero.Afero) (exists bool) {
	path, _ := os.Getwd()
	ext := ".hookz"
	p := fmt.Sprintf("%s/%s", path, ".git/hooks")
	dirFiles, _ := afs.ReadDir(p)

	for index := range dirFiles {
		file := dirFiles[index]

		name := file.Name()
		fullPath := fmt.Sprintf("%s/%s", p, name)
		info, _ := afs.Stat(fullPath)
		isHookzFile := strings.Contains(info.Name(), ext)
		if isHookzFile {
			return true
		}
	}

	return false
}

func genTemplate(hookType string) (t *template.Template) {

	content := `#!/bin/bash

# This file was generated by Hookz
# For more information, check out https://github.com/devops-kung-fu/hookz

echo -e "$(tput bold)Hookz$(tput sgr0)"
echo -e "https://github.com/devops-kung-fu/hookz"
echo -e "Version: 2.3.0"
echo

shasum=$(cat .git/hooks/hookz.shasum)
check=$(shasum -a 256 .hookz.yaml | cut -d " " -f 1)

if [ "$check" != "$shasum" ]; then
	echo -e "$(tput setab 124 && tput setaf 248;) FAIL $(tput sgr0) Configuration change detected"
	echo
	echo "It appears your configuration has changed."
	echo "Please regenerate your hooks with the following"
	echo "command and try again."
	echo
	echo "        hookz reset [--verbose] [--debug] [--verbose-output]"
	echo
	echo "Run 'hookz --help' for usage."
	echo
	exit 1
fi

echo -e "$(tput bold)Hookz$(tput sgr0): Running $(basename $0)"

{{range .}}

{{if .Debug}}
	echo -e "$(tput setaf 5) >> START:$(tput sgr0) {{.Name}}"
{{end}}

if ! [ -x "$(command -v  {{.ShortCommand}})" ]; then
	echo -e "$(tput setab 214 && tput setaf 238;) WARN $(tput sgr0) $(tput bold)Hooks$(tput sgr0): {{.ShortCommand}} cannot be run. Command doesn't exist.({{.Type}})"
else

{{.FullCommand}}
commandexit=$?
if [ $commandexit -eq 0 ]
	then
			echo -e "$(tput setab 34 && tput setaf 238;) PASS $(tput sgr0) $(tput bold)Hookz$(tput sgr0): {{.Name}} ({{.Type}})"
	else
			echo -e "$(tput setab 124 && tput setaf 248;) FAIL $(tput sgr0) $(tput bold)Hookz$(tput sgr0): {{.Name}} ({{.Type}})"
			exit $commandexit
	fi
fi
{{if .Debug}}
	echo -e "$(tput setaf 5) >> END:$(tput sgr0) {{.Name}}"
	echo -e "$(tput setaf 248;)----------------------------------------------------------------------------------------$(tput sgr0)"
{{end}}

{{end}}
`
	return template.Must(template.New(hookType).Parse(content))
}
