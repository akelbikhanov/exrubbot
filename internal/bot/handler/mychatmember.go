package handler

import (
	"fmt"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot/models"
)

// handleMyChatMember обрабатывает события изменения статуса бота в чате.
func (h *Handler) handlerMyChatMember(cm *models.ChatMemberUpdated) {
	if cm.NewChatMember.Type == models.ChatMemberTypeBanned &&
		h.noty.Unsubscribe(cm.Chat.ID) {
		h.logg.Info(fmt.Sprintf(text.InfoChatMemberBanned, cm.Chat.ID))
	}
}
