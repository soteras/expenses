package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/soteras/expenses/internal/auth/model"
	"github.com/soteras/expenses/internal/shared"
)

type Account struct {
	shared.BaseModel `valid:"-"`
	Name             string      `valid:"required"`
	User             *model.User `valid:"-"`
	UserId           uuid.UUID   `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewAccount(name string, user *model.User) (*Account, error) {
	account := &Account{
		Name: name,
		User: user,
	}

	if err := account.Validate(); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) Validate() error {
	_, err := govalidator.ValidateStruct(a)
	return err
}
