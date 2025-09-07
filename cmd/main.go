package main

import (
	"fmt"

	"github.com/geniusroom/catalog/internal/config"
)

func main() {
	cfg := config.GetDefault()
	fmt.Printf("%+v", cfg)
}
