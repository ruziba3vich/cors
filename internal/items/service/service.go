package service

import (
	"context"
	"log"

	"github.com/ruziba3vich/cors/internal/items/models"
	"github.com/ruziba3vich/cors/internal/items/repo"
	storage "github.com/ruziba3vich/cors/internal/items/strage"
)

type (
	UserServiceImple struct {
		storage *storage.Storage
		logger  *log.Logger
	}
)

func New(storage *storage.Storage, logger *log.Logger) repo.UserServiceRepo {
	return &UserServiceImple{
		storage: storage,
		logger:  logger,
	}
}

func (u *UserServiceImple) RegisterUser(ctx context.Context, req *models.User) (*models.User, error) {
	u.logger.Println("-- RECEIVED A REQUEST INTO `Register` SERVICE")
	return u.storage.Register(ctx, req)
}

func (u *UserServiceImple) LoginUser(ctx context.Context, req *models.User) (*models.LoginUserResponse, error) {
	u.logger.Println("-- RECEIVED A REQUEST INTO `LoginUser` SERVICE")
	return u.storage.LoginUser(ctx, req)
}
