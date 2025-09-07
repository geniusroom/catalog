package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	App    App    `envPrefix:"APP_"`
	Log    Log    `envPrefix:"LOG_"`
	Server Server `envPrefix:"SERVER_"`
}

type App struct {
	Name string `env:"NAME,required" envDefault:"geniusroom"`
	Env  string `env:"ENV,required" envDefault:"dev"`
}

type Log struct {
	Level int `env:"LEVEL,required" envDefault:"0"`
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
