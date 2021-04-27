package cmd

type Configuration struct {
	Version string `json:"version"`
	Hooks   []Hook `json:"hooks"`
}

type Hook struct {
	Type    string   `json:"type"`
	Actions []Action `json:"actions"`
}

func (hook *Hook) Create() (err error) {

	// for _, a := range hook.Actions {

	// }

	return
}

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
