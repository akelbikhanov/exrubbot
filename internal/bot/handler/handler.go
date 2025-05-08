// Package handler содержит обработчики запросов от пользователей Telegram.
package handler

import (
	"context"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/notifier"
	"github.com/akelbikhanov/exrubbot/pkg/entity"
	"github.com/akelbikhanov/exrubbot/pkg/logger"
)

// Handler отвечает за обработку обновлений и управление подписками.
type Handler struct {
	logg            *logger.Logger
	noty            *notifier.Notifier
	feeds           map[string]entity.Feed
	retryTime       time.Duration
	cancel          context.CancelFunc
	TerminateReason error
}

// New создаёт новый экземпляр Handler с инициализированной мапой подписок.
func New(logg *logger.Logger, feeds map[string]entity.Feed, retryTime time.Duration, cancel context.CancelFunc) *Handler {
	return &Handler{
		logg:      logg,
		noty:      notifier.New(),
		feeds:     feeds,
		retryTime: retryTime,
		cancel:    cancel,
	}
}
