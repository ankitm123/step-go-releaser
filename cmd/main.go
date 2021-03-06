package main

import (
	"os"

	"github.com/jenkins-x/step-goreleaser/pkg"
)

// Entrypoint for the command
func main() {
	err := pkg.NewCmdGoReleaser().Execute()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
