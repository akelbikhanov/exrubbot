package handler

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/akelbikhanov/exrubbot/pkg/entity"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// handleQuoteCallback обрабатывает callback-команды /quote.
// Возможные форматы:
// - quote:<feedID>           → показать выбор интервала
// - quote:<feedID>:0         → разовая котировка
// - quote:<feedID>:<seconds> → подписка
// - quote:<feedID>:back      → вернуться к выбору источника
func (h *Handler) handleQuoteCallback(ctx context.Context, b *bot.Bot, cb *models.CallbackQuery, parts []string) {
	switch len(parts) {
	case 2:
		h.sendTimerSelector(ctx, b, cb, parts[1])
	case 3:
		h.sendQuote(ctx, b, cb, parts[1], parts[2])
	default:
		msg := fmt.Sprintf(text.MessageCallbackIncorrect, cb.Data)
		h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, msg, nil)
	}
}

// sendTimerSelector
func (h *Handler) sendTimerSelector(ctx context.Context, b *bot.Bot, cb *models.CallbackQuery, feedID string) {
	feed, ok := h.feeds[feedID]
	if !ok {
		msg := fmt.Sprintf(text.MessageCallbackIncorrect, cb.Data)
		h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, msg, nil)
		return
	}

	kb := h.keyboardTimer(feed.ID())
	h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, text.MessageTimer, kb)
}

// sendQuote формирует и отправляет пользователю сообщение с котировкой.
func (h *Handler) sendQuote(ctx context.Context, b *bot.Bot, cb *models.CallbackQuery, feedID, timer string) {
	// quote:<feedID>:back → вернуть клавиатуру выбора источника
	if timer == text.CallbackBack {
		h.backToFeeds(ctx, b, cb.From.ID, cb.Message.Message.ID)
		return
	}

	feed, ok1 := h.feeds[feedID]
	if !ok1 {
		msg := fmt.Sprintf(text.MessageCallbackIncorrect, cb.Data)
		h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, msg, nil)
		return
	}

	interval, err := strconv.Atoi(timer)
	if err != nil {
		msg := fmt.Sprintf(text.MessageCallbackIncorrect, cb.Data)
		h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, msg, nil)
		return
	}

	// quote:<feedID>:0 → разовая котировка
	if interval == 0 {
		msg := h.buildQuoteMessage(ctx, feed)
		h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, msg, nil)
		return
	}

	// quote:<feedID>:<seconds> → подписка
	button, ok2 := findIntervalButton(interval)
	if !ok2 {
		msg := fmt.Sprintf(text.MessageCallbackIncorrect, cb.Data)
		h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, msg, nil)
	} else {
		sendFunc := func() {
			msg := h.buildQuoteMessageWithTimer(ctx, feed, button)
			h.sendText(ctx, b, cb.From.ID, msg, nil)
		}
		h.noty.Subscribe(ctx, cb.From.ID, time.Duration(interval)*time.Second, sendFunc)

		msg := h.buildQuoteMessageWithTimer(ctx, feed, button)
		h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, msg, nil)

	}
}

func findIntervalButton(sec int) (TimerButton, bool) {
	for _, btn := range TimerButtons {
		if btn.Sec == sec {
			return btn, true
		}
	}
	return TimerButton{}, false
}

// buildQuoteMessageWithTimer надстройка над buildQuoteMessage.
// Добавляет в конец сообщения информацию об интервале.
func (h *Handler) buildQuoteMessageWithTimer(ctx context.Context, feed entity.Feed, timer TimerButton) string {
	return h.buildQuoteMessage(ctx, feed) + fmt.Sprintf(text.QuoteStop, timer.Label)
}

// buildQuoteMessage формирует итоговое сообщение с котировкой
// для дальнейшей передачи либо в sendText, либо в editText
func (h *Handler) buildQuoteMessage(ctx context.Context, feed entity.Feed) string {
	msg := text.QuoteHead + feed.Name()
	quote, err := feed.GetQuote(ctx)
	if err != nil {
		h.logg.Error("", err, 2)
		msg += fmt.Sprintf("%s\n%v", text.QuoteFailed, err)
	} else {
		msg += fmt.Sprintf(
			text.QuoteFormat,
			formatNumberOrDash(quote.AskPrice, 2),
			formatNumberOrDash(quote.AskVolume, 0),
			formatNumberOrDash(quote.BidPrice, 2),
			formatNumberOrDash(quote.BidVolume, 0),
		)
	}

	msg += currentTimeMSK()

	return msg
}

// formatNumberOrDash форматирует числовую строку с заданной точностью.
// Если строка содержит нечисловое значение - возвращает как есть.
// Если строка пуста — возвращает дефис.
func formatNumberOrDash(s string, precision int) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "-"
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return s
	}
	return fmt.Sprintf("%.*f", precision, f)
}

// currentTimeMSK текущее время по Москве, в виде строки, в заданном формате.
func currentTimeMSK() string {
	const timeZone = " MSK (UTC+3)"
	msk := time.FixedZone("MSK", 3*60*60)
	now := time.Now().In(msk)
	return "\n\n" + now.Format(text.QuoteTimeLayout) + timeZone
}
func (h *Handler) sendFeeds(ctx context.Context, b *bot.Bot, chatID int64) {
	if kb := h.keyboardFeeds(); kb != nil {
		h.sendText(ctx, b, chatID, text.MessageFeed, kb)
	} else {
		h.sendText(ctx, b, chatID, text.MessageNoFeed, nil)
	}
}

func (h *Handler) backToFeeds(ctx context.Context, b *bot.Bot, chatID int64, msgID int) {
	if kb := h.keyboardFeeds(); kb != nil {
		h.editText(ctx, b, chatID, msgID, text.MessageFeed, kb)
	} else {
		h.editText(ctx, b, chatID, msgID, text.MessageNoFeed, nil)
	}
}
