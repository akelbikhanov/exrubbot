package config

import (
	"fmt"
	"sync"
	"syscall"
	"time"

	"github.com/akelbikhanov/garantex_bot/internal/common"
	"github.com/joho/godotenv"
)

// Config - структура для хранения конфигурации приложения
type Config struct {
	BotToken       string        // Токен Telegram бота
	DefaultTimeout time.Duration // Таймаут запросов по умолчанию
}

var (
	cfg  Config
	once sync.Once
)

// Get загружает переменные окружения из файла .env (если он есть)
// а затем пытается их прочесть из ОС
func Get() *Config {
	once.Do(func() {
		var (
			exists bool
			err    error
		)

		// подгружаем переменные окружения из файла (если есть)
		if err = godotenv.Load(); err != nil {
			common.LogError(fmt.Errorf(common.ErrLoadEnv, err))
		}

		cfg.BotToken, exists = syscall.Getenv(common.EnvBotToken)
		if !exists {
			common.LogError(fmt.Errorf(common.ErrMissingEnvVar, common.EnvBotToken))
		}

		cfg.DefaultTimeout = common.DefaultTimeout
	})
	return &cfg
}
