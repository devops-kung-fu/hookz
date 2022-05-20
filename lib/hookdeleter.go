//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/devops-kung-fu/common/util"
	"github.com/spf13/afero"
)

//RemoveHooks purges all hooks from the filesystem that Hookz has created
//and deletes any generated scripts
func RemoveHooks(afs *afero.Afero, verbose bool) (err error) {
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
			util.DoIf(verbose, func() {
				util.PrintTabbed(fmt.Sprintf("Deleted %s", parts[len(parts)-1]))
			})
		}
	}

	removeShasum(afs)
	return
}

func removeShasum(afs *afero.Afero) {
	filename, _ := filepath.Abs(".git/hooks/hookz.shasum")
	_ = afs.Remove(filename)
}
