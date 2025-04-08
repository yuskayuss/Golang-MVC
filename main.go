package main

import (
	"go-mvc-project/config"
	"go-mvc-project/models"
	"go-mvc-project/routes"
)

func main() {
    config.ConnectDB()
    models.MigrateUser()

    r := routes.SetupRoutes()
    r.Run(":9292")
}
