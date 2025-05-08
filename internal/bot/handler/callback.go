package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// handlerCallback обработчик update.CallbackQuery.
func (h *Handler) handlerCallback(ctx context.Context, b *bot.Bot, cb *models.CallbackQuery) {
	defer h.answerCallbackQuery(ctx, b, cb.From.ID, cb.ID)

	parts := strings.Split(cb.Data, text.CallbackSeparator)
	switch parts[0] {
	case text.CallbackQuotePrefix:
		h.handleQuoteCallback(ctx, b, cb, parts)
	default:
		msg := fmt.Sprintf(text.MessageCallbackIncorrect, cb.Data)
		h.editText(ctx, b, cb.From.ID, cb.Message.Message.ID, msg, nil)
	}
}
