//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var (
	f FileSystem = FileSystem{
		fs: afero.NewMemMapFs(),
	}
	version string = "1.0.0"
	config  Configuration
)

func TestDeps_ReadConfig(t *testing.T) {

	config, _ = f.createConfig(version)
	readConfig, err := f.ReadConfig(version)

	assert.NoError(t, err, "ReadConfig should not have generated an error")
	assert.Equal(t, version, readConfig.Version, "Versions should match")
}

func TestDeps_checkVersion(t *testing.T) {
	readConfig, err := f.ReadConfig(version)
	assert.NoError(t, err, "ReadConfig should not have generated an error")

	err = checkVersion(readConfig, version)
	assert.NoError(t, err, "Check version should not have generated an error")

	err = checkVersion(readConfig, "2.0")
	assert.Error(t, err, "Version mismatch not caught")

	readConfig.Version = ""
	err = checkVersion(readConfig, version)
	assert.Error(t, err, "An empty config version should throw an error")
}
