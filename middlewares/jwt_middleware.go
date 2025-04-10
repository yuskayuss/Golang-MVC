package middlewares

import (
	"go-mvc-project/config"
	"go-mvc-project/models"
	"go-mvc-project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, claims, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		// Ambil user_id dari claims
		userID := uint(claims["user_id"].(float64))

		// Cari user dari database
		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
			c.Abort()
			return
		}

		// Set user ke context supaya bisa dibaca di RequireRole
		c.Set("user", user)
		c.Next()
	}
}
