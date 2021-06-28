//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestDeps_CreateScriptFile(t *testing.T) {
	content := "Test Script"
	filename, err := CreateScriptFile(fs, content)
	assert.NoError(t, err, "CreateScriptFile should not have generated an error")
	assert.NotEmpty(t, filename, "A filename should have been returned")

	path, _ := os.Getwd()
	fullFileName := fmt.Sprintf("%s/%s/%s", path, ".git/hooks", filename)
	contains, _ := fs.Afero().FileContainsBytes(fullFileName, []byte(content))
	assert.True(t, contains, "Script file should have the phrase `Test Script` in it")
}

func Test_genTemplate(t *testing.T) {
	content := "hooktype"
	template := genTemplate(content)
	assert.Equal(t, content, template.ParseName, "Template should have a name `hooktype`")
}

func Test_buildFullCommand(t *testing.T) {
	config, err := CreateConfig(fs, version)
	assert.NoError(t, err, "createConfig should not have generated an error")

	action := config.Hooks[0].Actions[0]
	assert.NotNil(t, action, "Action should not be nil")

	command := buildFullCommand(action, true)
	assert.Equal(t, "echo -e Hello Hookz!", command, "Values are not equal")
}

func Test_WriteHooks(t *testing.T) {
	config, err := CreateConfig(fs, version)
	assert.NoError(t, err, "createConfig should not have generated an error")

	err = WriteHooks(fs, config, true, true)
	assert.NoError(t, err, "WriteHooks should not have generated an error")

	filename, _ := filepath.Abs(".git/hooks/pre-commit")
	contains, _ := fs.Afero().FileContainsBytes(filename, []byte("Hookz"))
	assert.True(t, contains, "Generated hook should have the word Hookz in it")
}

func Test_createFile(t *testing.T) {
	err := CreateFile(fs, "test")
	assert.NoError(t, err, "Create file should not generate an error")
	exists, _ := fs.Afero().Exists("test")
	assert.True(t, exists, "A file should have been created")
}

func Test_writeTemplate(t *testing.T) {
	err := writeTemplate(fs, nil, "")
	assert.Error(t, err, "writeTemplate should throw an error if there is no file created")
}

func Test_HasExistingHookz(t *testing.T) {
	exists := HasExistingHookz(fs)
	assert.False(t, exists, "No hookz files should exist")

	config, err := CreateConfig(fs, version)
	assert.NoError(t, err, "createConfig should not have generated an error")

	err = WriteHooks(fs, config, true, true)
	assert.NoError(t, err, "WriteHooks should not have generated an error")

	exists = HasExistingHookz(fs)
	assert.True(t, exists, "hookz files should exist")

}

func Test_buildExec(t *testing.T) {
	newFs := FileSystem{
		fs: afero.NewMemMapFs(),
	}

	script := "#!/bin/bash"
	action := Action{
		Script: &script,
	}
	err := buildExec(newFs, &action)

	assert.NoError(t, err, "buildExec shouldn't have generated an error")
	assert.NotNil(t, action.Exec, "action.Exec field should not be nil")
}
