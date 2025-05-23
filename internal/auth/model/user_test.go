package model_test

import (
	"testing"

	"github.com/soteras/expenses/internal/auth/model"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	user, _ := model.NewUser("test@gmail.com", "pass1234")

	assert.Equal(t, "test@gmail.com", user.Email)
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte("pass1234"))
	assert.Nil(t, err)
}
