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
