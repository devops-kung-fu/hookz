//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var (
	version string = "1.0.0"
	config  Configuration
)

func TestDeps_ReadConfig(t *testing.T) {

	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	_, err := ReadConfig(afs, version)
	assert.Error(t, err, "There should be no config created so an error should be thrown.")
	config, _ = CreateConfig(afs, version)
	readConfig, err := ReadConfig(afs, version)

	assert.NoError(t, err, "ReadConfig should not have generated an error")
	assert.Equal(t, version, readConfig.Version, "Versions should match")

	_, err = ReadConfig(afs, "")
	assert.Error(t, err, "Passing an empty string should cause an error")
}
