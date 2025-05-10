package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var ErrNotFound = errors.New("user not found")
var ErrEmailRegistered = errors.New("email already registered")

func RegisterUser(name, email, password string) (*models.User, error) {
	normlizedEmail := normalizeEmail(email)

	if err := validateInput(name, normlizedEmail, password); err != nil {
		return nil, err
	}

	if exists, _ := isEmailTaken(normlizedEmail); exists {
		return nil, ErrEmailRegistered
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:      name,
		Email:     normlizedEmail,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repository.CreateUser(user); err != nil {
		return nil, err
	}

	user.Password = "" // hide password in response
	return user, nil
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func validateInput(name, email, password string) error {
	if name == "" || email == "" || password == "" {
		return errors.New("all fields are required")
	}

	return nil
}

func isEmailTaken(email string) (bool, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return false, ErrNotFound
	}

	return user.ID != 0, nil
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing password")
	}

	return string(hashed), nil
}
