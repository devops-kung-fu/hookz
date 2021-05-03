//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeps_CreateScriptFile(t *testing.T) {
	content := "Test Script"
	filename, err := deps.CreateScriptFile(content)
	assert.NoError(t, err, "CreateScriptFile should not have generated an error")
	assert.NotEmpty(t, filename, "A filename should have been returned")
}

func Test_genTemplate(t *testing.T) {
	content := "hooktype"
	template := genTemplate(content)
	assert.Equal(t, content, template.ParseName, "Template should have a name `hooktype`")
}

func Test_buildFullCommand(t *testing.T) {
	config, err := deps.createConfig(version)
	assert.NoError(t, err, "createConfig should not have generated an error")

	action := config.Hooks[0].Actions[0]
	assert.NotNil(t, action, "Action should not be nil")

	command := buildFullCommand(action, true)
	assert.Equal(t, "echo -e Hello Hookz!", command, "Values are not equal")
}
