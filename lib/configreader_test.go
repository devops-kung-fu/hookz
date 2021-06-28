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

	config, _ = createConfig(fs, version)
	readConfig, err := ReadConfig(fs, version)

	assert.NoError(t, err, "ReadConfig should not have generated an error")
	assert.Equal(t, version, readConfig.Version, "Versions should match")

	_, err = ReadConfig(fs, "")
	assert.Error(t, err, "Passing an empty string should cause an error")
}

func TestDeps_checkVersion(t *testing.T) {
	readConfig, err := ReadConfig(fs, version)
	assert.NoError(t, err, "ReadConfig should not have generated an error")

	err = checkVersion(readConfig, version)
	assert.NoError(t, err, "Check version should not have generated an error")

	err = checkVersion(readConfig, "2.0")
	assert.Error(t, err, "Version mismatch not caught")

	readConfig.Version = ""
	err = checkVersion(readConfig, version)
	assert.Error(t, err, "An empty config version should throw an error")
}

func Test_promptCreateConfig(t *testing.T) {

	config, err := promptCreateConfig(fs, version)
	assert.Equal(t, version, config.Version, "Version mismatch")
	assert.NoError(t, err, "Expected no error to be thrown")
}
