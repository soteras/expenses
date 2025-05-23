package main

import (
	"fmt"
	"os"

	"github.com/soteras/expenses/config"
)

func main() {
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", config.EnvDevelopment)
	}

	config.LoadEnv()

	fmt.Println(config.GetEnv("GO_ENV"))
}
