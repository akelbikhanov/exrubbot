package handler

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// sendText отправляет тестовое сообщение с клавиатурой (опционально).
func (h *Handler) sendText(ctx context.Context, b *bot.Bot, chatID int64, text string, kb models.ReplyMarkup) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        bot.EscapeMarkdown(text),
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: kb,
	})

	h.processRequestResult("bot.SendMessage()", err, chatID, text)
}

// editText меняет текст в ранее отправленном сообщении бота.
func (h *Handler) editText(ctx context.Context, b *bot.Bot, chatID int64, messageID int, text string, kb models.ReplyMarkup) {
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      chatID,
		MessageID:   messageID,
		Text:        bot.EscapeMarkdown(text),
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: kb,
	})

	h.processRequestResult("bot.EditMessageText()", err, chatID, text)
}

// answerCallbackQuery сообщает telegram, что бот получил запрос обратного вызова (callbackQuery).
func (h *Handler) answerCallbackQuery(ctx context.Context, b *bot.Bot, chatID int64, callbackID string) {
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callbackID,
		//Text:            "Готово",
		//ShowAlert:       false,
	})

	h.processRequestResult("bot.AnswerCallbackQuery()", err, chatID, callbackID)
}
