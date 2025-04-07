package bot

import (
	"context"
	"errors"
	"os/signal"
	"sync"
	"syscall"

	"github.com/akelbikhanov/exrubbot/internal/bot/handler"
	"github.com/akelbikhanov/exrubbot/internal/common"
	"github.com/akelbikhanov/exrubbot/internal/config"
	"github.com/go-telegram/bot"
)

var (
	once    sync.Once
	initErr error
)

// Run инициализирует и запускает бота (гарантируя, что он создаётся только один раз)
// все последующие попытки возвращают ошибку ErrBotAlreadyRunning
func Run() error {
	initErr = errors.New(common.ErrBotAlreadyRunning)

	once.Do(func() {
		// Создаём контекст с graceful shutdown
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		h := handler.New(ctx)

		options := []bot.Option{
			bot.WithCheckInitTimeout(config.Get().DefaultTimeout),
			bot.WithDefaultHandler(h.DefaultHandler()),
			bot.WithErrorsHandler(h.ErrorsHandler(cancel, config.Get().DefaultTimeout)),
		}

		var b *bot.Bot
		b, initErr = bot.New(config.Get().BotToken, options...)
		if initErr != nil {
			return
		}

		// Передаем ссылку на бота в Handler.
		h.SetBot(b)

		// Запускаем бота
		b.Start(ctx)
	})

	return initErr
}
