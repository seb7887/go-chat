package services

import (
	"fmt"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/storage"
)

type UserService interface {
	AddUser(username string, password string) (*models.User, error)
	LoginUser(username string, password string) (*models.Login, error)
}

type userService struct {
	repository storage.UserRespository
}

func NewUserService(repository storage.UserRespository) UserService {
	return &userService{
		repository,
	}
}

func (s *userService) AddUser(username string, password string) (*models.User, error) {
	hash, err := helpers.HashAndSalt(password)
	if err != nil {
		return nil, err
	}

	newUser, err := s.repository.Create(username, string(hash))
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) LoginUser(username string, password string) (*models.Login, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("User not found")
	}

	if ok := helpers.VerifyPassword(password, user.Password); !ok {
		return nil, fmt.Errorf("Invalid credentials")
	}

	return helpers.GenerateJwt(user.ID, user.Username)
}
