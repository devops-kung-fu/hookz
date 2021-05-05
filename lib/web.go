//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
)

//UpdateExecutables parses the configuration for URL's and re-downloads
//the contents into the .git/hooks folder
func (f FileSystem) UpdateExecutables(config Configuration) (err error) {
	var updateCount = 0
	for _, hook := range config.Hooks {
		for _, action := range hook.Actions {
			if action.URL != nil {
				updateCount++
				_, _ = f.DownloadURL(*action.URL)
			}
		}
	}
	if updateCount == 0 {
		fmt.Println("Nothing to Update!")
	}

	return
}

//DownloadURL downloads content from the provided URL and returns the
//filename after saving the content to the .git/hooks folder. Returns an
//error if there were any problems.
func (f FileSystem) DownloadURL(URL string) (filename string, err error) {
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
