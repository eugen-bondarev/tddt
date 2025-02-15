package config

type Config struct {
	Port int    `env:"PORT" envDefault:"8080"`
	Mode string `env:"MODE" envDefault:"release"`
	WithMySQLConfig
	WithPGConfig
	WithGCPConfig
	WithBasicAuthConfig
}
