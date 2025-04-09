package models

import "go-mvc-project/config"

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}


func MigrateUser() {
    config.DB.AutoMigrate(&User{})
}
