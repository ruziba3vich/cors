package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/cors/internal/items/models"
	"github.com/ruziba3vich/cors/internal/items/repo"
	"github.com/ruziba3vich/cors/internal/pkg/utils"
)

type (
	Storage struct {
		rdb         *redis.Client
		logger      *log.Logger
		utils       *utils.Utils
		corsStorage repo.CORSRepo
	}
)

func New(rdb *redis.Client, logger *log.Logger, utils *utils.Utils) *Storage {
	return &Storage{
		logger: logger,
		rdb:    rdb,
		utils:  utils,
	}
}

func (s *Storage) Register(ctx context.Context, req *models.User) (*models.User, error) {
	_, err := s.rdb.Get(ctx, req.Username).Result()
	if err != nil {
		if err == redis.Nil {
			hashedPwd, err := s.utils.HashPassword(req.Password)
			if err != nil {
				s.logger.Println(err)
				return nil, err
			}
			req.Password = hashedPwd
			if err := s.rdb.Set(ctx, req.Username, req.Password, time.Hour*24).Err(); err != nil {
				s.logger.Println(err)
				return nil, err
			}
			return req, nil
		}
		return nil, err
	}
	s.logger.Printf("user with username %s already exists\n", req.Username)
	return nil, fmt.Errorf("user with username %s already exists", req.Username)
}

func (s *Storage) LoginUser(ctx context.Context, req *models.User) (*models.LoginUserResponse, error) {
	pwd, err := s.rdb.Get(ctx, req.Username).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("user with username %s does not exist", req.Username)
		}
		return nil, err
	}
	if ok := s.utils.CheckPasswordHash(req.Password, pwd); ok {
		newToken, err := s.utils.GenerateJWT(req.Username)
		if err != nil {
			s.logger.Println(err)
			return nil, err
		}

		s.corsStorage.CreateOrigin(ctx, &models.CreateOriginRequest{
			Username: req.Username,
			Origin:   "http://localhost:7777/origins",
		})

		return &models.LoginUserResponse{
			User:  req,
			Token: newToken,
		}, nil
	}
	return nil, fmt.Errorf("password missmatch")
}
