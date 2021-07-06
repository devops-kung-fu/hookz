package lib

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_generateShasum(t *testing.T) {
	newFs := FileSystem{
		fs: afero.NewMemMapFs(),
	}

	CreateConfig(newFs, version)

	shasum, err := generateShasum(newFs)
	assert.NoError(t, err, "Likely .hookz.yaml couldn't be read")
	assert.Equal(t, "bae340195673fed855200d34dace2fc3c6ed44ff37678a2a941b1e58557882e9", shasum, "shasums do not match, but should")
}

func Test_WriteShasum(t *testing.T) {
	newFs := FileSystem{
		fs: afero.NewMemMapFs(),
	}

	CreateConfig(newFs, version)

	err := WriteShasum(newFs)
	assert.NoError(t, err, "A shasum write should not have caused an error")

	path, _ := os.Getwd()
	filename := fmt.Sprintf("%s/%s", path, ".git/hooks/hookz.shasum")

	contains, _ := newFs.Afero().FileContainsBytes(filename, []byte("bae340195673fed855200d34dace2fc3c6ed44ff37678a2a941b1e58557882e9"))
	assert.True(t, contains, "The expected shasum was not written to the hookz.shasum file")
}

func TestDeps_CheckVersion(t *testing.T) {
	newFs := FileSystem{
		fs: afero.NewMemMapFs(),
	}

	CreateConfig(newFs, version)

	readConfig, err := ReadConfig(newFs, version)
	assert.NoError(t, err, "ReadConfig should not have generated an error")

	err = ValidateVersion(readConfig, version)
	assert.NoError(t, err, "Check version should not have generated an error")

	err = ValidateVersion(readConfig, "2.0")
	assert.Error(t, err, "Version mismatch not caught")

	readConfig.Version = ""
	err = ValidateVersion(readConfig, version)
	assert.Error(t, err, "An empty config version should throw an error")
}
