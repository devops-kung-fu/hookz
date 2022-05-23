//Package lib Functionality for the Hookz CLI
package lib

import (
	"errors"
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

//HasExistingHookzYaml returns true if an existing .hookz.yaml file exists
func HasExistingHookzYaml(afs *afero.Afero) bool {
	filename, _ := filepath.Abs(".hookz.yaml")
	_, readErr := afs.ReadFile(filename)
	return readErr == nil
}

//ReadConfig reads the .hookz.yaml file in from the filesystem and
//ensures it matches the provided version
func ReadConfig(afs *afero.Afero, version string) (config Configuration, err error) {
	filename, _ := filepath.Abs(".hookz.yaml")
	yamlFile, readErr := afs.ReadFile(filename)
	if readErr != nil {
		err = errors.New("NO_CONFIG")
		return
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return
	}

	err = ValidateVersion(config, version)

	return
}
