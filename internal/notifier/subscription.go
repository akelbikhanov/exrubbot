package notifier

import (
	"context"
	"sync"
	"time"
)

// subscription хранит параметры подписки пользователя.
type subscription struct {
	interval time.Duration
	ticker   *time.Ticker
	send     SendFunc
	done     chan struct{}
	once     sync.Once
}

// start запускает рассылку уведомлений с заданным интервалом.
func (s *subscription) start(ctx context.Context) {
	for {
		select {
		case <-s.ticker.C:
			s.send()
		case <-s.done:
			s.ticker.Stop()
			return
		case <-ctx.Done():
			s.ticker.Stop()
			return
		}
	}
}

// terminate завершает работу подписки, закрывая канал done.
func (s *subscription) terminate() {
	s.once.Do(func() {
		close(s.done)
	})
}
