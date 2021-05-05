//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateExecutables(t *testing.T) {
	botchedConfig := Configuration{
		Version: "s",
		Hooks: []Hook{
			{
				Type: "pre-commit",
				Actions: []Action{
					{
						Name: "test",
					},
				},
			},
		},
	}
	err := f.UpdateExecutables(botchedConfig)
	assert.NoError(t, err, "UpdateExecutables should only happen if action.URL != nil")
}

func Test_DownloadURL(t *testing.T) {
	_, err := f.DownloadURL("x")
	assert.Error(t, err, "URL should be a valid URI")
}
