package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GarantexResponse - структура для парсинга ответа API
type GarantexResponse struct {
	Bids []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
	} `json:"bids"`
	Asks []struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
	} `json:"asks"`
}

// GetUSDTPrice получает текущий курс USDT/RUB
func GetUSDTPrice() (buyPrice, sellPrice string, err error) {
	const garantexAPI = "https://garantex.org/api/v2/depth?market=usdtrub"

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(garantexAPI)
	if err != nil {
		return "", "", fmt.Errorf("ошибка запроса к API Garantex: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("ошибка API Garantex: код %d", resp.StatusCode)
	}

	var result GarantexResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	if len(result.Bids) == 0 || len(result.Asks) == 0 {
		return "", "", fmt.Errorf("недостаточно данных в ответе API")
	}

	return result.Bids[0].Price, result.Asks[0].Price, nil
}
