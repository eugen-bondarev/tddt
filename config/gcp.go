package config

type GCPConfig struct {
	Credentials string `env:"CREDENTIALS"`
}

type WithGCPConfig struct {
	GCPConfig GCPConfig `envPrefix:"GOOGLE_CLOUD_"`
}
