package datafeed

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akelbikhanov/exrubbot/pkg/entity"
	"github.com/akelbikhanov/exrubbot/pkg/text"
)

const (
	grinexID     = "usdtrub_grinex"
	grinexName   = "USDT/RUB • Grinex"
	grinexAPIURL = "https://grinex.io/api/v2/depth?market=usdtrub"
)

// grinexFeed реализует интерфейс entity.Feed.
// Поддерживается только пара USDT/RUB.
type grinexFeed struct {
	hc *http.Client
}

// NewGrinex возвращает экземпляр grinexFeed.
func NewGrinex(hc *http.Client) entity.Feed {
	return &grinexFeed{
		hc: hc,
	}
}

// Name возвращает название источника данных.
func (f *grinexFeed) Name() string {
	return grinexName
}

// ID возвращает название источника данных.
func (f *grinexFeed) ID() string {
	return grinexID
}

// GetQuote возвращает рыночную котировку для валютной пары USDT/RUB.
func (f *grinexFeed) GetQuote(ctx context.Context) (entity.Quote, error) {
	var result grinexResponse

	// Формируем HTTP-запрос с привязкой к контексту.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, grinexAPIURL, nil)
	if err != nil {
		return entity.Quote{}, fmt.Errorf("%s: %w", text.ErrRequestBuild, err)
	}

	// Выполняем HTTP-запрос.
	resp, err := f.hc.Do(req)
	if err != nil {
		return entity.Quote{}, fmt.Errorf("%s: %w", text.ErrRequestSend, err)
	}

	// Проверяем код ответа.
	if resp.StatusCode != http.StatusOK {
		return entity.Quote{}, fmt.Errorf("%s: %d", text.ErrRequestStatus, resp.StatusCode)
	}

	// Проверяем тело ответа.
	if resp.Body == nil {
		return entity.Quote{}, fmt.Errorf(text.ErrRequestBody)
	}

	// Декодируем JSON.
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return entity.Quote{}, fmt.Errorf("%s: %w", text.ErrDecodeJSON, err)
	}

	// Закрываем соединение и делаем это явно здесь, а не в defer, чтобы при возникновении
	// ошибки закрытия соединения не перезаписать ею - другую ошибку, возникшую в этом блоке ранее.
	if err = resp.Body.Close(); err != nil {
		return entity.Quote{}, fmt.Errorf("%s: %w", text.ErrRequestClose, err)
	}

	// Проверяем содержимое. Возможно структура пустая.
	if len(result.Asks) == 0 || len(result.Bids) == 0 {
		return entity.Quote{}, fmt.Errorf(text.ErrRequestNoQuotes)
	}

	return entity.Quote{
		AskPrice:  result.Asks[0].Price,
		AskVolume: result.Asks[0].Volume,
		BidPrice:  result.Bids[0].Price,
		BidVolume: result.Bids[0].Volume,
	}, nil
}

// grinexResponse описывает структуру JSON-ответа от Grinex API.
// Структура ответа актуальна по состоянию на 14.04.2025.
type grinexResponse struct {
	Bids []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
	} `json:"bids"`
	Asks []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
	} `json:"asks"`
}
