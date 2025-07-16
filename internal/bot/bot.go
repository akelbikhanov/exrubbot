// Package bot содержит основную логику работы приложения (оркестратор).
package bot

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/scheduler"
	"github.com/akelbikhanov/exrubbot/internal/storage"
	"github.com/akelbikhanov/exrubbot/pkg/datafeed"
	"github.com/akelbikhanov/exrubbot/pkg/entity"
)

// Bot координирует работу всех компонентов.
type Bot struct {
	logger     *slog.Logger
	storage    storage.Storage
	scheduler  *scheduler.Scheduler
	feeds      map[string]entity.Feed
	retryDelay time.Duration
	stop       context.CancelFunc
	StopError  error
}

// Options параметры для создания бота.
type Options struct {
	Logger     *slog.Logger
	Storage    storage.Storage
	FeedClient *http.Client
	RetryDelay time.Duration
	Stop       context.CancelFunc
}

// New создаёт новый экземпляр бота.
func New(opts Options) *Bot {
	if opts.RetryDelay <= 0 {
		opts.RetryDelay = 5 * time.Second
	}

	return &Bot{
		logger:     opts.Logger,
		storage:    opts.Storage,
		scheduler:  scheduler.New(opts.Logger),
		feeds:      datafeed.RegisterFeeds(opts.FeedClient),
		retryDelay: opts.RetryDelay,
		stop:       opts.Stop,
	}
}
