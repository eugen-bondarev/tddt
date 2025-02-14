package config

import "fmt"

type PGConfig struct {
	Host     string `env:"HOST" envDefault:"127.0.0.1"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD"`
}

type WithPGConfig struct {
	PG PGConfig `envPrefix:"PG_"`
}

func (c PGConfig) GetCommand(database, outputFile string) string {
	return fmt.Sprintf("PGPASSWORD=%s pg_dump -U %s -h %s -d %s >> %s", c.Password, c.User, c.Host, database, outputFile)
}
