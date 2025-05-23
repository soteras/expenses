package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewDb(t *testing.T) {
	db := NewDb()

	assert.NotNil(t, db)
	assert.IsType(t, &Database{}, db)
}

func TestConnec(t *testing.T) {
	db := NewDb()

	err := db.Connect()

	assert.Nil(t, err)
	assert.IsType(t, &gorm.DB{}, db.Conn)
}
