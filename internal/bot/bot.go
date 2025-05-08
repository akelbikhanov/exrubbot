// Package bot инициализирует и запускает бота.
package bot

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/akelbikhanov/exrubbot/internal/bot/handler"
	"github.com/akelbikhanov/exrubbot/internal/config"
	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/akelbikhanov/exrubbot/pkg/datafeed"
	"github.com/akelbikhanov/exrubbot/pkg/logger"
	"github.com/go-telegram/bot"
)

var (
	once sync.Once
)

// RunOnce - singleton обвязка.
func RunOnce(ctx context.Context) (err error) {
	alreadyStarted := true
	defer func() {
		if alreadyStarted {
			err = errors.New(text.ErrBotAlreadyRunning)
		}
	}()

	once.Do(func() {
		alreadyStarted = false

		err = run(ctx)
	})

	return err
}

// run инициализирует и запускает бота.
func run(parentCtx context.Context) error {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	var (
		cfg *config.Config
		b   *bot.Bot
		err error
	)

	l := logger.NewStdout()

	if cfg, err = config.Get(l); err != nil {
		l.Error("", err, 1)
		return err
	}

	tr := &http.Transport{}
	botClient := &http.Client{
		Timeout:   cfg.BotTimeout,
		Transport: tr,
	}
	feedClient := &http.Client{
		Timeout:   cfg.FeedTimeout,
		Transport: tr,
	}

	f := datafeed.RegisterFeeds(feedClient)
	h := handler.New(l, f, cfg.BotRetry, cancel)

	options := []bot.Option{
		bot.WithHTTPClient(cfg.BotTimeout, botClient),
		bot.WithNotAsyncHandlers(),
		bot.WithDefaultHandler(h.DefaultHandler()),
		bot.WithErrorsHandler(h.ErrorHandler()),
	}

	if b, err = bot.New(cfg.BotToken, options...); err != nil {
		return err
	}

	b.Start(ctx)

	if h.TerminateReason != nil {
		err = h.TerminateReason
	}
	return err
}
