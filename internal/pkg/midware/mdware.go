package midware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/cors/internal/items/models"
	"github.com/ruziba3vich/cors/internal/items/repo"
	"github.com/ruziba3vich/cors/internal/pkg/utils"
)

type (
	MiddleWare struct {
		logger  *log.Logger
		service repo.CORSRepo
		utils   *utils.Utils
	}
)

func New(logger *log.Logger, service repo.CORSRepo, utils *utils.Utils) *MiddleWare {
	return &MiddleWare{
		logger:  logger,
		service: service,
		utils:   utils,
	}
}

func (m *MiddleWare) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "" {
			c.Next()
			return
		}

		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		trustedOrigins, err := m.service.GetOriginsByUsername(c, username.(string))

		if err != nil || !contains(trustedOrigins, origin) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Origin not allowed"})
			c.Abort()
			return
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", fmt.Sprintf("%s, %s, %s", models.CREATE, models.RETRIEVE, models.REMOVE))
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		c.Next()
	}
}

func (m *MiddleWare) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		m.logger.Println(token)

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		claims, err := m.utils.ValidateJWT(token)
		if err != nil {
			m.logger.Printf("token validation error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func contains(slice []string, item string) bool {
	for i := range slice {
		if slice[i] == item {
			return true
		}
	}
	return false
}
