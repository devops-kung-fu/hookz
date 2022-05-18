package lib

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_generateShasum(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	CreateConfig(afs, version)

	shasum, err := generateShasum(afs)
	assert.NoError(t, err, "Likely .hookz.yaml couldn't be read")
	assert.Equal(t, "0213f04ee70cc7b48d6de58e7dd62338259d16ae8e52016d19f83559051dd57c", shasum, "shasums do not match, but should")
}

func Test_WriteShasum(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	CreateConfig(afs, version)

	err := WriteShasum(afs)
	assert.NoError(t, err, "A shasum write should not have caused an error")

	filename, _ := filepath.Abs(".git/hooks/hookz.shasum")

	exists, _ := afs.Exists(filename)
	assert.True(t, exists)

	contains, _ := afs.FileContainsBytes(filename, []byte("0213f04ee70cc7b48d6de58e7dd62338259d16ae8e52016d19f83559051dd57c"))
	assert.True(t, contains, "The expected shasum was not written to the hookz.shasum file")
}

func TestDeps_CheckVersion(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	CreateConfig(afs, version)

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
