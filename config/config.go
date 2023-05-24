package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvConfig(VARIABLE_NAME string) string {

	// workingDir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("Failed to get the current working directory:", err)
	// }

	// fmt.Println(workingDir)
	// envFilePath := filepath.Join(workingDir, ".env")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	ENV_VAR := os.Getenv(VARIABLE_NAME)
	if ENV_VAR == "" {
		log.Fatalf("Unable to load %s from .env file", VARIABLE_NAME)
	}

	return ENV_VAR
}
