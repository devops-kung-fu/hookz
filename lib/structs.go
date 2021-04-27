//Package lib Functionality for the Hookz CLI
package lib

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

func (hook *Hook) Create() (err error) {

	// for _, a := range hook.Actions {

	// }

	return
}

//Action is a task that will execute when a hook executes.
type Action struct {
	Name   string   `json:"name"`
	URL    *string  `json:"URL,omitempty"`
	Args   []string `json:"args,omitempty"`
	Exec   *string  `json:"exec,omitempty"`
	Script *string  `json:"script,omitempty"`
}

func (action *Action) Create() (err error) {
	return
}
