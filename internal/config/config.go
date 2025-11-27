package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	TelegramAPIToken string
	Debug            bool
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(".env no encontrado; usando variables de entorno del sistema")
	}

	token := os.Getenv("TELEGRAM_APITOKEN")
	if token == "" {
		return nil, os.ErrNotExist
	}

	return &Config{
		TelegramAPIToken: token,
		Debug:            true, // Default to true as per original code, can be made configurable
	}, nil
}
