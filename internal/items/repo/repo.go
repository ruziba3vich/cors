package repo

import (
	"context"

	"github.com/ruziba3vich/cors/internal/items/models"
)

type (
	UserServiceRepo interface {
		RegisterUser(context.Context, *models.User) (*models.User, error)
		LoginUser(context.Context, *models.User) (*models.LoginUserResponse, error)
	}

	CORSRepo interface {
		GetOriginsByUsername(context.Context, string) ([]string, error)
		CreateOrigin(context.Context, *models.CreateOriginRequest) (string, error)
		DeleteOrigin(context.Context, string, string) (string, error)
	}
)
