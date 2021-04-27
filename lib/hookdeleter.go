//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

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
