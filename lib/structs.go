//Package lib Functionality for the Hookz CLI
package lib

import (
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

//Configuration represents the content of .hookz.yaml
type Configuration struct {
	Version string `json:"version"`
	Hooks   []Hook `json:"hooks"`
}

//Hook is the definition of a collection of actions to be run at
//a specified stage in the commit process. Example Type: pre-commit
type Hook struct {
	Type    string   `json:"type"`
	Actions []Action `json:"actions"`
}

//Action is a task that will execute when a hook executes.
type Action struct {
	Name   string   `json:"name"`
	URL    *string  `json:"URL,omitempty"`
	Args   []string `json:"args,omitempty"`
	Exec   *string  `json:"exec,omitempty"`
	Script *string  `json:"script,omitempty"`
}

type FileSystem struct {
	fs afero.Fs
}

func NewOsFs() FileSystem {
	var d FileSystem
	d.fs = afero.NewOsFs()
	return d
}

func (f FileSystem) Afero() (afs *afero.Afero) {
	afs = &afero.Afero{Fs: f.fs}
	return
}

func (f FileSystem) createConfig(version string) (config Configuration, err error) {
	command := "echo"
	config = Configuration{
		Version: version,
		Hooks: []Hook{
			{
				Type: "pre-commit",
				Actions: []Action{
					{
						Name: "Hello Hookz!",
						Exec: &command,
						Args: []string{"-e", "Hello Hookz!"},
					},
				},
			},
		},
	}

	file, memoryErr := yaml.Marshal(config)
	if memoryErr != nil {
		err = memoryErr
		return
	}
	filename, _ := filepath.Abs(".hookz.yaml")
	err = f.Afero().WriteFile(filename, file, 0644)
	if err != nil {
		return
	}

	return
}
