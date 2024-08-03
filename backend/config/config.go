package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %s", err)
	}

	envPath := filepath.Join(pwd, "../.env")
	log.Printf("Attempting to load .env file from: %s", envPath)

	// Use "../.env" because main.go inside /cmd
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
