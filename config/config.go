package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	envPath := os.Getenv("ENV_PATH")
	if envPath == "" {
		envPath = "config/.env.dev"
	}
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("[config] Warning: no env file found at %s", envPath)
	}
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
