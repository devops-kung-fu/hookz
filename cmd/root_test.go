// cmd/root_test.go
package cmd

import (
	"testing"

	"github.com/devops-kung-fu/common/util"
	"github.com/stretchr/testify/assert"
)

func Test_rootCommand(t *testing.T) {
	// Set up a test command
	cmd := updateCmd

	// Capture standard output for testing
	output := util.CaptureOutput(func() {
		_ = cmd.Execute()
	})

	assert.Contains(t, output, "Manages commit hooks inside a local git repository")
}
