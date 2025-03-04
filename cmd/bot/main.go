package main

import (
	"github.com/akelbikhanov/garantex_bot/internal/bot"
	"github.com/akelbikhanov/garantex_bot/internal/common"
	"os"
)

func main() {
	if err := bot.Run(); err != nil {
		common.LogError(err)
		os.Exit(1)
	}
}
