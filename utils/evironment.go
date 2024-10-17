package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariables() {
	env := os.Getenv("ENV")
	if env != "local" {
		err := godotenv.Load()
		if err != nil {
			LogError("Error loading .env file")
			panic(err)
		}
		LogSuccess("Using " + env + " environment variables")
	} else {
		err := godotenv.Load(".env.local")
		if err != nil {
			LogError("Error loading .env.local file")
			panic(err)
		}
		LogWarning("Using local environment variables")
	}
}
