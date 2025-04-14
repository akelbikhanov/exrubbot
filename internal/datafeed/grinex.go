package datafeed

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/akelbikhanov/exrubbot/internal/logger"
	"github.com/akelbikhanov/exrubbot/internal/text"
	"net/http"
)

// Response - структура для парсинга ответа API
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

// getGrinexQuote получает текущий биржи Grinex и приводит её к формату Quote.
func getGrinexQuote(ctx context.Context) (Quote, error) {
	data, err := getGrinexRAW(ctx)
	if err != nil {
		return Quote{}, err
	}

	if len(data.Asks) == 0 || len(data.Bids) == 0 {
		return Quote{}, fmt.Errorf(text.ErrRequestNoQuotes)
	}

	return Quote{
		AskPrice:  data.Asks[0].Price,
		AskVolume: data.Asks[0].Volume,
		BidPrice:  data.Bids[0].Price,
		BidVolume: data.Bids[0].Volume,
	}, nil
}

// getGrinexRAW получает текущий курс USDT/A7A5 с биржи Grinex.
// Вместо RUB используется нестандартный идентификатор A7A5.
func getGrinexRAW(ctx context.Context) (grinexResponse, error) {
	const GrinexAPIURL = "https://grinex.io/api/v2/depth?market=usdta7a5"
	var result grinexResponse

	// Формируем HTTP-запрос с привязкой к контексту.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, GrinexAPIURL, nil)
	if err != nil {
		return result, fmt.Errorf("%s: %w", text.ErrRequestBuild, err)
	}

	// Выполняем HTTP-запрос.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return result, fmt.Errorf("%s: %w", text.ErrRequestTimeout, err)
		case errors.Is(err, context.Canceled):
			return result, fmt.Errorf(text.ErrRequestCancelled)
		default:
			return result, fmt.Errorf("%s: %w", text.ErrRequestSend, err)
		}
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			logger.Error(text.ErrRequestClose, cerr)
		}
	}()

	// Проверяем код ответа.
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("%s: %d", text.ErrRequestAPI, resp.StatusCode)
	}

	// Декодируем JSON.
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("%s: %w", text.ErrDecodeJSON, err)
	}

	return result, nil
}
