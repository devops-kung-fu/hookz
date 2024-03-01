package lib

import (
	"log"
	"os/exec"

	"gopkg.in/alessio/shellescape.v1"
)

// InstallSource installs a go repository that is found in the Sources section of the .hookz.yaml file.
func InstallSource(source Source) (err error) {
	cmd := exec.Command("go", "install", shellescape.Quote(source.Source))
	log.Println(cmd.String())
	err = cmd.Run()
	return
}
