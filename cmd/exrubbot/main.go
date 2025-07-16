// Package main - точка входа в приложение.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/akelbikhanov/exrubbot/internal/exrubbot"
	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/akelbikhanov/exrubbot/internal/version"
)

// main инициализирует конфигурацию, логгер и запускает приложение.
func main() {
	versionFlag := flag.Bool(text.FlagVersion, false, text.FlagVersionDesc)
	logLevelFlag := flag.Int(text.FlagLogLevel, 0, text.FlagLogLevelDesc)
	flag.Parse()

	if *versionFlag {
		fmt.Println(version.GetVersion())
		os.Exit(0)
	}

	if err := exrubbot.Run(*logLevelFlag); err != nil {
		os.Exit(1)
	}
}
