package main

import (
	"fmt"

	"github.com/soteras/expenses/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	fmt.Println("Running command")
}
