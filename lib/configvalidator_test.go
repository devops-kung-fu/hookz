package lib

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_generateShasum(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	_, _ = CreateConfig(afs, version)

	shasum := generateShasum(afs)
	assert.Equal(t, "d6e393b32ffa1a804b705d0a60acedd9c983a6d2e01cd1871a2e75ec358a5c20", shasum, "shasums do not match, but should")
}

func Test_WriteShasum(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	_, _ = CreateConfig(afs, version)

	err := WriteShasum(afs)
	assert.NoError(t, err, "A shasum write should not have caused an error")

	filename, _ := filepath.Abs(".git/hooks/hookz.shasum")

	exists, _ := afs.Exists(filename)
	assert.True(t, exists)

	contains, _ := afs.FileContainsBytes(filename, []byte("d6e393b32ffa1a804b705d0a60acedd9c983a6d2e01cd1871a2e75ec358a5c20"))
	assert.True(t, contains, "The expected shasum was not written to the hookz.shasum file")
}

func TestDeps_CheckVersion(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	_, _ = CreateConfig(afs, version)

	readConfig, err := ReadConfig(afs, version)
	assert.NoError(t, err, "ReadConfig should not have generated an error")

	err = ValidateVersion(readConfig, version)
	assert.NoError(t, err, "Check version should not have generated an error")

	err = ValidateVersion(readConfig, "2.0")
	assert.Error(t, err, "Version mismatch not caught")

	readConfig.Version = ""
	err = ValidateVersion(readConfig, version)
	assert.Error(t, err, "An empty config version should throw an error")
}
