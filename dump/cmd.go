package dump

import (
	"fmt"
	"os/exec"
	"strings"
)

func executeCommand(command string) error {
	c := exec.Command("bash", "-c", command)
	var stderr strings.Builder
	c.Stderr = &stderr
	if err := c.Run(); err != nil {
		return fmt.Errorf("%w: %s", err, strings.TrimSpace(stderr.String()))
	}
	return nil
}

type GetDumpCommand interface {
	GetCommand(database, outputFile string) string
}
