//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var (
	deps Deps = Deps{
		fs: afero.NewMemMapFs(),
	}
	version string = "1.0.0"
	config  Configuration
)

func TestDeps_ReadConfig(t *testing.T) {

	config, _ = deps.createConfig(version)
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
