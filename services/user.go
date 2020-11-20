package services

import (
	"go-template/connectors"
	"go-template/models"
)

// UserService struct
type UserService struct{}

// Create method
func (s *UserService) Create(name string, email string, password string) (models.User, error) {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	r := connectors.DB.Create(&user)
	return user, r.Error
}

// GetByID method
func (s *UserService) GetByID(ID int) (models.User, error) {
	user := models.User{}
	r := connectors.DB.Where("ID = ?", ID).First(&user)
	return user, r.Error
}
