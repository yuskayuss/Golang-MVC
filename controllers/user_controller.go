package controllers

import (
	"go-mvc-project/models"
	"net/http"

	"go-mvc-project/config"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    var users []models.User
    config.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Create(&user)
    c.JSON(http.StatusOK, user)
}
