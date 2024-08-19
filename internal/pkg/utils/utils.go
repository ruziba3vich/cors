package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ruziba3vich/cors/internal/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type (
	Claims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}

	Utils struct {
		jwtSecret string
	}
)

func New(cfg *config.Config) *Utils {
	return &Utils{
		jwtSecret: cfg.GetJwtSecret(),
	}
}

func (u *Utils) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *Utils) CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (u *Utils) GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.jwtSecret))
}

func (u *Utils) ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return u.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
