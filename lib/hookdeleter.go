//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func (f FileSystem) RemoveHooks() (err error) {
	ext := ".hookz"
	p, _ := filepath.Abs(fmt.Sprintf(".git/hooks/"))

	dirFiles, err := f.Afero().ReadDir(p)
	if err != nil {
		return err
	}

	for index := range dirFiles {
		file := dirFiles[index]

		name := file.Name()
		fullPath := fmt.Sprintf("%s/%s", p, name)
		r, err := regexp.MatchString(ext, fullPath)
		if err == nil && r {
			removeErr := f.fs.Remove(fullPath)
			if removeErr != nil {
				return removeErr
			}
			var hookName = fullPath[0 : len(fullPath)-len(ext)]
			removeErr = f.fs.Remove(hookName)
			if removeErr != nil {
				return removeErr
			}
			parts := strings.Split(hookName, "/")
			fmt.Printf("    	Deleted %s\n", parts[len(parts)-1])
		}
	}
	fmt.Println("[*] Successfully removed existing hooks!")

	return
}
