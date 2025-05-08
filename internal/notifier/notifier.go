// Package notifier управляет подписками с периодической отправкой уведомлений.
package notifier

import (
	"context"
	"sync"
	"time"
)

// SendFunc — функция отправки сообщений подписчику.
type SendFunc func()

// Notifier отвечает за обработку обновлений и управление подписками.
type Notifier struct {
	mu            sync.Mutex
	subscriptions map[int64]*subscription // ключ — chatID
}

// New создаёт новый экземпляр Notifier с инициализированной мапой подписок.
func New() *Notifier {
	return &Notifier{
		subscriptions: make(map[int64]*subscription),
	}
}

// Subscribe создаёт новую подписку.
func (n *Notifier) Subscribe(ctx context.Context, chatID int64, duration time.Duration, send SendFunc) {
	n.mu.Lock()
	defer n.mu.Unlock()

	// Если подписка уже существует, останавливаем её.
	if sub, exists := n.subscriptions[chatID]; exists {
		sub.terminate()
		delete(n.subscriptions, chatID)
	}

	// Создаем новую подписку.
	sub := &subscription{
		interval: duration,
		ticker:   time.NewTicker(duration),
		send:     send,
		done:     make(chan struct{}),
	}
	n.subscriptions[chatID] = sub

	// Запускаем горутину для рассылки уведомлений.
	go sub.start(ctx)
}

// Unsubscribe прекращает подписку для указанного chatID.
func (n *Notifier) Unsubscribe(chatID int64) bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	if sub, exists := n.subscriptions[chatID]; exists {
		sub.terminate()
		delete(n.subscriptions, chatID)
		return true
	}

	return false
}
