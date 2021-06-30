//Package lib Functionality for the Hookz CLI
package lib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

//HasExistingHookzYaml returns true if an existing .hookz.yaml file exists
func HasExistingHookzYaml(fs FileSystem) bool {
	filename, _ := filepath.Abs(".hookz.yaml")
	_, readErr := fs.Afero().ReadFile(filename)
	return readErr == nil
}

//ReadConfig reads the .hookz.yaml file in from the filesystem and
//ensures it matches the provided version
func ReadConfig(fs FileSystem, version string) (config Configuration, err error) {

	filename, _ := filepath.Abs(".hookz.yaml")
	yamlFile, readErr := fs.Afero().ReadFile(filename)
	if readErr != nil {
		fmt.Println(".hookz.yaml file not found")
		fmt.Println("\nTo create a sample configuration run:")
		fmt.Println("        hookz init config")
		fmt.Println("\nRun 'hookz --help' for usage.")
		fmt.Println()
		os.Exit(1)
	} else {
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			return
		}
	}

	err = checkVersion(config, version)

	return
}

func checkVersion(config Configuration, version string) (err error) {
	if config.Version == "" {
		err = errors.New("no configuration version value found in .hookz.yaml")
		return
	}
	if version == "" {
		err = errors.New("a version should not be empty")
		return
	}
	ver := strings.Split(config.Version, ".")
	verMatch := strings.Split(version, ".")
	if fmt.Sprintf("%v.%v", ver[0], ver[1]) != fmt.Sprintf("%v.%v", verMatch[0], verMatch[1]) {
		err = fmt.Errorf("version mismatch: Expected v%v.%v - Check your .hookz.yaml configuration", verMatch[0], verMatch[1])
	}
	return
}
