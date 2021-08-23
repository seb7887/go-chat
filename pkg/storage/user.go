package storage

import (
	"github.com/challenge/pkg/models"
)

type UserRespository interface {
	Create(username string, password string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRespository {
	return &userRepository{}
}

func (r *userRepository) Create(username string, password string) (*models.User, error) {
	user := models.User{
		Username: username,
		Password: password,
	}

	db := GetInstance()
	res := db.Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
