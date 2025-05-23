package service

import (
	"github.com/soteras/expenses/internal/auth/model"
	"github.com/soteras/expenses/internal/auth/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s userService) CreateUser(email string, password string) (*model.User, error) {
	user, err := model.NewUser(email, password)

	if err != nil {
		return nil, err
	}

	err = s.repo.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
