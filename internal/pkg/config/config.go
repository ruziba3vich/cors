package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	jwtSecret           string
	redisHost           string
	redisPort           string
	appHost             string
	rateLimitingSeconds int
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	strrls := os.Getenv("RATE_LIMITING_SECONDS")
	val, err := strconv.Atoi(strrls)
	if err != nil {
		return nil, err
	}
	return &Config{
		jwtSecret:           os.Getenv("JWT_SECRET"),
		redisHost:           os.Getenv("REDIS_HOST"),
		redisPort:           os.Getenv("REDIS_PORT"),
		appHost:             os.Getenv("APP_HOST"),
		rateLimitingSeconds: val,
	}, nil
}

func (c *Config) GetJwtSecret() string {
	return c.jwtSecret
}

func (c *Config) GetRedisHost() string {
	return c.redisHost
}

func (c *Config) GetRedisPort() string {
	return c.redisPort
}

func (c *Config) GetAppHost() string {
	return c.appHost
}

func (c *Config) GetRLS() int {
	return c.rateLimitingSeconds
}
