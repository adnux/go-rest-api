package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariables() {
	env := os.Getenv("ENV")
	if env != "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		LogSuccess("Using " + env + " environment variables")
	} else {
		err := godotenv.Load(".env.local")
		if err != nil {
			LogError("Error loading .env file")
		}
		LogWarning("Using local environment variables")
	}
}
