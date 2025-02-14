package config

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`
	WithMySQLConfig
	WithPGConfig
	WithGCPConfig
}
