package handler

import (
	"context"
	"github.com/akelbikhanov/garantex_bot/internal/common"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SendText(ctx context.Context, b *bot.Bot, ChatID int64, text string, kb models.ReplyMarkup) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      ChatID,
		Text:        bot.EscapeMarkdown(text),
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: kb,
	})
	if err != nil {
		common.LogError(err)
	}
}

func EditText(ctx context.Context, b *bot.Bot, ChatID int64, MessageID int, text string) {
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    ChatID,
		MessageID: MessageID,
		Text:      bot.EscapeMarkdown(text),
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		common.LogError(err)
	}
}
