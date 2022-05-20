package lib

import (
	"log"
	"os/exec"
)

func InstallSources(sources []Source) (err error) {
	for _, s := range sources {
		log.Printf("installing: %s", s.Source)
		cmd := exec.Command("go", "install", s.Source)

		err = cmd.Run()
		if err != nil {
			log.Print(err)
		}
	}
	return
}
