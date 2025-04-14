// Package handler содержит обработчики запросов от пользователей Telegram.
package handler

import (
	"context"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/notifier"
)

// Handler отвечает за обработку обновлений и управление подписками.
type Handler struct {
	n           *notifier.Notifier
	cancel      context.CancelFunc
	timeout     time.Duration
	CancelError error
}

// New создаёт новый экземпляр Handler с инициализированной мапой подписок.
func New(cancel context.CancelFunc, timeout time.Duration) *Handler {
	return &Handler{
		n:       notifier.New(),
		cancel:  cancel,
		timeout: timeout,
	}
}
