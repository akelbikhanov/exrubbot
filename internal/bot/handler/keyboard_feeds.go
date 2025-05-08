package handler

import (
	"sort"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot/models"
)

// maxFeedsShown ограничивает количество кнопок/поставщиков.
// В дальнейшем можно сделать пагинацию.
const maxFeedsShown = 5

// keyboardFeeds формирует инлайн-клавиатуру из списка поставщиков данных.
// Возвращает nil, если список пуст.
func (h *Handler) keyboardFeeds() *models.InlineKeyboardMarkup {
	if len(h.feeds) == 0 {
		return nil
	}

	// Сбор и сортировка
	type pair struct {
		id   string
		name string
	}
	var sorted []pair
	for id, feed := range h.feeds {
		sorted = append(sorted, pair{id: id, name: feed.Name()})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].name < sorted[j].name
	})

	// Ограничение до maxFeedsShown
	if len(sorted) > maxFeedsShown {
		sorted = sorted[:maxFeedsShown]
	}

	// Генерация кнопок
	var rows [][]models.InlineKeyboardButton
	for _, f := range sorted {
		btn := models.InlineKeyboardButton{
			Text:         f.name,
			CallbackData: text.CallbackQuotePrefix + text.CallbackSeparator + f.id,
		}
		rows = append(rows, []models.InlineKeyboardButton{btn})
	}

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}
