package config

type BasicAuthConfig struct {
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
}

type WithBasicAuthConfig struct {
	BasicAuth BasicAuthConfig `envPrefix:"BASIC_AUTH_"`
}
