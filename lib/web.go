//Package lib Functionality for the Hookz CLI
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
)

//WriteCounter encapsulates the total number of bytes captured and rendered
type WriteCounter struct {
	Total    uint64
	FileName string
}

//Write increments the total number of bytes and prints progress to STDOUT
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

//PrintProgress prints the current download progress to STDOUT
func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\r  Downloading %s... %s complete", wc.FileName, humanize.Bytes(wc.Total))
}

//UpdateExecutables parses the configuration for URL's and re-downloads
//the contents into the .git/hooks folder
func UpdateExecutables(fs FileSystem, config Configuration) (err error) {
	var updateCount = 0
	for _, hook := range config.Hooks {
		for _, action := range hook.Actions {
			if action.URL != nil {
				updateCount++
				_, err = DownloadFile(fs, ".git/hooks", *action.URL)
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
func DownloadFile(fs FileSystem, filepath string, URL string) (filename string, err error) {
	URL = platformURLIfDefined(URL)
	_, err = url.ParseRequestURI(URL)
	if err != nil {
		return
	}
	filename = path.Base(URL)
	fullFileName := fmt.Sprintf("%s/%s", filepath, filename)
	out, err := fs.Afero().Create(fmt.Sprintf("%s.tmp", fullFileName))
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

// func DownloadURL(URL string) (filename string, err error) {
// 	URL = platformURLIfDefined(URL)
// 	_, err = url.ParseRequestURI(URL)
// 	if err != nil {
// 		return
// 	}
// 	client := grab.NewClient()
// 	req, err := grab.NewRequest(".git/hooks", URL)
// 	if err != nil {
// 		return
// 	}

// 	fmt.Printf("Downloading %v...\n", req.URL())
// 	resp := client.Do(req)
// 	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

// 	t := time.NewTicker(500 * time.Millisecond)
// 	defer t.Stop()

// Loop:
// 	for {
// 		select {
// 		case <-t.C:
// 			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
// 				resp.BytesComplete(),
// 				resp.Size,
// 				100*resp.Progress())

// 		case <-resp.Done:
// 			break Loop
// 		}
// 	}

// 	if err := resp.Err(); err != nil {
// 		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
// 		return resp.Filename, err
// 	}

// 	fmt.Printf("Download saved to ./%v \n", resp.Filename)
// 	err = os.Chmod(resp.Filename, 0777)
// 	if err != nil {
// 		return resp.Filename, err
// 	}
// 	return resp.Filename, err
// }

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
