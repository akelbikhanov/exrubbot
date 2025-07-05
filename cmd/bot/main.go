// Package main - точка входа в приложение.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/akelbikhanov/exrubbot/internal/bot"
	"github.com/akelbikhanov/exrubbot/internal/version"
)

// main запускает приложение.
func main() {
	versionFlag := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version.Full())
		os.Exit(0)
	}

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
