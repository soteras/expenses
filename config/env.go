package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
	EnvTest        = "test"
)

func LoadEnv() {
	fmt.Println(os.Getenv("GO_ENV"))
	if os.Getenv("GO_ENV") == EnvTest {
		godotenv.Load(".env.test")
	}

	if os.Getenv("GO_ENV") == EnvDevelopment {
		godotenv.Load()
	}
}

func GetEnv(key string, fallback ...string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if len(fallback) > 0 {
		return fallback[0]
	}

	return ""
}
