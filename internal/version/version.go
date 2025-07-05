// Package version предоставляет информацию о сборке приложения.
package version

import (
	"fmt"
	"runtime"
)

var (
	// Заполняется при сборке через ldflags
	version   = "unknown"
	gitCommit = "unknown"
	buildTime = "unknown"
	goVersion = runtime.Version()
)

// Short возвращает краткую версию приложения
func Short() string {
	return fmt.Sprintf("%s, %s", version, goVersion)
}

// Full возвращает полную информацию о версии
func Full() string {
	return fmt.Sprintf("Version: %s\nCommit: %s\nBuild: %s\nGo: %s\n",
		version, gitCommit, buildTime, goVersion)
}
