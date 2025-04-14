// Package logger инициализирует и предоставляет потокобезопасный доступ к логгеру приложения.
package logger

import (
	"log/slog"
	"os"
	"runtime"
	"sync/atomic"
)

// currentLogger — потокобезопасное хранилище активного логгера.
var currentLogger atomic.Value

// Инициализируем логгер по умолчанию (stdout, текстовый формат).
func init() {
	defaultLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	currentLogger.Store(defaultLogger)
}

// Init переопределяет глобальный логгер, используя переданный handler.
func Init(handler slog.Handler) {
	logger := slog.New(handler)
	currentLogger.Store(logger)
}

// Info логирует информационное сообщение.
func Info(msg string) {
	logger := currentLogger.Load().(*slog.Logger)
	logger.Info(msg)
}

// Warn логирует предупреждение.
func Warn(msg string, err error) {
	logger := currentLogger.Load().(*slog.Logger)
	logger.Warn(msg, slog.Any("error", err))
}

// Error логирует ошибку с указанием функции-источника вызова.
func Error(msg string, err error) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "?"
	if ok {
		if f := runtime.FuncForPC(pc); f != nil {
			funcName = f.Name()
		}
	}

	logger := currentLogger.Load().(*slog.Logger)
	logger.Error(msg,
		slog.String("func", funcName),
		slog.Any("error", err),
	)
}
