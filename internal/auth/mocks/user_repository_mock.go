package mocks

import (
	"github.com/google/uuid"
	"github.com/soteras/expenses/internal/auth/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindByID(id uuid.UUID) (*model.User, error) {
	args := m.Called(id)
	user, _ := args.Get(0).(*model.User)
	return user, args.Error(1)
}
