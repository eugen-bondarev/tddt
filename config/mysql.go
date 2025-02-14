package config

import "fmt"

type MySQLConfig struct {
	Host     string `env:"HOST" envDefault:"127.0.0.1"`
	User     string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD"`
}

type WithMySQLConfig struct {
	MySQL MySQLConfig `envPrefix:"MYSQL_"`
}

func (c MySQLConfig) GetCommand(database, outputFile string) string {
	return fmt.Sprintf("mysqldump -u %s -h %s -p%s %s > %s", c.User, c.Host, c.Password, database, outputFile)
}
