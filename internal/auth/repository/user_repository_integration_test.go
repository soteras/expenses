package repository

import (
	"testing"

	"github.com/soteras/expenses/config"
	"github.com/soteras/expenses/internal/auth/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupDb() *gorm.DB {
	db := config.NewDb()
	conn, _ := db.Connect()
	return conn
}

func TestCreate_Create(t *testing.T) {
	conn := setupDb()
	conn.Exec("DELETE FROM users")
	repo := NewUserRepository(conn)
	user, _ := model.NewUser("test@gmail.com", "pass1234")

	repo.Create(user)
	target, _ := repo.FindByID(user.ID)

	assert.Equal(t, user.ID.String(), target.ID.String())
	assert.Equal(t, "test@gmail.com", target.Email)
}
