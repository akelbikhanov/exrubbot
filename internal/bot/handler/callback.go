package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	kbRepeat = &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "10 сек", CallbackData: "interval:10s"},
				{Text: "10 мин", CallbackData: "interval:10m"},
			},
			{
				{Text: "30 мин", CallbackData: "interval:30m"},
				{Text: "1 час", CallbackData: "interval:1h"},
			},
		},
	}
)

// callbackHandler
func (h *Handler) callbackHandler(ctx context.Context, b *bot.Bot, cb *models.CallbackQuery) {
	if cb == nil {
		return
	}

	defer h.answerCallbackQuery(ctx, b, cb.ID)

	_, err := parseInterval(strings.TrimPrefix(cb.Data, "interval:"))
	if err != nil {
		h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, text.MessageUnknownCommand)
		return
	}

	// editText(ctx, b, cb.From.ID, cb.Message.Message.ID, garantex.GetPriceText()+repeatEnding(interval))
	// h.n.Subscribe(ctx, cb.From.ID, interval)
}

// parseInterval парсит строку интервала в time.Duration
func parseInterval(intervalStr string) (time.Duration, error) {
	switch intervalStr {
	case "10s":
		return 10 * time.Second, nil
	case "10m":
		return 10 * time.Minute, nil
	case "30m":
		return 30 * time.Minute, nil
	case "1h":
		return 1 * time.Hour, nil
	default:
		return 0, fmt.Errorf("неподдерживаемый интервал")
	}
}
