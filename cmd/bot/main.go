package main

import (
	"os"

	"github.com/akelbikhanov/exrubbot/internal/bot"
	"github.com/akelbikhanov/exrubbot/internal/common"
)

func main() {
	if err := bot.Run(); err != nil {
		common.LogError(err)
		os.Exit(1)
	}
}
