package controllers

import (
	"go-mvc-project/models"
	"net/http"

	"go-mvc-project/config"

	"github.com/gin-gonic/gin"

	"go-mvc-project/utils"

	"golang.org/x/crypto/bcrypt"
)
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := config.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
        return
    }

    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.Name = input.Name
    user.Email = input.Email

    config.DB.Save(&user)
    c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := config.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
        return
    }

    config.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}


func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := config.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, user)
}


func Register(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    input.Password = string(hash)

    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal registrasi"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil"})
}

func Login(c *gin.Context) {
    var input models.User
    var user models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Where("email = ?", input.Email).First(&user)
    if user.ID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
        return
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
        return
    }

    token, _ := utils.GenerateToken(user.Email)
    c.JSON(http.StatusOK, gin.H{"token": token})
}

func Profile(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Selamat datang di endpoint terproteksi!"})
}


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
