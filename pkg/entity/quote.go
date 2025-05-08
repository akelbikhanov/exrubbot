// Package entity содержит основные бизнес-сущности, используемые для работы с рыночными данными.
package entity

// Quote представляет собой рыночную котировку, полученную от поставщика данных.
type Quote struct {
	AskPrice  string
	AskVolume string
	BidPrice  string
	BidVolume string
}
