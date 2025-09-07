package main

import (
	"github.com/geniusroom/catalog/internal/config"
	"github.com/geniusroom/catalog/internal/logger"
)

func main() {
	cfg := config.GetDefault()
	log := logger.New(cfg)

	_ = log
}
