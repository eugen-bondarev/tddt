package main

import "fmt"

type MySQLConfig struct {
	Host     string `env:"HOST" envDefault:"127.0.0.1"`
	User     string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD"`
}

type PostgresConfig struct {
	Host     string `env:"HOST" envDefault:"127.0.0.1"`
	User     string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD"`
}

type Config struct {
	MySQL    MySQLConfig    `envPrefix:"MYSQL_"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
}

func (c MySQLConfig) getCommand(database, outputFile string) string {
	return fmt.Sprintf("mysqldump -u %s -h %s -p%s %s > %s", c.User, c.Host, c.Password, database, outputFile)
}

func (c PostgresConfig) getCommand(database, outputFile string) string {
	return fmt.Sprintf("pg_dump -U %s -h %s -p %s -d %s -F c -b -v -f %s", c.User, c.Host, c.Password, database, outputFile)
}
