package config_test

import (
	"testing"

	"github.com/soteras/expenses/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewDb(t *testing.T) {
	db := config.NewDb()

	assert.NotNil(t, db)
	assert.IsType(t, &config.Database{}, db)
}

func TestConnec(t *testing.T) {
	db := config.NewDb()

	conn, err := db.Connect()

	assert.Nil(t, err)
	assert.IsType(t, &gorm.DB{}, conn)
}
