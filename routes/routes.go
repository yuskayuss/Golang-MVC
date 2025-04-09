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

		// ✅ Protected endpoints (pakai JWT)
		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/users", controllers.GetUsers)
			protected.GET("/users/:id", controllers.GetUserByID)
			protected.POST("/users", controllers.CreateUser)
			protected.PUT("/users/:id", controllers.UpdateUser)
			protected.DELETE("/users/:id", controllers.DeleteUser)

			protected.GET("/profile", controllers.Profile)
		}
	}

	return router
}

