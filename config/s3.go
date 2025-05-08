package config

type S3Config struct {
	AccessKeyID     string `env:"ACCESS_KEY_ID"`
	SecretAccessKey string `env:"SECRET_ACCESS_KEY"`
	Endpoint        string `env:"ENDPOINT"`
}

type WithS3Config struct {
	S3Config S3Config `envPrefix:"S3_"`
}

func (c *S3Config) IsDefined() bool {
	return len(c.AccessKeyID) > 0 && len(c.SecretAccessKey) > 0
}
