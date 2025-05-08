// Package main - точка входа в приложение.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/akelbikhanov/exrubbot/internal/bot"
)

// main запускает приложение.
func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

// run выполняет инициализацию контекста и передаёт управление боту.
func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	return bot.RunOnce(ctx)
}
