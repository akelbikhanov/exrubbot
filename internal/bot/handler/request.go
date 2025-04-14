package handler

import (
	"context"
	"errors"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/logger"
	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// ErrorHandler функция обработки ошибок, получаемых при запросе данных
// здесь перехватываются ошибки, связанные с фоновым взаимодействием бота и сервиса
func (h *Handler) ErrorHandler() bot.ErrorsHandler {
	return func(e error) {
		logger.Error(text.ErrRequestUpdates, e)
		h.isFatalError(e)
	}
}

// sendText отправляет тестовое сообщение с клавиатурой (опционально).
func (h *Handler) sendText(ctx context.Context, b *bot.Bot, chatID int64, text string, kb models.ReplyMarkup) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        bot.EscapeMarkdown(text),
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: kb,
	})
	if err != nil {
		logger.Error("", err)
		h.isFatalError(err)
	}
}

// editText меняет текст в ранее отправленном сообщении бота.
func (h *Handler) editText(ctx context.Context, b *bot.Bot, chatID int64, messageID int, text string) {
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    chatID,
		MessageID: messageID,
		Text:      bot.EscapeMarkdown(text),
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		logger.Error("", err)
		h.isFatalError(err)
	}
}

// answerCallbackQuery сообщает telegram, что бот получил запрос обратного вызова (callbackQuery).
func (h *Handler) answerCallbackQuery(ctx context.Context, b *bot.Bot, callbackID string) {
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callbackID,
		//Text:            "Готово",
		//ShowAlert:       false,
	})
	if err != nil {
		logger.Error("", err)
		h.isFatalError(err)
	}
}

// isFatalError проверяет ошибку, полученную от telegram при выполнении запросов.
// Если она отличается от ErrorTooManyRequests, завершает работу бота/приложения.
func (h *Handler) isFatalError(err error) {
	if err == nil {
		return
	}

	switch {
	case errors.Is(err, bot.ErrorTooManyRequests):
		time.Sleep(h.timeout)
	default:
		h.CancelError = err
		h.cancel()
	}
}
