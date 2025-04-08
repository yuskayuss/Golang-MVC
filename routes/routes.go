package routes

import (
	"go-mvc-project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    router := gin.Default()

    api := router.Group("/api")
    {
        api.GET("/users", controllers.GetUsers)
        api.POST("/users", controllers.CreateUser)
    }

    return router
}
