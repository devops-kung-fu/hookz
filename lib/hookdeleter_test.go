//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFileSystem_RemoveHooks(t *testing.T) {
	path, _ := os.Getwd()

	content := "Test Script"
	CreateScriptFile(fs, content)

	p := fmt.Sprintf("%s/%s", path, ".git/hooks")
	dirFiles, _ := fs.Afero().ReadDir(p)
	assert.Equal(t, 2, len(dirFiles), "Incorrect number of created script files")

	err := RemoveHooks(fs, true)
	assert.NoError(t, err, "RemoveHooks should not have generated an error")

	dirFiles, _ = fs.Afero().ReadDir(p)
	assert.Equal(t, 0, len(dirFiles), "Incorrect number of created script files")

}

func Test_removeShasum(t *testing.T) {
	newFs := FileSystem{
		fs: afero.NewMemMapFs(),
	}

	CreateConfig(newFs, version)
	_ = WriteShasum(newFs)

	path, _ := os.Getwd()
	filename := fmt.Sprintf("%s/%s", path, ".git/hooks/hookz.shasum")
	exists, _ := newFs.Afero().Exists(filename)

	assert.True(t, exists, "A shasum file needs to exist in order to test removal")

	removeShasum(newFs)

	exists, _ = newFs.Afero().Exists(filename)
	assert.False(t, exists, "The shasum file should no longer exist")

}
