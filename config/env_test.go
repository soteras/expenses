package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	assert.Equal(t, "test", GetEnv("GO_ENV"))
}

func TestGetEnv_WithDefaultValue(t *testing.T) {
	assert.Equal(t, "default", GetEnv("ENV", "default"))
}
