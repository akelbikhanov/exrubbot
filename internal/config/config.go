// Package config отвечает за загрузку переменных окружения.
package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/akelbikhanov/exrubbot/pkg/logger"
	"github.com/joho/godotenv"
)

const (
	envBotToken = "EXRUBBOT_TOKEN"
)

// Config - структура для хранения конфигурации приложения
type Config struct {
	BotToken    string        // Токен Telegram бота
	BotTimeout  time.Duration // Таймаут long-polling запросов (getUpdates)
	BotRetry    time.Duration // Таймаут при ErrorTooManyRequests (SendMessage, etc.)
	FeedTimeout time.Duration // Таймаут http-запросов к поставщикам данных, биржам
}

var (
	cfg     Config
	initErr error
	once    sync.Once
)

// Get загружает переменные окружения из файла .env (если он есть)
// а затем пытается их прочесть из окружения
func Get(l *logger.Logger) (*Config, error) {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			if os.IsNotExist(err) {
				l.Info(text.InfoEnvFileNotFound)
			} else {
				initErr = fmt.Errorf("%s: %w", text.ErrEnvFileLoad, err)
				return
			}
		}

		cfg.BotToken = os.Getenv(envBotToken)
		if cfg.BotToken == "" {
			initErr = fmt.Errorf("%s '%s'", text.ErrEnvMissingVar, envBotToken)
			return
		}

		cfg.BotTimeout = 60 * time.Second
		cfg.BotRetry = 5 * time.Second
		cfg.FeedTimeout = 5 * time.Second
	})

	if initErr != nil {
		return nil, initErr
	}

	return &cfg, nil
}
