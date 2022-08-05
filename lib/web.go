// Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/spf13/afero"
)

// WriteCounter encapsulates the total number of bytes captured and rendered
type WriteCounter struct {
	Total    uint64
	FileName string
}

// Write increments the total number of bytes and prints progress to STDOUT
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

// PrintProgress prints the current download progress to STDOUT
func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\r	Downloading %s... %s complete", wc.FileName, humanize.Bytes(wc.Total))
}

// UpdateExecutables parses the configuration for URL's and re-downloads
// the contents into the .git/hooks folder
func UpdateExecutables(afs *afero.Afero, config Configuration) (updateCount int, err error) {
	for _, hook := range config.Hooks {
		for _, action := range hook.Actions {
			if action.URL != nil {
				updateCount++
				_, err = DownloadFile(afs, ".git/hooks", *action.URL)
			}
		}
	}
	return
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFile(afs *afero.Afero, filepath string, URL string) (filename string, err error) {
	URL = platformURLIfDefined(URL)
	_, err = url.ParseRequestURI(URL)
	if err != nil {
		return
	}
	filename = path.Base(URL)
	fullFileName := fmt.Sprintf("%s/%s", filepath, filename)
	out, err := afs.Create(fmt.Sprintf("%s.tmp", fullFileName))
	if err != nil {
		return
	}
	defer func() {
		err = out.Close()
	}()

	resp, err := http.Get(URL)
	if err != nil {
		return
	}
	defer func() {
		err = resp.Body.Close()
	}()

	counter := &WriteCounter{FileName: filename}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return
	}

	fmt.Print("\n")

	err = os.Rename(fmt.Sprintf("%s.tmp", fullFileName), fullFileName)
	if err != nil {
		return
	}

	err = os.Chmod(fullFileName, 0777)
	filename = fullFileName

	return
}

func platformURLIfDefined(URL string) string {
	return strings.Replace(URL, "%%PLATFORM%%", getPlatformName(), 1)
}

func getPlatformName() string {
	platform := strings.ToLower(runtime.GOOS)
	if platform == "penguin" {
		platform = "linux"
	}
	return platform
}
