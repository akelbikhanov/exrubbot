// Package datafeed предоставляет унифицированный интерфейс получения рыночных данных
// от внешних источников (бирж, агрегаторов и прочих провайдеров).
package datafeed

import (
	"context"
	"fmt"
	"github.com/akelbikhanov/exrubbot/internal/text"
)

// Quote содержит цену на покупку и продажу с указанием объёмов.
type Quote struct {
	AskPrice  string
	AskVolume string
	BidPrice  string
	BidVolume string
}

// FeedName — строковое имя поставщика.
type FeedName string

const (
	Grinex FeedName = "Grinex"
)

// AvailableFeeds возвращает список поддерживаемых источников данных.
func AvailableFeeds() []FeedName {
	return []FeedName{
		Grinex,
	}
}

// GetQuote получает последнюю (актуальную) котировку от указанного источника.
func GetQuote(ctx context.Context, name FeedName) (Quote, error) {
	switch name {
	case Grinex:
		return getGrinexQuote(ctx)
	// case Binance:
	//	return getBinancePrice(ctx)
	default:
		return Quote{}, fmt.Errorf("%s '%s'", text.ErrDataFeedName, name)
	}
}
