//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"strings"
)

func RemoveHooks(fs FileSystem, verbose bool) (err error) {
	PrintIf(func() {
		fmt.Println("[*] Removing existing hooks...")
	}, verbose)

	path, _ := os.Getwd()

	ext := ".hookz"
	p := fmt.Sprintf("%s/%s", path, ".git/hooks")

	dirFiles, _ := fs.Afero().ReadDir(p)

	for index := range dirFiles {
		file := dirFiles[index]

		name := file.Name()
		fullPath := fmt.Sprintf("%s/%s", p, name)
		info, _ := fs.Afero().Stat(fullPath)
		isHookzFile := strings.Contains(info.Name(), ext)
		if isHookzFile {
			var hookName = fullPath[0 : len(fullPath)-len(ext)]
			removeErr := fs.fs.Remove(fullPath)
			if removeErr != nil {
				return removeErr
			}
			removeErr = fs.fs.Remove(hookName)
			if removeErr != nil {
				return removeErr
			}
			parts := strings.Split(hookName, "/")
			PrintIf(func() {
				fmt.Printf("    	Deleted %s\n", parts[len(parts)-1])
			}, verbose)
		}
	}
	PrintIf(func() {
		fmt.Println("[*] Successfully removed existing hooks!")
	}, verbose)

	PrintIf(func() {
		fmt.Println()
	}, verbose)

	return
}
