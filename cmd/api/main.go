package main

import (
	"fmt"
	"os"

	"github.com/soteras/expenses/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	fmt.Println("API running..")
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", config.EnvDevelopment)
	}
}
