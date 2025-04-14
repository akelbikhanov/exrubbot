// Package config отвечает за загрузку переменных окружения.
package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/logger"
	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/joho/godotenv"
)

// Config - структура для хранения конфигурации приложения
type Config struct {
	BotToken string        // Токен Telegram бота
	Timeout  time.Duration // Таймаут запросов по умолчанию
}

var (
	cfg     Config
	initErr error
	once    sync.Once
)

// Get загружает переменные окружения из файла .env (если он есть)
// а затем пытается их прочесть из окружения
func Get() (*Config, error) {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			if os.IsNotExist(err) {
				logger.Info(text.InfoEnvFileNotFound)
			} else {
				initErr = fmt.Errorf("%s: %w", text.ErrEnvFileLoad, err)
				return
			}
		}

		cfg.BotToken = os.Getenv(text.EnvBotToken)
		if cfg.BotToken == "" {
			initErr = fmt.Errorf("%s '%s'", text.ErrEnvMissingVar, text.EnvBotToken)
			return
		}

		cfg.Timeout = 5 * time.Second
	})

	if initErr != nil {
		return nil, initErr
	}

	return &cfg, nil
}
