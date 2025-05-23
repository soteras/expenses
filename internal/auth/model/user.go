package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/soteras/expenses/internal/shared"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	shared.BaseModel `valid:"-"`
	Email            string `valid:"email,required"`
	Password         string `gorm:"-" valid:"required,minstringlength(6)"`
	PasswordHash     string `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewUser(email string, password string) (*User, error) {
	user := &User{
		Email:    email,
		Password: password,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.PasswordHash = string(hash)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}
