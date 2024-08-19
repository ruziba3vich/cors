package storage

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/cors/internal/items/models"
)

type (
	CorsStorage struct {
		rdb    *redis.Client
		logger *log.Logger
	}
)

func NewCorsStorage(rdb *redis.Client, logger *log.Logger) *CorsStorage {
	return &CorsStorage{
		rdb:    rdb,
		logger: logger,
	}
}

func (c CorsStorage) AddOriginForUser(ctx context.Context, req *models.CreateOriginRequest) (string, error) {
	if err := c.rdb.SAdd(ctx, "user:"+req.Username+":origins", req.Origin).Err(); err != nil {
		c.logger.Println(err)
		return "", err
	}
	return "OK", nil
}

func (c *CorsStorage) GetOriginsByUsername(ctx context.Context, username string) ([]string, error) {
	result, err := c.rdb.SMembers(ctx, username).Result()
	if err != nil {
		c.logger.Println(err)
		return nil, err
	}
	return result, nil
}

func (c *CorsStorage) DeleteOriginByUsername(ctx context.Context, username, origin string) (string, error) {
	if err := c.rdb.SRem(ctx, "user:"+username+":origins", origin).Err(); err != nil {
		c.logger.Println(err)
		return "", err
	}
	return "OK", nil
}
