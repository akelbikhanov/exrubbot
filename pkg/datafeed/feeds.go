// Package datafeed содержит реализации и регистрацию поставщиков рыночных данных.
package datafeed

import (
	"net/http"

	"github.com/akelbikhanov/exrubbot/pkg/entity"
)

// RegisterFeeds возвращает список всех зарегистрированных поставщиков данных.
func RegisterFeeds(hc *http.Client) map[string]entity.Feed {
	return map[string]entity.Feed{
		// NewABCEX(hc),
		grinexRubID:  NewGrinex(hc, grinexRubID, grinexRubName, grinexRubAPIURL),
		grinexA7A5ID: NewGrinex(hc, grinexA7A5ID, grinexA7A5Name, grinexA7A5APIURL),
	}
}
