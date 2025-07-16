package bot

import (
	"fmt"
	"sort"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/akelbikhanov/exrubbot/pkg/entity"
	"github.com/go-telegram/bot/models"
)

// Интервалы для подписки
var intervals = [...]struct {
	seconds int
	label   string
}{
	{60, text.BtnInterval1M},
	{300, text.BtnInterval5M},
	{1800, text.BtnInterval30M},
	{3600, text.BtnInterval1H},
	{21600, text.BtnInterval6H},
	{86400, text.BtnInterval1D},
}

// intervalsKeyboard создаёт клавиатуру выбора интервала.
func intervalsKeyboard(feedID string) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{{
				Text:         text.BtnQuoteOnce,
				CallbackData: fmt.Sprintf(text.CallbackFormatFull, text.CallbackPrefixQuote, text.CallbackSeparator, feedID, text.CallbackSeparator, 0),
			}}, {{
				Text:         intervals[0].label,
				CallbackData: fmt.Sprintf(text.CallbackFormatInterval, text.CallbackPrefixQuote, feedID, intervals[0].seconds),
			}, {
				Text:         intervals[1].label,
				CallbackData: fmt.Sprintf(text.CallbackFormatInterval, text.CallbackPrefixQuote, feedID, intervals[1].seconds),
			}}, {{
				Text:         intervals[2].label,
				CallbackData: fmt.Sprintf(text.CallbackFormatInterval, text.CallbackPrefixQuote, feedID, intervals[2].seconds),
			}, {
				Text:         intervals[3].label,
				CallbackData: fmt.Sprintf(text.CallbackFormatInterval, text.CallbackPrefixQuote, feedID, intervals[3].seconds),
			}}, {{
				Text:         intervals[4].label,
				CallbackData: fmt.Sprintf(text.CallbackFormatInterval, text.CallbackPrefixQuote, feedID, intervals[4].seconds),
			}, {
				Text:         intervals[5].label,
				CallbackData: fmt.Sprintf(text.CallbackFormatInterval, text.CallbackPrefixQuote, feedID, intervals[5].seconds),
			}}, {{
				Text:         text.BtnBack,
				CallbackData: fmt.Sprintf(text.CallbackFormatFull, text.CallbackPrefixQuote, text.CallbackSeparator, feedID, text.CallbackSeparator, -1),
			}},
		},
	}
}

// feedsKeyboard создаёт клавиатуру выбора источника котировок.
func feedsKeyboard(feeds map[string]entity.Feed) *models.InlineKeyboardMarkup {
	if len(feeds) == 0 {
		return nil
	}

	// Сортируем источники по имени
	type item struct {
		id   string
		name string
	}

	items := make([]item, 0, len(feeds))
	for id, feed := range feeds {
		items = append(items, item{id: id, name: feed.Name()})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].name < items[j].name
	})

	// Создаём кнопки
	rows := make([][]models.InlineKeyboardButton, 0, len(items))
	for _, it := range items {
		btn := models.InlineKeyboardButton{
			Text:         it.name,
			CallbackData: fmt.Sprintf(text.CallbackFormatFeed, text.CallbackPrefixQuote, text.CallbackSeparator, it.id),
		}
		rows = append(rows, []models.InlineKeyboardButton{btn})
	}

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}
