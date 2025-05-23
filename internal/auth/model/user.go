package model

import (
	"github.com/soteras/expenses/internal/shared"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	shared.BaseModel
	Email        string
	PasswordHash string
}

func NewUser(email string, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		Email:        email,
		PasswordHash: string(hash),
	}, nil
}
