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
	err := UpdateExecutables(fs, botchedConfig)
	assert.NoError(t, err, "UpdateExecutables should only happen if action.URL != nil")
}

func Test_DownloadURL(t *testing.T) {
	_, err := DownloadURL("x")
	assert.Error(t, err, "URL should be a valid URI")
}

// func Test_DownloadURLWithPlatform(t *testing.T) {
// 	URL := "https://github.com/devops-kung-fu/hinge/releases/download/v0.1.0/hinge-0.1.0-%%PLATFORM%%-amd64"
// 	_, _ = DownloadURL(URL)
// }
