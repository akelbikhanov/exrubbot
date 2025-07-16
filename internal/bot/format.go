package bot

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/akelbikhanov/exrubbot/pkg/entity"
)

// formatQuote форматирует полученные котировки
// в строку вида "\n🔼 %s (%s)\n🔽 %s (%s)"
func formatQuote(q entity.Quote) string {
	var result strings.Builder

	result.WriteString("\n")
	result.WriteString(text.SymbolAskPrice)
	result.WriteString(" ")
	result.WriteString(formatNumber(q.AskPrice, 2))
	result.WriteString(" (")
	result.WriteString(formatNumber(q.AskVolume, 0))
	result.WriteString(")\n")
	result.WriteString(text.SymbolBidPrice)
	result.WriteString(" ")
	result.WriteString(formatNumber(q.BidPrice, 2))
	result.WriteString(" (")
	result.WriteString(formatNumber(q.BidVolume, 0))
	result.WriteString(")")

	return result.String()
}

// formatNumber форматирует числовую строку с заданной точностью.
// Если строка содержит нечисловое значение - возвращает как есть.
// Если строка пуста — возвращает дефис.
func formatNumber(price string, precision int) string {
	price = strings.TrimSpace(price)
	if price == "" {
		return "-"
	}

	if f, err := strconv.ParseFloat(price, 64); err == nil {
		return fmt.Sprintf("%.*f", precision, f)
	}
	return price
}

// formatInterval форматирует интервал времени.
func formatInterval(seconds int) string {
	if seconds == 0 {
		return text.SuffixZero
	}

	var result strings.Builder

	if days := seconds / 86400; days > 0 {
		result.WriteString(strconv.Itoa(days))
		result.WriteString(text.SuffixDays)
		seconds %= 86400
	}
	if hours := seconds / 3600; hours > 0 {
		result.WriteString(strconv.Itoa(hours))
		result.WriteString(text.SuffixHours)
		seconds %= 3600
	}
	if minutes := seconds / 60; minutes > 0 {
		result.WriteString(strconv.Itoa(minutes))
		result.WriteString(text.SuffixMinutes)
		seconds %= 60
	}
	if seconds > 0 {
		result.WriteString(strconv.Itoa(seconds))
		result.WriteString(text.SuffixSeconds)
	}

	return result.String()
}

// formatTimeMSK форматирует текущее время по Москве.
func formatTimeMSK() string {
	msk := time.FixedZone(text.QuoteTimeZone, 3*60*60)
	now := time.Now().In(msk)
	return now.Format(text.QuoteTimeFormat)
}
