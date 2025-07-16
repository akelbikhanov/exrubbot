// Package version предоставляет информацию о версии приложения.
package version

import (
	"fmt"
	"runtime"
)

var (
	// Заполняется при сборке образа, через ldflags
	name      = "exrubbot"
	version   = "unknown"
	gitCommit = "unknown"
	buildTime = "unknown"
	goVersion = runtime.Version()
)

// GetVersion возвращает версию приложения
// За основу взят вывод версии из утилиты golangci-lint:
// golangci-lint has version 2.2.2 built with go1.24.4 from e9d42511 on 2025-07-11T12:00:50Z
func GetVersion() string {
	return fmt.Sprintf("%s has version %s built with %s from %s on %s",
		name, version, goVersion, gitCommit, buildTime)
}
