//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
)

//RemoveHooks purges all hooks from the filesystem that Hookz has created
//and deletes any generated scripts
func RemoveHooks(fs FileSystem, verbose bool) (err error) {
	DoIf(func() {
		color.Style{color.FgLightYellow}.Print("■")
		fmt.Println(" Removing existing hooks...")
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
			DoIf(func() {
				fmt.Printf("  Deleted %s\n", parts[len(parts)-1])
			}, verbose)
		}
	}

	removeShasum(fs)

	DoIf(func() {
		color.Style{color.FgGreen}.Print("■")
		fmt.Print(" Successfully removed existing hooks!\n")
	}, verbose)

	DoIf(func() {
		fmt.Println()
	}, verbose)

	return
}

func removeShasum(fs FileSystem) {
	filename, _ := filepath.Abs(".git/hooks/hookz.shasum")
	_ = fs.Afero().Remove(filename)
}
