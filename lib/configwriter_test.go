//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_CreateConfig(t *testing.T) {

	newFs := FileSystem{
		fs: afero.NewMemMapFs(),
	}

	exists := HasExistingHookzYaml(newFs)
	assert.False(t, exists, "No hookz.yaml file should exist")

	_, err := CreateConfig(newFs, version)
	assert.NoError(t, err, "CreateConfig should not have generated an error")

	exists = HasExistingHookzYaml(newFs)
	assert.True(t, exists, "A hookz.yaml file should exist")

	readConfig, _ := ReadConfig(newFs, version)
	assert.Equal(t, version, readConfig.Version, "A Version should have been detected in the config after creation.")
}
