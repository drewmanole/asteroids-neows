package app

import (
	"github.com/gin-gonic/gin"
	"github.com/drewmanole/asteroids-neows/config"
	"github.com/drewmanole/asteroids-neows/internal/transport"
	"github.com/drewmanole/asteroids-neows/pkg/httpserver"
)

func Run(cfg *config.Config) {

	// HTTP Server
	handler := gin.New()
	transport.NewRouter(handler, nil)
	httpServer := httpserver.New(handler, cfg.HTTP.Port)

	// TODO: Обработать ошибку
	httpServer.Run()
}
