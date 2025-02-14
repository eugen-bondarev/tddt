package main

import "fmt"

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

type DBType int

const (
	MySQL DBType = iota
	Postgres
)

func (d Dump) Create(c Config, dbType DBType) error {
	var config getDumpCommand

	switch dbType {
	case MySQL:
		config = c.MySQL
	case Postgres:
		config = c.Postgres
	default:
		return fmt.Errorf("invalid database type: %v", dbType)
	}

	command := config.getCommand(d.database, d.outputFile)
	return executeCommand(command)
}
