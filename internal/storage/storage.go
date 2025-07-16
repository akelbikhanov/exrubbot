// Package storage обеспечивает загрузку и запись подписок в хранилище.
package storage

import "github.com/akelbikhanov/exrubbot/internal/entity"

// Storage абстрактный интерфейс хранилища.
type Storage interface {
	Load() ([]entity.Subscription, error)
	Save(items []entity.Subscription) error
}
