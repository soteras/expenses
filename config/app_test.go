package config_test

import (
	"testing"

	"github.com/soteras/expenses/config"
	"github.com/stretchr/testify/assert"
)

func TestNewAPP(t *testing.T) {
	app := config.NewApp()

	assert.Equal(t, "test", app.ENV)
}
