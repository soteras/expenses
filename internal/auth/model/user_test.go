package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	user, _ := NewUser("test@gmail.com", "pass1234")

	assert.Equal(t, "test@gmail.com", user.Email)
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte("pass1234"))
	assert.Nil(t, err)
}

func TestNewUser_EmailInvalid(t *testing.T) {
	_, err := NewUser("wrong", "pass1234")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Email: wrong does not validate as email")
}

func TestNewUser_PasswordIsBlank(t *testing.T) {
	_, err := NewUser("email@gmail.com", "")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Password: non zero value required")
}

func TestNewUser_PasswordIsLessThan6(t *testing.T) {
	_, err := NewUser("email@gmail.com", "pass")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Password: pass does not validate as minstringlength(6)")
}
