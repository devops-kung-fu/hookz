// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	//Verbose identifies if extended output should be configured during init and reset
	Version float64
	Verbose bool
	rootCmd = &cobra.Command{
		Use:     "hookz",
		Short:   `Manages commit hooks inside a local git repository`,
		Version: "0.0.1",
	}
)

// Execute creates the command tree and handles any error condition returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	Version = 1.1
}

func readConfig() (config Configuration, err error) {

	filename, _ := filepath.Abs(".hookz.yaml")
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		if err != nil {
			return
		}
	}

	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return
	}
	return
}

func hookzHeader() {
	fmt.Println("Hookz (https://github.com/devops-kung-fu/hookz)\n")
}

func isError(err error, pre string) error {
	if err != nil {
		log.Printf("%v: %v", pre, err)
	}
	return err
}

func isErrorBool(err error, pre string) (b bool) {
	if err != nil {
		log.Printf("%v: %v", pre, err)
		b = true
	}
	return
}

func checkExt(ext string, pathS string) (files []string, err error) {
	filepath.Walk(pathS, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return err
	})
	return files, nil
}

func removeHooks() (err error) {
	ext := ".hookz"
	p := ".git/hooks/"

	dirRead, _ := os.Open(p)
	dirFiles, _ := dirRead.Readdir(0)

	for index := range dirFiles {
		file := dirFiles[index]

		name := file.Name()
		fullPath := fmt.Sprintf("%s%s", p, name)

		r, err := regexp.MatchString(ext, fullPath)
		if err == nil && r {
			os.Remove(fullPath)
			var hookName = fullPath[0 : len(fullPath)-len(ext)]
			os.Remove(hookName)
			parts := strings.Split(hookName, "/")
			fmt.Println(fmt.Sprintf("[*] Deleted %s", parts[len(parts)-1]))
		}
	}

	return
}

func downloadURL(url string) (filename string, err error) {
	client := grab.NewClient()
	req, _ := grab.NewRequest(".git/hooks", url)

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
