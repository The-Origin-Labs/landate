package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvConfig(VARIABLE_NAME string) string {
	err := godotenv.Load(".dev.env") // for production leave it empty
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	ENV_VAR := os.Getenv(VARIABLE_NAME)
	if ENV_VAR == "" {
		fmt.Printf("Unable to load %s from .env file", VARIABLE_NAME)
	}

	return ENV_VAR
}
