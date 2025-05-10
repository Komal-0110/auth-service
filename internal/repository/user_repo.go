package repository

import (
	"auth-service/internal/db"
	"auth-service/internal/models"
)

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := db.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}
