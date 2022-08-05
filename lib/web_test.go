// Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestUpdateExecutables(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}
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
	_, err := UpdateExecutables(afs, botchedConfig)
	assert.NoError(t, err, "UpdateExecutables should only happen if action.URL != nil")
}

func Test_DownloadFile(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}
	_, err := DownloadFile(afs, "x", "x")
	assert.Error(t, err, "URL should be a valid URI")
}

// func Test_DownloadURLWithPlatform(t *testing.T) {
// 	URL := "https://github.com/devops-kung-fu/hinge/releases/download/v0.1.0/hinge-0.1.0-%%PLATFORM%%-amd64"
// 	_, _ = DownloadURL(URL)
// }

func Test_getPlatformName(t *testing.T) {
	platform := getPlatformName()
	assert.NotEmpty(t, platform, "There should be a platform returned")
}

func Test_platformURLIfDefined(t *testing.T) {
	processedURL := platformURLIfDefined("https://%%PLATFORM%%")
	assert.NotContains(t, processedURL, "%%PLATFORM%%", "The token %%PLATFORM%% should not exist in the return")
}

func TestWriteCounter_Write(t *testing.T) {
	wc := WriteCounter{}
	count, err := wc.Write([]byte("test"))
	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, 4, count, "4 bytes should have been written")
}

func TestDownloadFile(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}
	URL := "https://github.com/devops-kung-fu/hinge/releases/download/v0.1.0/hinge-0.1.0-linux-amd64"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", URL,
		httpmock.NewBytesResponder(200, []byte("test")))

	filename, err := DownloadFile(afs, ".git/hooks", URL)
	assert.NoError(t, err)
	assert.Equal(t, "hinge-0.1.0-linux-amd64", filename)

	httpmock.GetTotalCallCount()
}
