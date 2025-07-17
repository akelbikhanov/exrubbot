package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/akelbikhanov/exrubbot/internal/entity"
	"github.com/akelbikhanov/exrubbot/internal/text"
	tg "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// sendMessage отправляет текстовое сообщение.
func (b *Bot) sendMessage(ctx context.Context, bot *tg.Bot, chatID int64, msg string, kb models.ReplyMarkup) {
	_, err := bot.SendMessage(ctx, &tg.SendMessageParams{
		ChatID:      chatID,
		Text:        tg.EscapeMarkdown(msg),
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: kb,
	})

	b.logger.Debug(text.LogDebugSendMessage,
		text.LogFieldChatID, chatID,
		text.LogFieldMessageText, msg)
	b.handleError(err, chatID)
}

// editMessage редактирует/обновляет существующее сообщение.
func (b *Bot) editMessage(ctx context.Context, bot *tg.Bot, chatID int64, messageID int, msg string, kb models.ReplyMarkup) {
	_, err := bot.EditMessageText(ctx, &tg.EditMessageTextParams{
		ChatID:      chatID,
		MessageID:   messageID,
		Text:        tg.EscapeMarkdown(msg),
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: kb,
	})

	b.logger.Debug(text.LogDebugEditMessage,
		text.LogFieldChatID, chatID,
		text.LogFieldMessageID, messageID,
		text.LogFieldMessageText, msg)
	b.handleError(err, chatID)
}

// answerCallback отвечает на callback query.
func (b *Bot) answerCallback(ctx context.Context, bot *tg.Bot, callbackID string) {
	_, err := bot.AnswerCallbackQuery(ctx, &tg.AnswerCallbackQueryParams{
		CallbackQueryID: callbackID,
		// Text:         "Готово",
		// ShowAlert:    false,
	})

	b.logger.Debug(text.LogDebugAnswerCallback, text.LogFieldCallbackID, callbackID)
	b.handleError(err, 0)
}

// createSendFunc создаёт функцию-замыкание для отправки котировки.
// messageID > 0 - editMessage, 0 - sendMessage.
// Interval > 0 - добавляем в сообщение строку с таймером.
func (b *Bot) createSendFunc(ctx context.Context, bot *tg.Bot, sub entity.Subscription, messageID int) (func(), error) {
	// проверяем входные переменные
	feed, ok := b.feeds[sub.FeedID]
	if !ok {
		return nil, fmt.Errorf(text.LogErrFeedNotFound, sub.FeedID)
	}
	if sub.IntervalSec < 0 {
		return nil, fmt.Errorf(text.LogErrInvalidInterval, sub.IntervalSec)
	}

	return func() {
		// строим ответ с котировкой
		var msg strings.Builder
		//msg.WriteString(text.QuoteHeader)
		msg.WriteString(feed.Name())

		quote, err := feed.GetQuote(ctx)
		if err != nil {
			msg.WriteString(text.QuoteError)
			b.logger.Error(text.LogErrQuoteFetch,
				text.LogFieldChatID, sub.ChatID,
				text.LogFieldFeedID, sub.FeedID,
				text.LogFieldError, err)
		} else {
			msg.WriteString(formatQuote(quote))
		}

		// добавляем московское время
		msg.WriteString("\n\n")
		msg.WriteString(formatTimeMSK())

		// добавляем подпись "/stop ⏱..."
		if sub.IntervalSec != 0 {
			msg.WriteString(fmt.Sprintf(text.QuoteStopHint, formatInterval(sub.IntervalSec)))
		}

		// отправляем ответ
		if messageID > 0 {
			b.editMessage(ctx, bot, sub.ChatID, messageID, msg.String(), nil)
		} else {
			b.sendMessage(ctx, bot, sub.ChatID, msg.String(), nil)
		}
	}, nil
}
