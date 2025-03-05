package handler

import (
	"context"
	"sync"

	"github.com/go-telegram/bot"
)

// Handler отвечает за обработку обновлений и управление подписками.
type Handler struct {
	ctx context.Context
	b   *bot.Bot

	mu            sync.Mutex
	subscriptions map[int64]*subscription // ключ — chatID
}

// New создаёт новый экземпляр Handler с инициализированной мапой подписок.
func New(ctx context.Context) *Handler {
	return &Handler{
		ctx:           ctx,
		subscriptions: make(map[int64]*subscription),
	}
}

// SetBot устанавливает указатель на инициализированного бота.
func (h *Handler) SetBot(b *bot.Bot) {
	h.b = b
}
