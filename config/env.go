package config

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
	EnvTest        = "test"
)

func LoadEnv() {
	if os.Getenv("GO_ENV") == EnvTest {
		projectDirName := "expenses"
		projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
		currentWorkDirectory, _ := os.Getwd()
		rootPath := projectName.Find([]byte(currentWorkDirectory))

		err := godotenv.Load(string(rootPath) + `/.env`)

		if err != nil {
			panic(err)
		}
	}

	if os.Getenv("GO_ENV") == EnvDevelopment {
		err := godotenv.Load()

		if err != nil {
			panic(err)
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
