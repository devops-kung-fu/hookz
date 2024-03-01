package cmd

import (
	"testing"

	"github.com/devops-kung-fu/common/util"
	"github.com/stretchr/testify/assert"

	"github.com/devops-kung-fu/hookz/lib"
)

func Test_InstallSources(t *testing.T) {
	sources := []lib.Source{
		{
			Source: "github.com/devops-kung-fu/hinge@latest",
		},
	}
	output := util.CaptureOutput(func() {
		_ = InstallSources(sources)
	})

	assert.NotNil(t, output)
	assert.Contains(t, output, "go install github.com/devops-kung-fu/hinge@latest\n")

	sources = []lib.Source{
		{
			Source: "yeah",
		},
	}
	output = util.CaptureOutput(func() {
		_ = InstallSources(sources)
	})
	assert.Contains(t, output, "exit status 1\n")
}

func TestNoConfig(t *testing.T) {
	output := util.CaptureOutput(func() {
		NoConfig()
	})
	assert.NotNil(t, output)
}
