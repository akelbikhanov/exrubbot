// Package bot инициализирует и запускает бота.
package bot

import (
	"context"
	"errors"
	"sync"

	"github.com/akelbikhanov/exrubbot/internal/bot/handler"
	"github.com/akelbikhanov/exrubbot/internal/config"
	"github.com/akelbikhanov/exrubbot/internal/logger"
	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot"
)

var (
	once sync.Once
	cfg  *config.Config
	b    *bot.Bot
	err  error
)

// Run инициализирует и запускает бота в единственном экземпляре.
func Run(parentCtx context.Context) error {
	alreadyStarted := true

	once.Do(func() {
		alreadyStarted = false

		if cfg, err = config.Get(); err != nil {
			logger.Error("", err)
			return
		}

		ctx, cancel := context.WithCancel(parentCtx)
		defer cancel()

		h := handler.New(cancel, cfg.Timeout)

		options := []bot.Option{
			bot.WithNotAsyncHandlers(),
			bot.WithCheckInitTimeout(cfg.Timeout),
			bot.WithDefaultHandler(h.DefaultHandler()),
			bot.WithErrorsHandler(h.ErrorHandler()),
		}

		if b, err = bot.New(cfg.BotToken, options...); err != nil {
			return
		}

		b.Start(ctx)

		if h.CancelError != nil {
			err = h.CancelError
		}
	})

	if alreadyStarted {
		return errors.New(text.ErrBotAlreadyRunning)
	}

	return err
}
