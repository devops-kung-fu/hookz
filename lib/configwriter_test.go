// Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_CreateConfig(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	exists := HasExistingHookzYaml(afs)
	assert.False(t, exists, "No hookz.yaml file should exist")

	_, err := CreateConfig(afs, version)
	assert.NoError(t, err, "CreateConfig should not have generated an error")

	exists = HasExistingHookzYaml(afs)
	assert.True(t, exists, "A hookz.yaml file should exist")

	readConfig, _ := ReadConfig(afs, version)
	assert.Equal(t, version, readConfig.Version, "A Version should have been detected in the config after creation.")
}
