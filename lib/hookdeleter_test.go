// Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFileSystem_RemoveHooks(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}
	path, _ := os.Getwd()

	content := "Test Script"
	_, _ = CreateScriptFile(afs, content)

	p := fmt.Sprintf("%s/%s", path, ".git/hooks")
	dirFiles, _ := afs.ReadDir(p)
	assert.Equal(t, 2, len(dirFiles), "Incorrect number of created script files")

	err := RemoveHooks(afs, true)
	assert.NoError(t, err, "RemoveHooks should not have generated an error")

	dirFiles, _ = afs.ReadDir(p)
	assert.Equal(t, 0, len(dirFiles), "Incorrect number of created script files")

}

func Test_removeShasum(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}

	_, _ = CreateConfig(afs, version)
	_ = WriteShasum(afs)

	path, _ := os.Getwd()
	filename := fmt.Sprintf("%s/%s", path, ".git/hooks/hookz.shasum")
	exists, _ := afs.Exists(filename)

	assert.True(t, exists, "A shasum file needs to exist in order to test removal")

	removeShasum(afs)

	exists, _ = afs.Exists(filename)
	assert.False(t, exists, "The shasum file should no longer exist")

}
