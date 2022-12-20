package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//LoadEnvVars initialized dotenv and loads env vars from .env file
func LoadEnvVars() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
