//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/afero"
)

//RemoveHooks purges all hooks from the filesystem that Hookz has created
//and deletes any generated scripts
func RemoveHooks(afs *afero.Afero, verbose bool) (err error) {
	DoIf(verbose, func() {
		color.Style{color.FgLightYellow}.Print("■")
		fmt.Println(" Removing existing hooks...")
	})

	path, _ := os.Getwd()

	ext := ".hookz"
	p := fmt.Sprintf("%s/%s", path, ".git/hooks")

	dirFiles, _ := afs.ReadDir(p)

	for index := range dirFiles {
		file := dirFiles[index]

		name := file.Name()
		fullPath := fmt.Sprintf("%s/%s", p, name)
		info, _ := afs.Stat(fullPath)
		isHookzFile := strings.Contains(info.Name(), ext)
		if isHookzFile {
			var hookName = fullPath[0 : len(fullPath)-len(ext)]
			removeErr := afs.Fs.Remove(fullPath)
			if removeErr != nil {
				return removeErr
			}
			removeErr = afs.Fs.Remove(hookName)
			if removeErr != nil {
				return removeErr
			}
			parts := strings.Split(hookName, "/")
			DoIf(verbose, func() {
				fmt.Printf("  Deleted %s\n", parts[len(parts)-1])
			})
		}
	}

	removeShasum(afs)

	DoIf(verbose, func() {
		color.Style{color.FgGreen}.Print("■")
		fmt.Print(" Successfully removed existing hooks!\n")
	})

	DoIf(verbose, func() {
		fmt.Println()
	})

	return
}

func removeShasum(afs *afero.Afero) {
	filename, _ := filepath.Abs(".git/hooks/hookz.shasum")
	_ = afs.Remove(filename)
}
