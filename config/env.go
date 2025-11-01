package config

import (
	"github.com/caarlos0/env/v11"
)

type EnvConfig struct {
	ApiKey string `env:"API_KEY,required"`
	Port   string `env:"PORT,required"`
}

func LoadEnv() (EnvConfig, error) {
	return env.ParseAs[EnvConfig]()
}
