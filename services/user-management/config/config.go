package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIBaseURL string
	APIKey     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	LogLevel   string
	ServerPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("warning: No .env file found, relying on system environment variables")
	}

	return &Config{
		APIBaseURL: getEnv("API_BASE_URL", ""),
		APIKey:     getEnv("API_KEY", ""),
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", ""),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		LogLevel:   getEnv("LOG_LEVEL", ""),
		ServerPort: getEnv("SERVER_PORT", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

// loadEnv loads environment variables from a .env file
func loadEnv() error {
	if _, err := os.Stat(".env"); err == nil {
		if err := os.Setenv("DB_HOST", "localhost"); err != nil {
			return err
		}
		if err := os.Setenv("DB_USER", "your_db_user"); err != nil {
			return err
		}
		if err := os.Setenv("DB_PASSWORD", "your_db_password"); err != nil {
			return err
		}
		if err := os.Setenv("DB_NAME", "your_db_name"); err != nil {
			return err
		}
		if err := os.Setenv("DB_PORT", "5432"); err != nil {
			return err
		}
		if err := os.Setenv("PORT", "8088"); err != nil {
			return err
		}
	}
	return nil
}
