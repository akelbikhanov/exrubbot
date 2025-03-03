package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config - структура для хранения конфигурации
type Config struct {
	TelegramToken string
}

// LoadConfig загружает .env, но позволяет переопределять переменные через ОС
func LoadConfig() *Config {
	_ = godotenv.Load() // Загружаем .env (если есть)

	config := &Config{
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}

	if config.TelegramToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не задан в .env или переменных окружения")
	}

	return config
}
