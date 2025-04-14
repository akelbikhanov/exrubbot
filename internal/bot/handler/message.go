package handler

import (
	"context"
	"github.com/akelbikhanov/exrubbot/internal/datafeed"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// messageHandler
func (h *Handler) messageHandler(ctx context.Context, b *bot.Bot, msg *models.Message) {
	if msg == nil {
		return
	}

	switch msg.Text {
	case text.CommandStart:
		h.sendText(ctx, b, msg.From.ID, text.MessageStart, nil)
	case text.CommandHelp:
		h.sendText(ctx, b, msg.From.ID, text.MessageHelp, nil)
	case text.CommandPrice:
		h.sendQuote(ctx, b, msg.From.ID, datafeed.Grinex, 0)
	case text.CommandRepeat:
		h.sendText(ctx, b, msg.From.ID, text.MessageRepeat, kbRepeat)
	case text.CommandStop:
		if h.n.Unsubscribe(msg.From.ID) {
			h.sendText(ctx, b, msg.From.ID, text.MessageStopYes, nil)
		} else {
			h.sendText(ctx, b, msg.From.ID, text.MessageStopNo, nil)
		}
	default:
		h.sendText(ctx, b, msg.From.ID, text.MessageUnknownCommand, nil)
	}
}
