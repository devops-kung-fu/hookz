//Package lib Functionality for the Hookz CLI
package lib

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var (
	deps Deps = Deps{
		fs: afero.NewMemMapFs(),
	}
	version string = "1.0.0"
	config  Configuration
)

func TestDeps_ReadConfig(t *testing.T) {

	config, _ = createTestConfig(version)
	readConfig, err := deps.ReadConfig(version)

	assert.NoError(t, err, "ReadConfig should not have generated an error")
	assert.Equal(t, version, readConfig.Version, "Versions should match")
}

func TestDeps_checkVersion(t *testing.T) {
	readConfig, err := deps.ReadConfig(version)
	assert.NoError(t, err, "ReadConfig should not have generated an error")
	err = checkVersion(readConfig, version)
	assert.NoError(t, err, "Check version should not have generated an error")

}

func createTestConfig(version string) (config Configuration, err error) {
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
	err = deps.Afero().WriteFile(filename, file, 0644)
	if err != nil {
		return
	}

	return
}
