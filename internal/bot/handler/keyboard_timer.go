package handler

import (
	"fmt"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot/models"
)

// TimerButton структура для хранения интервала подписки (время и подпись кнопки).
type TimerButton struct {
	Sec   int
	Label string
}

// TimerButtons хранит список интервалов для подписки.
var TimerButtons = []TimerButton{
	{60, text.TimerButton1M},
	{300, text.TimerButton5M},
	{1800, text.TimerButton30M},
	{3600, text.TimerButton1H},
	{21600, text.TimerButton6H},
	{86400, text.TimerButton1D},
}

// keyboardTimer формирует клавиатуру выбора таймера.
func (h *Handler) keyboardTimer(feedID string) *models.InlineKeyboardMarkup {
	if feedID == "" {
		return nil
	}

	var rows [][]models.InlineKeyboardButton

	rows = append(rows, []models.InlineKeyboardButton{{
		Text:         text.TimerButtonOnce,
		CallbackData: fmt.Sprintf("%s%s%s%s%d", text.CallbackQuotePrefix, text.CallbackSeparator, feedID, text.CallbackSeparator, 0),
	}})

	if len(TimerButtons) != 6 {
		return nil
	}

	rows = append(rows, []models.InlineKeyboardButton{
		h.intervalButton(TimerButtons[0].Label, TimerButtons[0].Sec, feedID),
		h.intervalButton(TimerButtons[1].Label, TimerButtons[1].Sec, feedID),
	})

	rows = append(rows, []models.InlineKeyboardButton{
		h.intervalButton(TimerButtons[2].Label, TimerButtons[2].Sec, feedID),
		h.intervalButton(TimerButtons[3].Label, TimerButtons[3].Sec, feedID),
	})

	rows = append(rows, []models.InlineKeyboardButton{
		h.intervalButton(TimerButtons[4].Label, TimerButtons[4].Sec, feedID),
		h.intervalButton(TimerButtons[5].Label, TimerButtons[5].Sec, feedID),
	})

	rows = append(rows, []models.InlineKeyboardButton{{
		Text:         text.TimerButtonBack,
		CallbackData: fmt.Sprintf("%s%s%s%s%s", text.CallbackQuotePrefix, text.CallbackSeparator, feedID, text.CallbackSeparator, text.CallbackBack),
	}})

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}

func (h *Handler) intervalButton(label string, seconds int, feedID string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:         label,
		CallbackData: fmt.Sprintf("%s:%s:%d", text.CallbackQuotePrefix, feedID, seconds),
	}
}
