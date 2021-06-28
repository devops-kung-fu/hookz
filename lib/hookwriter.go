//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/segmentio/ksuid"
)

type command struct {
	Name         string
	Type         string
	ShortCommand string
	FullCommand  string
	Debug        bool
}

func CreateFile(fs FileSystem, name string) (err error) {

	file, err := fs.fs.Create(name)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	return
}

func CreateScriptFile(fs FileSystem, content string) (name string, err error) {

	k, idErr := ksuid.NewRandom()
	name = k.String()
	if IsErrorBool(idErr, "ERROR") {
		err = idErr
		return
	}
	path, _ := os.Getwd()
	p := fmt.Sprintf("%s/%s", path, ".git/hooks")

	hookzFile := fmt.Sprintf("%s/%s.hookz", p, name)
	scriptName := fmt.Sprintf("%s/%s", p, name)

	err = CreateFile(fs, hookzFile)
	if err != nil {
		return
	}

	err = fs.Afero().WriteFile(scriptName, []byte(content), 0644)
	if err != nil {
		return
	}

	err = fs.fs.Chmod(scriptName, 0777)
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

func WriteHooks(fs FileSystem, config Configuration, verbose bool, debug bool) (err error) {

	for _, hook := range config.Hooks {

		var commands []command
		PrintIf(func() {
			fmt.Printf("\n[*] Writing %s \n", hook.Type)
		}, verbose)

		for _, action := range hook.Actions {
			if action.Exec == nil && action.URL != nil {
				filename, _ := DownloadURL(*action.URL)
				action.Exec = &filename
			}
			if action.Exec == nil && action.Script != nil {
				scriptFileName, err := CreateScriptFile(fs, *action.Script)
				if err != nil {
					return err
				}
				path, _ := os.Getwd()
				fullScriptFileName := fmt.Sprintf("%s/%s/%s", path, ".git/hooks", scriptFileName)
				action.Exec = &fullScriptFileName
			}

			PrintIf(func() {
				fmt.Printf("    	Adding %s action: %s\n", hook.Type, action.Name)
			}, verbose)

			fullCommand := buildFullCommand(action, debug)

			commands = append(commands, command{
				Name:         action.Name,
				Type:         hook.Type,
				ShortCommand: *action.Exec,
				FullCommand:  fullCommand,
				Debug:        debug,
			})
		}
		err = writeTemplate(fs, commands, hook.Type)
		if err != nil {
			return
		}
		PrintIf(func() {
			fmt.Println("[*] Successfully wrote " + hook.Type)
		}, verbose)

		PrintIf(func() {
			fmt.Println()
		}, verbose)
	}
	return nil
}

func writeTemplate(fs FileSystem, commands []command, hookType string) (err error) {
	path, _ := os.Getwd()
	p := fmt.Sprintf("%s/%s", path, ".git/hooks")

	hookzFile := fmt.Sprintf("%s/%s.hookz", p, hookType)
	err = CreateFile(fs, hookzFile)
	if err != nil {
		return
	}

	filename := fmt.Sprintf("%s/%s", p, hookType)
	file, err := fs.Afero().Create(filename)
	if err != nil {
		return err
	}
	t := genTemplate(hookType)
	err = t.ExecuteTemplate(file, hookType, commands)
	if err != nil {
		return err
	}
	err = fs.fs.Chmod(filename, 0777)
	if err != nil {
		return err
	}

	return
}

func HasExistingHookz(fs FileSystem) (exists bool) {
	path, _ := os.Getwd()
	ext := ".hookz"
	p := fmt.Sprintf("%s/%s", path, ".git/hooks")
	dirFiles, _ := fs.Afero().ReadDir(p)

	for index := range dirFiles {
		file := dirFiles[index]

		name := file.Name()
		fullPath := fmt.Sprintf("%s/%s", p, name)
		info, _ := fs.Afero().Stat(fullPath)
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

reset='\033[0m'         # Text Reset
red='\033[41m'          # Red Background
green='\033[42m'        # Green Background
blackText='\033[0;30m'  # Black Text
yellowText='\033[0;33m' # Purple Text
boldWhite='\e[1m'  		# Bold White
orange='\e[30;48;5;208m'	# Orange Background

echo -e "Hookz: Running $(basename $0)"

{{range .}}

{{if .Debug}}
echo -e "$yellowText >> START:$reset {{.Name}}"
{{end}}

if ! [ -x "$(command -v  {{.ShortCommand}})" ]; then
echo -e "$blackText$orange WARN $reset Hookz: {{.ShortCommand}} cannot be run. Command doesn't exist.({{.Type}})"
else

{{.FullCommand}}
commandexit=$?
if [ $commandexit -eq 0 ]
	then
			echo -e "$blackText$green PASS $reset Hookz: {{.Name}} ({{.Type}})"
	else
			echo -e "$blackText$red FAIL $reset Hookz: {{.Name}} ({{.Type}})"
			exit $commandexit
	fi
fi
{{if .Debug}}
echo -e "$yellowText >> END:$reset {{.Name}}"
echo -e "--------------------------------------------"
{{end}}

{{end}}
`
	return template.Must(template.New(hookType).Parse(content))
}
