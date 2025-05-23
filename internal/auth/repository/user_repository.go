package repository

import (
	"github.com/google/uuid"
	"github.com/soteras/expenses/internal/auth/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id uuid.UUID) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	err := r.db.Create(&user).Error
	return err
}

func (r *userRepository) FindByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}
