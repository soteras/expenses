package config_test

import (
	"testing"

	"github.com/soteras/expenses/config"
	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	config.LoadEnv()

	assert.Equal(t, "test", config.GetEnv("GO_ENV"))
}

func TestGetEnv_WithDefaultValue(t *testing.T) {
	config.LoadEnv()

	assert.Equal(t, "default", config.GetEnv("ENV", "default"))
}
