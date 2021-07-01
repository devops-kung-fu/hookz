package lib

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

func generateShasum(fs FileSystem) (shasum string, err error) {
	filename, _ := filepath.Abs(".hookz.yaml")
	yamlFile, err := fs.Afero().ReadFile(filename)
	if err != nil {
		return
	}
	shasum = fmt.Sprintf("%x", sha256.Sum256(yamlFile))
	return
}

//WriteShasum writes the shasum of the JSON representation of the configuration to hookz.shasum
func WriteShasum(fs FileSystem) (err error) {
	shasum, err := generateShasum(fs)
	if err != nil {
		return err
	}
	filename, _ := filepath.Abs(".git/hooks/hookz.shasum")
	err = fs.Afero().WriteFile(filename, []byte(shasum), 0644)
	return
}

//ValidateVersion ensures that the configuration that is read matches the hookz binary version
func ValidateVersion(config Configuration, version string) (err error) {
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
