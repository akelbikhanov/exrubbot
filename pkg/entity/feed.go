package entity

import "context"

// Feed описывает интерфейс поставщика рыночных данных.
type Feed interface {
	ID() string
	Name() string
	GetQuote(context.Context) (Quote, error)
}
