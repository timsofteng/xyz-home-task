package config

import "github.com/caarlos0/env/v11"

type Config struct {
	HTTPServerHost string `env:"HOST,notEmpty"`
	LogLevel       string `env:"LOG_LEVEL"`
	HTTPServerPort string `env:"HTTPServerPort,notEmpty"`
}

func ReadConfig() (Config, error) {
	return env.ParseAs[Config]()
}
