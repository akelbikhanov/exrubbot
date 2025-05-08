package datafeed

import (
	"context"
	"fmt"
	"net/http"

	"github.com/akelbikhanov/exrubbot/pkg/entity"
	"github.com/akelbikhanov/exrubbot/pkg/text"
)

const (
	abcexID     = "abcex"
	abcexName   = "ABCEX"
	abcexAPIURL = "https://abcex.io/..."
)

// abcexFeed реализует интерфейс entity.Feed.
// Поддерживается только пара USDT/RUB.
type abcexFeed struct {
	hc *http.Client
}

// NewABCEX возвращает экземпляр grinexFeed.
func NewABCEX(hc *http.Client) entity.Feed {
	return &abcexFeed{
		hc: hc,
	}
}

// Name возвращает название источника данных.
func (f *abcexFeed) Name() string {
	return abcexName
}

// ID возвращает название источника данных.
func (f *abcexFeed) ID() string {
	return abcexID
}

// GetQuote возвращает рыночную котировку для валютной пары USDT/RUB.
func (f *abcexFeed) GetQuote(_ context.Context) (entity.Quote, error) {
	return entity.Quote{}, fmt.Errorf(text.ErrFeedNotSupport)
}
