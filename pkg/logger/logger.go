// Package logger инициализирует и предоставляет потокобезопасный доступ к логгеру приложения.
package logger

import (
	"log/slog"
	"os"
	"runtime"
)

// Logger is logger.
type Logger struct {
	logger *slog.Logger
}

// NewStdout создаёт новый экземпляр логгера с выводом в Stdout.
func NewStdout() *Logger {
	return &Logger{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

// Info логирует информационное сообщение.
func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

// Warn логирует предупреждение.
func (l *Logger) Warn(msg string, err error) {
	l.logger.Warn(msg, slog.Any("error", err))
}

// Error логирует ошибку с указанием функции-источника вызова.
// Если skip = 1, будет напечатана функция, откуда вызывался Error().
func (l *Logger) Error(msg string, err error, skip int) {
	pc, _, _, ok := runtime.Caller(skip)
	funcName := "?"
	if ok {
		if f := runtime.FuncForPC(pc); f != nil {
			funcName = f.Name()
		}
	}

	l.logger.Error(msg,
		slog.String("func", funcName),
		slog.Any("error", err),
	)
}
