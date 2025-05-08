package handler

import (
	"context"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// handlerMessage обработчик update.Message.
func (h *Handler) handlerMessage(ctx context.Context, b *bot.Bot, msg *models.Message) {
	switch msg.Text {
	case text.CommandStart:
		h.sendText(ctx, b, msg.From.ID, text.MessageStart, nil)
	case text.CommandQuote:
		h.sendFeeds(ctx, b, msg.From.ID)
	case text.CommandStop:
		if h.noty.Unsubscribe(msg.From.ID) {
			h.sendText(ctx, b, msg.From.ID, text.MessageStopYes, nil)
		} else {
			h.sendText(ctx, b, msg.From.ID, text.MessageStopNo, nil)
		}
	default:
		h.sendText(ctx, b, msg.From.ID, text.MessageCommandUnknown, nil)
	}
}
