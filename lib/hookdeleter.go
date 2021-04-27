package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func checkExt(ext string, pathS string) (files []string, err error) {
	filepath.Walk(pathS, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			match, _ := regexp.MatchString(ext, f.Name())
			if match {
				files = append(files, f.Name())
			}
		}
		return err
	})
	return files, nil
}

func RemoveHooks() (err error) {
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
			fmt.Printf("    	Deleted %s\n", parts[len(parts)-1])
		}
	}
	fmt.Println("[*] Successfully removed existing hooks!")

	return
}
