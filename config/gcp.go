package config

type GCPConfig struct {
	Credentials string `env:"CREDENTIALS"`
}

type WithGCPConfig struct {
	GCPConfig GCPConfig `envPrefix:"GOOGLE_CLOUD_"`
}

func (c *GCPConfig) IsDefined() bool {
	return len(c.Credentials) > 0
}
