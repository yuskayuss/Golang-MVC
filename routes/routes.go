package routes

import (
	"go-mvc-project/controllers"
	"go-mvc-project/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	{
		// ✅ Public endpoints
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// ✅ Protected endpoints (butuh JWT token)
		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			// 🛡️ Endpoint yang hanya boleh diakses admin
			adminOnly := protected.Group("/")
			adminOnly.Use(middlewares.RequireRole("admin"))
			{
				adminOnly.GET("/users", controllers.GetUsers)
				adminOnly.GET("/users/:id", controllers.GetUserByID)
				adminOnly.POST("/users", controllers.CreateUser)
				adminOnly.PUT("/users/:id", controllers.UpdateUser)
				adminOnly.DELETE("/users/:id", controllers.DeleteUser)
			}

			// 🧍 Endpoint umum (semua user dengan JWT bisa akses)
			protected.GET("/profile", controllers.Profile)
		}
	}

	return router
}
