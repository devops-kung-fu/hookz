package lib

import (
	"testing"

	"github.com/devops-kung-fu/common/util"
	"github.com/stretchr/testify/assert"
)

func TestInstallSources(t *testing.T) {
	sources := []Source{
		{
			Source: "github.com/devops-kung-fu/hinge@latest",
		},
	}
	output := util.CaptureOutput(func() {
		_ = InstallSources(sources)
	})

	assert.NotNil(t, output)
	assert.Contains(t, output, "installing: github.com/devops-kung-fu/hinge@latest\n")
}
