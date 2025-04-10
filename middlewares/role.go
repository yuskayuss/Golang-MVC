package middlewares

import (
	"go-mvc-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		u := user.(models.User)
		if u.Role != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden - insufficient role"})
			return
		}

		c.Next()
	}
}
