// Package entity содержит основные бизнес-сущности, используемые в работе бота.
package entity

import "time"

// Subscription базовая структура подписки.
type Subscription struct {
	ChatID      int64     `json:"chat_id"`
	FeedID      string    `json:"feed_id"`
	IntervalSec int       `json:"interval_sec"`
	CreatedAt   time.Time `json:"created_at"`
}
