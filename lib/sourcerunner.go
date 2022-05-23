package lib

import (
	"log"
	"os/exec"
)

//InstallSource installs a go repository that is found in the Sources section of the .hookz.yaml file.
func InstallSource(source Source) (err error) {

	log.Printf("installing: %s", source.Source)
	cmd := exec.Command("go", "install", source.Source)

	err = cmd.Run()
	if err != nil {
		log.Print(err)
	}

	return
}
