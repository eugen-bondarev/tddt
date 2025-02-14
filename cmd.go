package main

import (
	"os/exec"
)

func executeCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	return cmd.Run()
}

type getDumpCommand interface {
	getCommand(database, outputFile string) string
}
