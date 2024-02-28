// Package lib Functionality for the Hookz CLI
package lib

import (
	"errors"
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

// HasExistingHookzYaml returns true if an existing .hookz.yaml file exists
func HasExistingHookzYaml(afs *afero.Afero) bool {
	filename, _ := filepath.Abs(".hookz.yaml")
	_, readErr := afs.ReadFile(filename)
	return readErr == nil
}

// ReadConfig method reads a file named .hookz.yaml from the filesystem using the provided Afero interface. It then checks if the version specified in the function parameter matches the version specified in the configuration file. If the versions match, the configuration file is parsed using YAML and the resulting Configuration struct is returned. If the versions don't match or there is an error while reading or parsing the file, an error is returned.
func ReadConfig(afs *afero.Afero, version string) (config Configuration, err error) {
	filename, _ := filepath.Abs(".hookz.yaml")
	yamlFile, readErr := afs.ReadFile(filename)
	if readErr != nil {
		err = errors.New("NO_CONFIG")
		return
	}

	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		err = errors.New("BAD_YAML")
		return
	}

	err = ValidateVersion(config, version)

	return
}
