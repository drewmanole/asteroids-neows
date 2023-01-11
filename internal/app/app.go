package app

import (
	"github.com/gin-gonic/gin"
	"github.com/drewmanole/asteroids-neows/config"
	"github.com/drewmanole/asteroids-neows/internal/repo/neows"
	"github.com/drewmanole/asteroids-neows/internal/repo/storage"
	"github.com/drewmanole/asteroids-neows/internal/transport"
	"github.com/drewmanole/asteroids-neows/internal/usecase"
	"github.com/drewmanole/asteroids-neows/pkg/httpserver"
	"github.com/drewmanole/asteroids-neows/pkg/postgres"
	"log"
)

func Run(cfg *config.Config) {
	// DB
	db, err := postgres.New(cfg.DB.GetURL())
	if err != nil {
		log.Fatal(err)
	}

	// Storage
	stg, err := storage.New(db)
	if err != nil {
		log.Fatal(err)
	}

	// NeoWs API
	api := neows.New(cfg.NeoWs.URL, cfg.NeoWs.ApiKey)

	// Use case
	uc := usecase.New(stg, api)

	// HTTP Server
	handler := gin.New()
	transport.NewRouter(handler, uc)
	httpServer := httpserver.New(handler, cfg.HTTP.Port)

	if err = httpServer.Run(); err != nil {
		log.Fatal(err)
	}
}
