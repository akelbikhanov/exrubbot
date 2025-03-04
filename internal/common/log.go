package common

import (
	"log/slog"
	"os"
	"runtime"
)

// Логгер (по умолчанию пишет в stderr, но можно менять)
var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

// InitLogger позволяет задать кастомный логгер (например, в файл)
func InitLogger(handler slog.Handler) {
	logger = slog.New(handler)
}

// LogError логирует ошибку с названием функции
func LogError(e error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)

	if ok && details != nil {
		logger.Error("Ошибка", slog.String("функция", details.Name()), slog.Any("error", e))
	} else {
		logger.Error("Ошибка", slog.String("функция", "?"), slog.Any("error", e))
	}
}

// LogInfo логирует информационные сообщения
func LogInfo(msg string) {
	logger.Info(msg)
}

// LogWarning логирует предупреждения
func LogWarning(msg string) {
	logger.Warn(msg)
}
