package handler

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/datafeed"
	"github.com/akelbikhanov/exrubbot/internal/logger"
	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot"
)

var msk = time.FixedZone("MSK", 3*60*60)

func (h *Handler) sendQuote(ctx context.Context, b *bot.Bot, chatID int64, name datafeed.FeedName, t time.Duration) {
	msg := text.MessageQuoteHead + string(name) + "\n%s\n\n" + currentTimeMSK()
	if t != 0 {
		msg += repeatSuffix(t)
	}

	quote, err := datafeed.GetQuote(ctx, name)
	if err != nil {
		logger.Error("", err)
		failText := fmt.Sprintf("%s\n%v", text.MessageQuoteFailed, err)
		h.sendText(ctx, b, chatID, fmt.Sprintf(msg, failText), nil)
		return
	}

	priceText := fmt.Sprintf(
		text.MessageQuoteFormat, // например: "🔼 %s (%s)\n🔽 %s (%s)"
		formatNumberOrDash(quote.AskPrice, 2),
		formatNumberOrDash(quote.AskVolume, 0),
		formatNumberOrDash(quote.BidPrice, 2),
		formatNumberOrDash(quote.BidVolume, 0),
	)

	h.sendText(ctx, b, chatID, fmt.Sprintf(msg, priceText), nil)
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
		return s // можно заменить на "-" если хочешь строгую проверку
	}
	return fmt.Sprintf("%.*f", precision, f)
}

func currentTimeMSK() string {
	now := time.Now().In(msk)
	return now.Format("2006.01.02 15:04:05")
}

func repeatSuffix(t time.Duration) string {
	seconds := int64(t.Seconds())
	if seconds < 60 {
		return fmt.Sprintf("%s%dс", text.MessageQuoteBottom, seconds)
	}

	minutes := int64(t.Minutes())
	if minutes < 60 {
		return fmt.Sprintf("%s%dм", text.MessageQuoteBottom, minutes)
	}

	hours := int64(t.Hours())
	if hours < 24 {
		return fmt.Sprintf("%s%dч", text.MessageQuoteBottom, hours)
	}

	days := hours / 24
	return fmt.Sprintf("%s%dд", text.MessageQuoteBottom, days)
}
