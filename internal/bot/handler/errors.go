package handler

import (
	"errors"
	"fmt"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot"
)

// ErrorHandler функция обработки ошибок, получаемых при запросе данных
// здесь перехватываются ошибки, связанные с фоновым взаимодействием бота и сервиса
func (h *Handler) ErrorHandler() bot.ErrorsHandler {
	return func(e error) {
		h.processRequestResult("bot.ErrorHandler()", e, 0, "")
	}
}

// processRequestResult проверяет ошибку, полученную от telegram при выполнении запросов.
func (h *Handler) processRequestResult(source string, e error, chatID int64, data string) {
	msg := fmt.Sprintf("%s, chat_id:%d, data:'%s'", source, chatID, data)
	if e != nil {
		switch {
		case errors.Is(e, bot.ErrorForbidden):
			if h.noty.Unsubscribe(chatID) {
				h.logg.Warn(text.WarnForbiddenSubscribe, e)
			} else {
				h.logg.Warn(text.WarnForbiddenUnsubscribe, e)
			}
		case errors.Is(e, bot.ErrorTooManyRequests):
			h.logg.Warn(fmt.Sprintf(text.WarnRequestDelay, h.retryTime), e)
			time.Sleep(h.retryTime)
		default:
			h.terminateWith(e)
		}
	} else {
		h.logg.Info(msg)
	}
}

// terminateWith завершает работу бота.
func (h *Handler) terminateWith(e error) {
	h.logg.Error(text.ErrBotTerminate, e, 4)
	h.TerminateReason = e
	h.cancel()
}
