// Package scheduler управляет подписками пользователей на периодические котировки.
package scheduler

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/entity"
	"github.com/akelbikhanov/exrubbot/internal/text"
)

// task структура задачи для планировщика.
type task struct {
	s      entity.Subscription
	send   func()
	ticker *time.Ticker
	done   chan struct{}
	once   sync.Once
	logger *slog.Logger
}

// Scheduler управляет всеми активными подписками.
type Scheduler struct {
	mu     sync.RWMutex
	subs   map[int64]*task // ключ - chatID
	logger *slog.Logger
}

// New создаёт новый менеджер подписок.
func New(logger *slog.Logger) *Scheduler {
	return &Scheduler{
		subs:   make(map[int64]*task),
		logger: logger,
	}
}

// Subscribe создаёт или обновляет задачу/подписку.
// Если подписка уже существует, она будет остановлена и заменена новой.
func (s *Scheduler) Subscribe(ctx context.Context, sub entity.Subscription, send func()) error {
	if sub.IntervalSec <= 0 {
		return fmt.Errorf(text.LogErrInvalidInterval, sub.IntervalSec)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Останавливаем существующую задачу
	if existing, ok := s.subs[sub.ChatID]; ok {
		existing.stop()
		delete(s.subs, sub.ChatID)
	}

	// Создаём новую задачу
	t := &task{
		s:      sub,
		send:   send,
		ticker: time.NewTicker(time.Duration(sub.IntervalSec) * time.Second),
		done:   make(chan struct{}),
	}

	s.subs[sub.ChatID] = t
	go t.run(ctx)

	// Логируем запуск задачи/планировщика
	s.logger.Debug(text.LogDebugSchedulerStart,
		text.LogFieldChatID, sub.ChatID,
		text.LogFieldFeedID, sub.FeedID,
		text.LogFieldIntervalSec, sub.IntervalSec)

	return nil
}

// Unsubscribe отменяет подписку пользователя.
// Возвращает true, если подписка существовала.
func (s *Scheduler) Unsubscribe(chatID int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.subs[chatID]
	if !ok {
		return false
	}

	t.stop()
	delete(s.subs, chatID)

	// Логируем остановку задачи/планировщика
	s.logger.Debug(text.LogDebugSchedulerStop, text.LogFieldChatID, chatID)

	return true
}

// GetSubscriptions возвращает слайс активных подписок.
func (s *Scheduler) GetSubscriptions() []entity.Subscription {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]entity.Subscription, 0, len(s.subs))
	for _, t := range s.subs {
		// Создаём копию для избежания race condition
		result = append(result, entity.Subscription{
			ChatID:      t.s.ChatID,
			FeedID:      t.s.FeedID,
			IntervalSec: t.s.IntervalSec,
			CreatedAt:   t.s.CreatedAt,
		})
	}
	return result
}

// run запускает периодическую отправку котировок.
func (t *task) run(ctx context.Context) {
	// Добавляем recover для защиты от паники в send()
	defer func() {
		if r := recover(); r != nil {
			t.logger.Error(text.LogErrSchedulerRun, text.LogFieldChatID, t.s.ChatID)
		}
		t.ticker.Stop()
	}()

	for {
		select {
		case <-t.ticker.C:
			t.send()
		case <-t.done:
			return
		case <-ctx.Done():
			return
		}
	}
}

// stop останавливает подписку.
func (t *task) stop() {
	t.once.Do(func() {
		close(t.done)
	})
}
