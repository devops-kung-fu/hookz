package lib

import (
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

// CreateConfig creates a starter .hookz.yaml file
func CreateConfig(afs *afero.Afero, version string) (config Configuration, err error) {
	command := "/bin/echo"
	config = Configuration{
		Version: version,
		Sources: []Source{
			{
				Source: "github.com/devops-kung-fu/hinge@latest",
			},
		},
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
	err = afs.WriteFile(filename, file, 0644)

	return
}
