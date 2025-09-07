package logger

import (
	"fmt"
	"os"

	"github.com/geniusroom/catalog/internal/config"
	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func New(cfg *config.Config) *Logger {
	stream := fmt.Sprintf("app=%s,env=%s", cfg.App.Name, cfg.App.Env)

	logger := zerolog.New(os.Stdout).With().
		Str("_stream", stream).
		Timestamp().
		Logger().
		Level(zerolog.Level(cfg.Log.Level))

	return &Logger{logger}
}
