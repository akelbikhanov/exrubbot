package handler

import (
	"context"
	"github.com/akelbikhanov/garantex_bot/internal/common"
	"github.com/akelbikhanov/garantex_bot/internal/service/garantex"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// messageHandler
func messageHandler(ctx context.Context, b *bot.Bot, msg *models.Message) {
	switch msg.Text {
	case common.CommandStart:
		sendTextHandler(ctx, b, msg.From.ID, common.MessageStart, nil)
	case common.CommandHelp:
		sendTextHandler(ctx, b, msg.From.ID, common.MessageHelp, nil)
	case common.CommandPrice:
		sendTextHandler(ctx, b, msg.From.ID, garantex.GetPriceText(), nil)
	case common.CommandRepeat:
	case common.CommandStop:
	default:
	}
}
