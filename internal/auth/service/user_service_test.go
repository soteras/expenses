package service

import (
	"testing"

	"github.com/soteras/expenses/internal/auth/mocks"
	"github.com/soteras/expenses/internal/auth/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUserService(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	email := "test@gmail.com"
	password := "pass1234"
	mockRepo.On("Create", mock.MatchedBy(func(u *model.User) bool {
		if u.Email != email {
			return false
		}

		err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
		return err == nil
	})).Return(nil)
	service := NewUserService(mockRepo)

	_, err := service.CreateUser("test@gmail.com", "pass1234")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
