package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/akelbikhanov/garantex_bot/internal/common"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// DefaultHandler обработчик по-умолчанию
// первая точка входа обработки всех запросов
func (h *Handler) DefaultHandler() bot.HandlerFunc {
	return func(_ context.Context, _ *bot.Bot, update *models.Update) {
		start := time.Now()
		defer func() {
			elapsed := time.Since(start)
			common.LogInfo(fmt.Sprintf(common.InfoUpdateProcessed, update.ID, elapsed))
		}()

		switch {
		case update.Message != nil:
			h.messageHandler(update.Message)
		case update.CallbackQuery != nil:
			h.callbackHandler(update.CallbackQuery)
		default:
			common.LogInfo(fmt.Sprintf(common.InfoUpdateSkip, update.ID))
		}
	}
}
