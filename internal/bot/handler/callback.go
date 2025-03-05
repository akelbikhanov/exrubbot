package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/akelbikhanov/garantex_bot/internal/common"
	"github.com/akelbikhanov/garantex_bot/internal/service/garantex"
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
func (h *Handler) callbackHandler(callback *models.CallbackQuery) {
	defer h.answerCallbackQuery(callback.ID)

	interval, err := parseInterval(strings.TrimPrefix(callback.Data, "interval:"))
	if err != nil {
		EditText(h.ctx, h.b, callback.From.ID, callback.Message.Message.ID, "")
		return
	}

	EditText(h.ctx, h.b, callback.From.ID, callback.Message.Message.ID, garantex.GetPriceText()+repeatEnding(interval))
	h.Subscribe(callback.From.ID, interval)
}

// answering callback query first to let Telegram know that we received the callback query,
// and we're handling it. Otherwise, Telegram might retry sending the update repetitively
// as it thinks the callback query doesn't reach to our application. learn more by
// reading the footnote of the https://core.telegram.org/bots/api#callbackquery type.
func (h *Handler) answerCallbackQuery(callbackID string) {
	_, err := h.b.AnswerCallbackQuery(h.ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callbackID,
		//Text:            "Готово",
		//ShowAlert:       false,
	})
	if err != nil {
		common.LogError(err)
	}
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

func repeatEnding(d time.Duration) string {
	const prefix = "\n" + common.CommandStop + " ⏱"
	seconds := int64(d.Seconds())
	if seconds < 60 {
		return fmt.Sprintf("%s%dс", prefix, seconds)
	}
	minutes := int64(d.Minutes())
	if minutes < 60 {
		return fmt.Sprintf("%s%dм", prefix, minutes)
	}
	hours := int64(d.Hours())
	if hours < 24 {
		return fmt.Sprintf("%s%dч", prefix, hours)
	}
	days := hours / 24
	return fmt.Sprintf("%s%dд", prefix, days)
}
