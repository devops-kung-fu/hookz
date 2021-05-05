//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeps_CreateScriptFile(t *testing.T) {
	content := "Test Script"
	filename, err := f.CreateScriptFile(content)
	assert.NoError(t, err, "CreateScriptFile should not have generated an error")
	assert.NotEmpty(t, filename, "A filename should have been returned")

	path, _ := os.Getwd()
	fullFileName := fmt.Sprintf("%s/%s/%s", path, ".git/hooks", filename)
	contains, _ := f.Afero().FileContainsBytes(fullFileName, []byte(content))
	assert.True(t, contains, "Script file should have the phrase `Test Script` in it")
}

func Test_genTemplate(t *testing.T) {
	content := "hooktype"
	template := genTemplate(content)
	assert.Equal(t, content, template.ParseName, "Template should have a name `hooktype`")
}

func Test_buildFullCommand(t *testing.T) {
	config, err := f.createConfig(version)
	assert.NoError(t, err, "createConfig should not have generated an error")

	action := config.Hooks[0].Actions[0]
	assert.NotNil(t, action, "Action should not be nil")

	command := buildFullCommand(action, true)
	assert.Equal(t, "echo -e Hello Hookz!", command, "Values are not equal")
}

func Test_WriteHooks(t *testing.T) {
	config, err := f.createConfig(version)
	assert.NoError(t, err, "createConfig should not have generated an error")

	err = f.WriteHooks(config, true)
	assert.NoError(t, err, "WriteHooks should not have generated an error")

	filename, _ := filepath.Abs(".git/hooks/pre-commit")
	contains, _ := f.Afero().FileContainsBytes(filename, []byte("Hookz"))
	assert.True(t, contains, "Generated hook should have the word Hookz in it")
}

func Test_createFile(t *testing.T) {
	err := f.CreateFile("test")
	assert.NoError(t, err, "Create file should not generate an error")
	// assert.FileExists(t, "./test", "A file should have been created")

	// err = f.CreateFile("")
	// assert.Error(t, err, "A file should have not been created and an error thrown")

}

func Test_writeTemplate(t *testing.T) {
	err := f.writeTemplate(nil, "")
	assert.Error(t, err, "writeTemplate should throw an error if there is no file created")
}
