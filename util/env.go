package util

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func LoadEnv[T any]() T {
	godotenv.Load()
	t, _ := env.ParseAs[T]()
	return t
}
