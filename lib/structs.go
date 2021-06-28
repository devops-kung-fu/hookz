//Package lib Functionality for the Hookz CLI
package lib

import (
	"github.com/spf13/afero"
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

//FileSystem wraps the afero.Fs interface to allow for mocking
type FileSystem struct {
	fs afero.Fs
}

//NewOsFs creates a new disk based file system
func NewOsFs() FileSystem {
	var d FileSystem
	d.fs = afero.NewOsFs()
	return d
}

//Afero returns a new Afero struct wraping the current file system
func (fs FileSystem) Afero() (afs *afero.Afero) {
	afs = &afero.Afero{Fs: fs.fs}
	return
}
