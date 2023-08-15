package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvConfig(VARIABLE_NAME string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	ENV_VAR := os.Getenv(VARIABLE_NAME)
	return ENV_VAR
}
