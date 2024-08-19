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
)
