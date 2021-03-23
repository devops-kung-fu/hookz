package cmd

// Generated by https://quicktype.io

type Configuration struct {
		Version float64 `json:"version"`
		Hooks   []Hook  `json:"hooks"`
}

type Hook struct {
	Type    string   `json:"type"`
	Actions []Action `json:"actions"`
}

type Action struct {
	Name string   `json:"name"`
	URL  *string  `json:"URL,omitempty"`
	Args []string `json:"args"`
	Exec *string  `json:"exec,omitempty"`
}
