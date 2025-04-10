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


// controllers/auth_controller.go



func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"` // admin atau user
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role, // ← role langsung disimpan
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register berhasil"})
}


func Login(c *gin.Context) {
    var input models.User
    var user models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Cari user dari DB
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
        return
    }

    // Cek password dengan bcrypt
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
        return
    }

    // ✅ Buat token pakai ID, email, dan role
    token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal generate token"})
        return
    }

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
