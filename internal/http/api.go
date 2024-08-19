package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/cors/internal/http/handlers"
	"github.com/ruziba3vich/cors/internal/items/service"
	storage "github.com/ruziba3vich/cors/internal/items/strage"
	"github.com/ruziba3vich/cors/internal/pkg/config"
	"github.com/ruziba3vich/cors/internal/pkg/midware"
	"github.com/ruziba3vich/cors/internal/pkg/utils"
)

func Run(logger *log.Logger) error {
	config, err := config.New()
	if err != nil {
		logger.Println(err)
		return err
	}
	rdb := storage.GetRedisConn(config)
	utilss := utils.New(config)
	handler := handlers.New(logger, service.New(storage.New(rdb, logger, utilss), logger))

	router := gin.Default()
	router.POST("/register", handler.Register)
	router.POST("login", handler.Login)

	corsService := service.NewCORSImpleService(storage.NewCorsStorage(rdb, logger), logger)
	moddleware := midware.New(logger, corsService, utilss)
	corsHandler := handlers.NewCORSHandler(logger, corsService)

	worker := router.Group("/origins")
	worker.Use(moddleware.AuthMiddleware())
	worker.Use(moddleware.CORSMiddleware())

	worker.POST("/", corsHandler.AddOriginToUser)
	worker.GET("/", corsHandler.GetOriginsByUsername)
	worker.DELETE("/", corsHandler.DeleteOriginByUsername)

	return router.Run(":" + config.GetAppHost())
}
