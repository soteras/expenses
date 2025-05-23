package service

import (
	"github.com/soteras/expenses/internal/auth/errors"
	"github.com/soteras/expenses/internal/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repo repository.UserRepository
}

func NewLoginService(repo repository.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s LoginService) Login(email string, password string) (string, error) {
	authError := errors.NewAuthError("Email or password invalid")
	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return "", authError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte("pass1234")); err != nil {
		return "", authError
	}

	return "", nil
}
