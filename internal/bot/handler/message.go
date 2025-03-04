package handler

import (
	"github.com/akelbikhanov/garantex_bot/internal/common"
	"github.com/akelbikhanov/garantex_bot/internal/service/garantex"
	"github.com/go-telegram/bot/models"
)

// messageHandler
func (h *Handler) messageHandler(msg *models.Message) {
	switch msg.Text {
	case common.CommandStart:
		SendText(h.ctx, h.b, msg.From.ID, common.MessageStart, nil)
	case common.CommandHelp:
		SendText(h.ctx, h.b, msg.From.ID, common.MessageHelp, nil)
	case common.CommandPrice:
		SendText(h.ctx, h.b, msg.From.ID, garantex.GetPriceText(), nil)
	case common.CommandRepeat:
		SendText(h.ctx, h.b, msg.From.ID, common.MessageRepeat, kbRepeat)
	case common.CommandStop:
		if h.Unsubscribe(msg.From.ID) {
			SendText(h.ctx, h.b, msg.From.ID, common.MessageStopYes, nil)
		} else {
			SendText(h.ctx, h.b, msg.From.ID, common.MessageStopNo, nil)
		}
	default:
	}
}
