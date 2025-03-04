package handler

import (
	"context"
	"errors"
	"github.com/akelbikhanov/garantex_bot/internal/common"
	"github.com/go-telegram/bot"
	"time"
)

// ErrorsHandler функция обработки ошибок, получаемых при запросе данных
// здесь перехватываются ошибки, связанные с фоновым взаимодействием бота и сервиса
func ErrorsHandler(cancelFunc context.CancelFunc, duration time.Duration) bot.ErrorsHandler {
	return func(e error) {
		common.LogError(e)

		if errors.Is(e, bot.ErrorTooManyRequests) {
			time.Sleep(duration)
		} else {
			cancelFunc()
		}
	}
}
