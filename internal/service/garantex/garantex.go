package garantex

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/akelbikhanov/garantex_bot/internal/common"
)

// Response - структура для парсинга ответа API
type response struct {
	Bids []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
	} `json:"bids"`
	Asks []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
	} `json:"asks"`
}

func GetPriceText() string {
	prices, err := getUSDTPrice()
	if err != nil {
		return common.ErrGarantex + err.Error()
	}

	return formatMessage(prices)
}

// getUSDTPrice получает текущий курс USDT/RUB
func getUSDTPrice() (response, error) {
	const garantexAPI = "https://garantex.org/api/v2/depth?market=usdtrub"
	var result response

	client := &http.Client{Timeout: common.DefaultTimeout}
	resp, err := client.Get(garantexAPI)
	if err != nil {
		return result, fmt.Errorf(common.ErrGarantexRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf(common.ErrGarantexAPI, resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, fmt.Errorf(common.ErrGarantexJSON, err)
	}

	if len(result.Bids) == 0 || len(result.Asks) == 0 {
		return result, fmt.Errorf(common.ErrGarantexNoQuotes)
	}

	return result, nil
}

// formatMessage форматирует сообщение с данными
func formatMessage(prices response) string {
	askPrice, _ := strconv.ParseFloat(prices.Asks[0].Price, 64)
	askVolume, _ := strconv.ParseFloat(prices.Asks[0].Volume, 64)
	bidPrice, _ := strconv.ParseFloat(prices.Bids[0].Price, 64)
	bidVolume, _ := strconv.ParseFloat(prices.Bids[0].Volume, 64)

	askVolumeRounded := math.Round(askVolume)
	bidVolumeRounded := math.Round(bidVolume)
	moscowTime := time.Now().In(time.FixedZone("MSK", 3*60*60))

	// Формат без экранирования вручную, так как используем EscapeMarkdown
	return fmt.Sprintf(common.MessagePrice,
		askPrice, askVolumeRounded,
		bidPrice, bidVolumeRounded,
		moscowTime.Format("2006.01.02 15:04:05"), // Оставляем без экранирования, EscapeMarkdown обработает
	)
}
