package main

import (
	"github.com/drewmanole/asteroids-neows/config"
	"github.com/drewmanole/asteroids-neows/internal/app"
)

func main() {
	cfg, _ := config.NewConfig()

	app.Run(cfg)
}
