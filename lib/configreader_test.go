// Package lib Functionality for the Hookz CLI
package lib

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var (
	version string = "1.0.0"
)

func Test_ReadConfig(t *testing.T) {

	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	_, err := ReadConfig(afs, version)
	assert.Error(t, err, "There should be no config created so an error should be thrown.")
	assert.Equal(t, "NO_CONFIG", err.Error())

	_, _ = CreateConfig(afs, "1.0.0")
	readConfig, err := ReadConfig(afs, version)

	assert.NoError(t, err, "ReadConfig should not have generated an error")
	assert.Equal(t, version, readConfig.Version, "Versions should match")

	readConfig, err = ReadConfig(afs, "0.1.0")
	assert.Error(t, err)

	_, err = ReadConfig(afs, "")
	assert.Error(t, err, "Passing an empty string should cause an error")
}

func Test_badConfig(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}
	filename, _ := filepath.Abs(".hookz.yaml")
	_ = afs.WriteFile(filename, badConfig(), 0644)

	_, err := ReadConfig(afs, version)
	assert.Error(t, err)
	assert.Equal(t, "BAD_YAML", err.Error())
}

func badConfig() []byte {
	config := `
	version: 2.4.2

	hooks:
	- type: pre-commit
		actions:
		- name: "Git Pull (Ensure there are no upstream changes)"
			exec: git
			args: ["pull"]
	- type: post-commit
		actions:
		- name: "Mark all done"
			exec: echo
			args: ["-e" "[x] Successfully committed upstream"]
	`
	return []byte(config)
}
