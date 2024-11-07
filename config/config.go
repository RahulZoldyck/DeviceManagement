package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

var AppConfig *Config

// LoadConfig loads configuration from .env file or environment variables
func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = &Config{
		Port:   getEnv("PORT", "8080"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "3306"), // Changed to MySQL default port
		DBUser: getEnv("DB_USER", ""),
		DBPass: getEnv("DB_PASSWORD", ""),
		DBName: getEnv("DB_NAME", ""),
	}

	return AppConfig
}

// getEnv retrieves an environment variable or returns a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
