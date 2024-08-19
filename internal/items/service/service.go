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

	CORSServiceImple struct {
		storage *storage.CorsStorage
		logger  *log.Logger
	}
)

func New(storage *storage.Storage, logger *log.Logger) repo.UserServiceRepo {
	return &UserServiceImple{
		storage: storage,
		logger:  logger,
	}
}

func NewCORSImpleService(storage *storage.CorsStorage, logger *log.Logger) repo.CORSRepo {
	return &CORSServiceImple{
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

func (c *CORSServiceImple) GetOriginsByUsername(ctx context.Context, username string) ([]string, error) {
	c.logger.Println("-- RECEIVED A REQUEST INTO `GetOriginsByUsername` SERVICE")
	return c.storage.GetOriginsByUsername(ctx, username)
}

func (c *CORSServiceImple) CreateOrigin(ctx context.Context, req *models.CreateOriginRequest) (string, error) {
	c.logger.Println("-- RECEIVED A REQUEST INTO `CreateOrigin` SERVICE")
	return c.storage.AddOriginForUser(ctx, req)
}

func (c *CORSServiceImple) DeleteOrigin(ctx context.Context, username string, origin string) (string, error) {
	c.logger.Println("-- RECEIVED A REQUEST INTO `DeleteOrigin` SERVICE")
	return c.storage.DeleteOriginByUsername(ctx, username, origin)
}

/*
	GetOriginsByUsername(context.Context, string) ([]string, error)
	CreateOrigin(context.Context, *models.CreateOriginRequest) (string, error)
	DeleteOrigin(context.Context, string, string) (string, error)
*/
