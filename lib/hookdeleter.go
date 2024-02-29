// Package lib Functionality for the Hookz CLI
package lib

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/devops-kung-fu/common/util"
	"github.com/spf13/afero"
)

// RemoveHooks removes hooks with a specific extension and their corresponding files in the Git hooks directory.
// It also optionally prints information about deleted hooks if verbose is set to true.
func RemoveHooks(afs *afero.Afero, verbose bool) error {
	path, _ := os.Getwd()

	ext := ".hookz"
	hooksPath := filepath.Join(path, ".git/hooks")

	dirFiles, err := afs.ReadDir(hooksPath)
	if err != nil {
		return err
	}

	for _, file := range dirFiles {
		fullPath := filepath.Join(hooksPath, file.Name())
		if strings.Contains(file.Name(), ext) {
			if removeErr := afs.Remove(fullPath); removeErr != nil {
				return removeErr
			}

			correspondingFile := fullPath[:len(fullPath)-len(ext)]
			if removeErr := afs.Remove(correspondingFile); removeErr != nil {
				return removeErr
			}

			util.DoIf(verbose, func() {
				util.PrintTabbedf("Deleted %s\n", filepath.Base(correspondingFile))
			})
		}
	}

	removeShasum(afs)
	return nil
}

// removeShasum removes the shasum file from the Git hooks directory.
func removeShasum(afs *afero.Afero) {
	filename, _ := filepath.Abs(".git/hooks/hookz.shasum")
	_ = afs.Remove(filename)
}
