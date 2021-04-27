//Package lib Functionality for the Hookz CLI
package lib

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func ReadConfig(version string) (config Configuration, err error) {

	filename, _ := filepath.Abs(".hookz.yaml")
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		if err != nil {
			return
		}
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return
	}

	if config.Version == "" {
		err = errors.New("no configuration version value found in .hookz.yaml")
		return
	}

	// Check version
	ver := strings.Split(config.Version, ".")
	verMatch := strings.Split(version, ".")
	if fmt.Sprintf("%v.%v", ver[0], ver[1]) != fmt.Sprintf("%v.%v", verMatch[0], verMatch[1]) {
		err = fmt.Errorf("version mismatch: Expected v%v.%v - Check your .hookz.yaml configuration", verMatch[0], verMatch[1])
	}
	return
}
