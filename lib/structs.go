//Package lib Functionality for the Hookz CLI
package lib

//Configuration represents the content of .hookz.yaml
type Configuration struct {
	Version string   `json:"version"`
	Hooks   []Hook   `json:"hooks"`
	Sources []Source `json:"source"`
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

//Source defines a go repository that should be installed when hookz is initializing, updating, or resetting
type Source struct {
	Source string `json:"source,omitempty"`
}
