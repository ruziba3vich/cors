package storage

import (
	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/cors/internal/pkg/config"
)

func GetRedisConn(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.GetRedisHost() + ":" + cfg.GetRedisPort(),
		Password: "",
		DB:       0,
	})
}
