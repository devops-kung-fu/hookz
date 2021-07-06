//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var (
	fs FileSystem = FileSystem{
		fs: afero.NewMemMapFs(),
	}
	version string = "1.0.0"
	config  Configuration
)

func TestDeps_ReadConfig(t *testing.T) {

	newFs := FileSystem{
		fs: afero.NewMemMapFs(),
	}

	_, err := ReadConfig(newFs, version)
	assert.Error(t, err, "There should be no config created so an error should be thrown.")
	config, _ = CreateConfig(newFs, version)
	readConfig, err := ReadConfig(newFs, version)

	assert.NoError(t, err, "ReadConfig should not have generated an error")
	assert.Equal(t, version, readConfig.Version, "Versions should match")

	_, err = ReadConfig(newFs, "")
	assert.Error(t, err, "Passing an empty string should cause an error")
}
