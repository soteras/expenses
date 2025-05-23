package model

import (
	"testing"

	authmodel "github.com/soteras/expenses/internal/auth/model"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	account, err := NewAccount("Checking account Itau", &authmodel.User{})

	assert.Nil(t, err)
	assert.Equal(t, "Checking account Itau", account.Name)
}

func TestValidate_NameIsBlank(t *testing.T) {
	_, err := NewAccount("", &authmodel.User{})

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Name: non zero value required")
}
