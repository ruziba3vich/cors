package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/cors/internal/http/handlers"
	"github.com/ruziba3vich/cors/internal/items/service"
	storage "github.com/ruziba3vich/cors/internal/items/strage"
	"github.com/ruziba3vich/cors/internal/pkg/config"
	"github.com/ruziba3vich/cors/internal/pkg/utils"
)

func Run(logger *log.Logger) error {
	config, err := config.New()
	if err != nil {
		logger.Println(err)
		return err
	}
	rdb := storage.GetRedisConn(config)
	handler := handlers.New(logger, service.New(storage.New(rdb, logger, utils.New(config)), logger))

	router := gin.Default()
	router.POST("/register", handler.Register)
	router.POST("login", handler.Login)
	return router.Run(":" + config.GetAppHost())
}
