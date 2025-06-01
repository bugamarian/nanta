package common

import (
	"log"
	"os"
	"os/exec"
)

func OpenFile(modifier, path string) {
	// Generic, takes any command as arg, for example: cat, nvim, glow
	cmd := exec.Command(modifier, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
}
