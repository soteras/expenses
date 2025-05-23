package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
	EnvTest        = "test"
)

func LoadEnv() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	root := filepath.Dir(wd)
	env := os.Getenv("GO_ENV")

	if env == "" {
		envPath := filepath.Join(root, ".env.test")
		env = EnvTest

		if err := godotenv.Load(envPath); err != nil {
			log.Fatalf("%s", fmt.Sprintf("Error to load .env.test: %v", err))
		}
	}

	if env == EnvDevelopment {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("%s", fmt.Sprintf("Error to load .env: %v", err))
		}
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
