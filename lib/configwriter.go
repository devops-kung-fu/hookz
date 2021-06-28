package lib

import (
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func createConfig(fs FileSystem, version string) (config Configuration, err error) {
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
	err = fs.Afero().WriteFile(filename, file, 0644)

	if err != nil {
		return
	}

	return
}
