//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/dustin/go-humanize"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

//UpdateExecutables parses the configuration for URL's and re-downloads
//the contents into the .git/hooks folder
func UpdateExecutables(fs FileSystem, config Configuration) (err error) {
	var updateCount = 0
	for _, hook := range config.Hooks {
		for _, action := range hook.Actions {
			if action.URL != nil {
				updateCount++
				_, err = DownloadURL(*action.URL)
			}
		}
	}
	if updateCount == 0 {
		fmt.Println("Nothing to Update!")
	}

	return
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	fmt.Print("\n")

	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}

//DownloadURL downloads content from the provided URL and returns the
//filename after saving the content to the .git/hooks folder. Returns an
//error if there were any problems.
func DownloadURL(URL string) (filename string, err error) {
	URL = platformURLIfDefined(URL)
	_, err = url.ParseRequestURI(URL)
	if err != nil {
		return
	}
	client := grab.NewClient()
	req, err := grab.NewRequest(".git/hooks", URL)
	if err != nil {
		return
	}

	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		return resp.Filename, err
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
	err = os.Chmod(resp.Filename, 0777)
	if err != nil {
		return resp.Filename, err
	}
	return resp.Filename, err
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
