package dump

import (
	"os/exec"
)

func executeCommand(command string) error {
	c := exec.Command("bash", "-c", command)
	return c.Run()
}

type GetDumpCommand interface {
	GetCommand(database, outputFile string) string
}
