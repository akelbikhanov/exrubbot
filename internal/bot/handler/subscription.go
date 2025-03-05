package handler

import (
	"context"
	"github.com/akelbikhanov/garantex_bot/internal/service/garantex"
	"github.com/go-telegram/bot"
	"sync"
	"time"
)

// subscription хранит параметры подписки пользователя.
type subscription struct {
	interval time.Duration
	ticker   *time.Ticker
	stop     chan struct{}
	once     sync.Once
}

func (h *Handler) Subscribe(chatID int64, duration time.Duration) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Если подписка уже существует, останавливаем её.
	if sub, exists := h.subscriptions[chatID]; exists {
		sub.stopSubscription()
		delete(h.subscriptions, chatID)
	}

	// Создаем новую подписку.
	sub := &subscription{
		interval: duration,
		ticker:   time.NewTicker(duration),
		stop:     make(chan struct{}),
	}
	h.subscriptions[chatID] = sub

	// Запускаем горутину для рассылки уведомлений.
	go sub.run(h.ctx, h.b, chatID)
}

// Unsubscribe прекращает подписку для указанного chatID.
func (h *Handler) Unsubscribe(chatID int64) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	if sub, exists := h.subscriptions[chatID]; exists {
		sub.stopSubscription()
		delete(h.subscriptions, chatID)
		return true
	}

	return false
}

// run запускает цикл рассылки уведомлений для подписки.
func (s *subscription) run(ctx context.Context, b *bot.Bot, chatID int64) {
	for {
		select {
		case <-s.ticker.C:
			SendText(ctx, b, chatID, garantex.GetPriceText()+repeatEnding(s.interval), nil)
		case <-s.stop:
			s.ticker.Stop()
			return
		case <-ctx.Done():
			s.ticker.Stop()
			return
		}
	}
}

// stopSubscription сигнализирует о завершении работы подписки, закрывая канал stop.
func (s *subscription) stopSubscription() {
	s.once.Do(func() {
		close(s.stop)
	})
}
