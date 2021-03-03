package main

import (
	"os"

	"github.com/devops-kung-fu/hookz/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
