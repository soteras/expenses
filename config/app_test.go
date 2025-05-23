package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPP(t *testing.T) {
	app := NewApp()

	assert.Equal(t, "test", app.ENV)
}
