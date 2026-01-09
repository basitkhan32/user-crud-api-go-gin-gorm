package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to laod env", err)
	}
	return err
}
