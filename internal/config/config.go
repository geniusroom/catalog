package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Server Server `envPrefix:"SERVER_"`
}

type Server struct {
	Port         string        `env:"PORT,required" envDefault:"8080"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT,required" envDefault:"5s"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT,required" envDefault:"5s"`
}

func GetDefault() *Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to read config: %s", err.Error())
	}
	return &cfg
}
