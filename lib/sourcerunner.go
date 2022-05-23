package lib

import (
	"log"
	"os/exec"
)

func InstallSource(source Source) (err error) {

	log.Printf("installing: %s", source.Source)
	cmd := exec.Command("go", "install", source.Source)

	err = cmd.Run()
	if err != nil {
		log.Print(err)
	}

	return
}
