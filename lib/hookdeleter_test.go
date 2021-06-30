//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"testing"

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
