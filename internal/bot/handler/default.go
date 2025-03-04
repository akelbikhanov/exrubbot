package handler

import (
	"context"
	"fmt"
	"github.com/akelbikhanov/garantex_bot/internal/common"
	"github.com/akelbikhanov/garantex_bot/internal/config"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"time"
)

// DefaultHandler обработчик по-умолчанию
// первая точка входа обработки всех запросов
func DefaultHandler() bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		ctx1, cancel1 := context.WithTimeout(ctx, config.Get().DefaultTimeout)
		defer cancel1()

		deadline, _ := ctx1.Deadline()

		switch {
		case update.Message != nil:
			messageHandler(ctx1, b, update.Message)
			common.LogInfo(fmt.Sprintf(common.InfoUpdateProcessed, update.ID, config.Get().DefaultTimeout-time.Until(deadline)))
		case update.CallbackQuery != nil:
			callbackHandler(ctx1, b, update.CallbackQuery)
			common.LogInfo(fmt.Sprintf(common.InfoUpdateProcessed, update.ID, config.Get().DefaultTimeout-time.Until(deadline)))
		default:
			common.LogInfo(fmt.Sprintf(common.InfoUpdateSkip, update.ID))
		}
	}
}
