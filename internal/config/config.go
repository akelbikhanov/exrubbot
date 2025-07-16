// Package config отвечает за загрузку конфигурации приложения из флагов и переменных окружения.
package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/joho/godotenv"
)

// Config - структура для хранения конфигурации приложения
type Config struct {
	BotToken    string        // Токен Telegram бота
	SubsPath    string        // Путь к файлу со списком подписчиков
	BotTimeout  time.Duration // Таймаут long-polling запросов (getUpdates)
	BotRetry    time.Duration // Таймаут при ErrorTooManyRequests (SendMessage, etc.)
	FeedTimeout time.Duration // Таймаут http-запросов к поставщикам данных, биржам
}

// LoadConfig загружает конфигурацию из .env файла и переменных окружения.
func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return Config{}, fmt.Errorf(text.ErrConfigEnvLoad, text.LogErrEnvFile, err)
	}

	SubsPath := strings.TrimSpace(os.Getenv(text.EnvStoragePath))
	botToken := strings.TrimSpace(os.Getenv(text.EnvBotToken))
	if botToken == "" {
		return Config{}, fmt.Errorf(text.ErrConfigTokenEmpty, text.LogErrEnvMissing, text.EnvBotToken)
	}

	return Config{
		BotToken:    botToken,
		SubsPath:    SubsPath,
		BotTimeout:  60 * time.Second,
		BotRetry:    5 * time.Second,
		FeedTimeout: 5 * time.Second,
	}, nil
}
