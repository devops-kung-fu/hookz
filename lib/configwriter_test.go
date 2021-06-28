//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createConfig(t *testing.T) {
	_, err := createConfig(fs, version)
	assert.NoError(t, err, "ReadConfig should not have generated an error")

	readConfig, _ := ReadConfig(fs, version)
	assert.Equal(t, version, readConfig.Version, "A Version should have been detected in the config after creation.")
}
