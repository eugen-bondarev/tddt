package dump

import (
	"fmt"

	"github.com/eugen-bondarev/backup-tool/config"
)

type Dump struct {
	database   string
	outputFile string
}

func NewDump(database, outputFile string) *Dump {
	return &Dump{
		database:   database,
		outputFile: outputFile,
	}
}

type DBType string

const (
	MySQL DBType = "mysql"
	PG    DBType = "pg"
)

func (d Dump) Create(c config.Config, dbType DBType) error {
	var config GetDumpCommand

	switch dbType {
	case MySQL:
		config = c.MySQL
	case PG:
		config = c.PG
	default:
		return fmt.Errorf("invalid database type: %v", dbType)
	}

	command := config.GetCommand(d.database, d.outputFile)
	return executeCommand(command)
}
