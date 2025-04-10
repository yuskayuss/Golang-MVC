package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret-key")

func GenerateToken(userID uint, email, role string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "email":   email,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 1).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}


func ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    return token, claims, err
}

