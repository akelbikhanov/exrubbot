// Package datafeed содержит реализации и регистрацию поставщиков рыночных данных.
package datafeed

import (
	"net/http"

	"github.com/akelbikhanov/exrubbot/pkg/entity"
)

// RegisterFeeds возвращает список всех зарегистрированных поставщиков данных.
func RegisterFeeds(hc *http.Client) map[string]entity.Feed {
	feeds := []entity.Feed{
		// NewABCEX(hc),
		NewGrinex(hc),
	}

	result := make(map[string]entity.Feed, len(feeds))
	for _, feed := range feeds {
		result[feed.ID()] = feed
	}
	return result
}
