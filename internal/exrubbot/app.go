// Package exrubbot инициализирует и запускает приложение.
package exrubbot

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/akelbikhanov/exrubbot/internal/bot"
	"github.com/akelbikhanov/exrubbot/internal/config"
	"github.com/akelbikhanov/exrubbot/internal/storage"
	"github.com/akelbikhanov/exrubbot/internal/text"
	tg "github.com/go-telegram/bot"
)

// Run инициализирует и запускает приложение.
func Run(logLevel int) error {
	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Создаём логгер
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(logLevel)}))

	// Читаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error(text.LogErrConfig, text.LogFieldError, err)
		return err
	}

	// Создаём HTTP клиенты
	clients := createHTTPClients(cfg)

	// Создаём хранилище
	store := storage.NewFileStorage(cfg.SubsPath)

	// Создаём главный класс приложения
	b := bot.New(bot.Options{
		Logger:     logger,
		Storage:    store,
		FeedClient: clients.feed,
		RetryDelay: cfg.BotRetry,
		Stop:       stop,
	})

	// Создаём Telegram клиента
	telegram, err1 := createBot(cfg, clients.tg, b)
	if err1 != nil {
		logger.Error(text.LogErrTelegramInit, text.LogFieldError, err1)
		return err1
	}

	// Восстанавливаем подписки
	b.RestoreSubscriptions(ctx, telegram)

	// Запускаем асинхронно бота и сохранение подписок
	logger.Info(text.LogInfoStarting)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		telegram.Start(ctx)
	}()
	go func() {
		defer wg.Done()
		b.SaveSubscriptions(ctx)
	}()

	wg.Wait()
	logger.Info(text.LogInfoStopped)

	// Возвращаем ошибку из Bot
	return b.StopError
}

// httpClients группирует HTTP клиенты для разных целей.
type httpClients struct {
	tg   *http.Client
	feed *http.Client
}

// createHTTPClients создает настроенные HTTP клиенты.
func createHTTPClients(cfg config.Config) httpClients {
	transport := &http.Transport{}

	return httpClients{
		tg: &http.Client{
			Timeout:   cfg.BotTimeout,
			Transport: transport,
		},
		feed: &http.Client{
			Timeout:   cfg.FeedTimeout,
			Transport: transport,
		},
	}
}

// createBot создает и настраивает экземпляр Telegram клиента.
func createBot(cfg config.Config, client *http.Client, b *bot.Bot) (*tg.Bot, error) {
	options := []tg.Option{
		tg.WithHTTPClient(cfg.BotTimeout, client),
		tg.WithNotAsyncHandlers(),
		tg.WithDefaultHandler(b.HandleUpdate),
		tg.WithErrorsHandler(b.HandleError),
	}

	return tg.New(cfg.BotToken, options...)
}
