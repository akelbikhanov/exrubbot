package bot

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/entity"
	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/akelbikhanov/exrubbot/internal/version"
	tg "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// HandleUpdate обрабатывает входящие обновления от Telegram.
func (b *Bot) HandleUpdate(ctx context.Context, bot *tg.Bot, update *models.Update) {
	start := time.Now()
	defer b.logger.Debug(fmt.Sprintf(text.LogDebugUpdateProcessed, update.ID, time.Since(start)))

	switch {
	case update.Message != nil:
		b.handleMessage(ctx, bot, update.ID, update.Message)
	case update.CallbackQuery != nil:
		b.handleCallback(ctx, bot, update.ID, update.CallbackQuery)
	case update.MyChatMember != nil:
		b.handleMyChatMember(update.MyChatMember)
	default:
		// Игнорируем неподдерживаемые типы обновлений
		b.logger.Debug(fmt.Sprintf(text.LogDebugUpdateUnknownType, update.ID))
	}
}

// handleMessage обрабатывает текстовые сообщения.
func (b *Bot) handleMessage(ctx context.Context, bot *tg.Bot, updateID int64, msg *models.Message) {
	if msg.From == nil {
		return
	}

	switch msg.Text {
	case text.CmdStart:
		b.handleCommandStart(ctx, bot, msg.From.ID)
	case text.CmdQuote:
		b.handleCommandQuote(ctx, bot, msg.From.ID)
	case text.CmdStop:
		b.handleCommandStop(ctx, bot, msg.From.ID)
	case text.CmdVersion:
		b.handleCommandVersion(ctx, bot, msg.From.ID)
	default:
		b.handleCommandUnknown(ctx, bot, updateID, msg.From.ID)
	}
}

// handleCommandStart обрабатывает команду /start.
func (b *Bot) handleCommandStart(ctx context.Context, bot *tg.Bot, chatID int64) {
	b.sendMessage(ctx, bot, chatID, text.MsgWelcome, nil)
}

// handleCommandQuote обрабатывает команду /quote.
func (b *Bot) handleCommandQuote(ctx context.Context, bot *tg.Bot, chatID int64) {
	b.handleShowFeeds(ctx, bot, chatID, 0)
}

// handleShowFeeds показывает выбор источников котировок.
// Если messageID == 0, отправляем как новое сообщение; иначе - изменяем существующее.
func (b *Bot) handleShowFeeds(ctx context.Context, bot *tg.Bot, chatID int64, messageID int) {
	kb := feedsKeyboard(b.feeds)
	if messageID == 0 {
		if kb == nil {
			b.sendMessage(ctx, bot, chatID, text.MsgProviderError, nil)
		} else {
			b.sendMessage(ctx, bot, chatID, text.MsgSelectProvider, kb)
		}
	} else {
		if kb == nil {
			b.editMessage(ctx, bot, chatID, messageID, text.MsgProviderError, nil)
		} else {
			b.editMessage(ctx, bot, chatID, messageID, text.MsgSelectProvider, kb)
		}
	}
}

// handleCommandStop обрабатывает команду /stop.
func (b *Bot) handleCommandStop(ctx context.Context, bot *tg.Bot, chatID int64) {
	if b.scheduler.Unsubscribe(chatID) {
		b.logger.Info(text.LogInfoSubsCancelled, text.LogFieldChatID, chatID)
		b.sendMessage(ctx, bot, chatID, text.MsgUnsubscribeOK, nil)
	} else {
		b.sendMessage(ctx, bot, chatID, text.MsgUnsubscribeNone, nil)
	}
}

// handleCommandVersion обрабатывает команду /version.
func (b *Bot) handleCommandVersion(ctx context.Context, bot *tg.Bot, chatID int64) {
	b.sendMessage(ctx, bot, chatID, version.GetVersion(), nil)
}

// handleCommandUnknown обрабатывает неизвестные команды.
func (b *Bot) handleCommandUnknown(ctx context.Context, bot *tg.Bot, updateID, chatID int64) {
	b.logger.Warn(fmt.Sprintf(text.LogWarnUnknownCommand, updateID))
	b.sendMessage(ctx, bot, chatID, text.MsgUnknownCommand, nil)
}

// handleCallback обрабатывает нажатия на инлайн-кнопки.
func (b *Bot) handleCallback(ctx context.Context, bot *tg.Bot, updateID int64, cb *models.CallbackQuery) {
	// Всегда отвечаем на callback
	defer b.answerCallback(ctx, bot, cb.ID)

	// Парсим данные callback: "quote:feedID:action"
	parts := strings.Split(cb.Data, text.CallbackSeparator)
	switch parts[0] {
	case text.CallbackPrefixQuote:
		b.handleQuoteCallback(ctx, bot, updateID, cb, parts)
	default:
		b.logger.Warn(fmt.Sprintf(text.LogWarnUnknownCommand, updateID))
		msg := fmt.Sprintf(text.MsgCallbackError, cb.Data)
		b.editMessage(ctx, bot, cb.From.ID, cb.Message.Message.ID, msg, nil)
	}
}

// handleQuoteCallback обрабатывает callback-команды /quote.
// Возможные форматы:
// - quote:<feedID>           		→ показать выбор интервала
// - quote:<feedID>:0         		→ разовая котировка
// - quote:<feedID>:<positive int>	→ подписка
// - quote:<feedID>:<negative int>	→ вернуться к выбору источника
func (b *Bot) handleQuoteCallback(ctx context.Context, bot *tg.Bot, updateID int64, cb *models.CallbackQuery, parts []string) {
	switch len(parts) {
	case 2:
		// quote:feedID - показать выбор интервала
		b.handleShowIntervals(ctx, bot, cb, parts[1])
	case 3:
		// quote:feedID:interval
		b.handleSelectInterval(ctx, bot, updateID, cb, parts[1], parts[2])
	default:
		// quote:???
		b.logger.Warn(fmt.Sprintf(text.LogWarnUnknownCommand, updateID))
		msg := fmt.Sprintf(text.MsgCallbackError, cb.Data)
		b.editMessage(ctx, bot, cb.From.ID, cb.Message.Message.ID, msg, nil)
	}
}

// handleShowIntervals показывает выбор интервала для подписки.
func (b *Bot) handleShowIntervals(ctx context.Context, bot *tg.Bot, cb *models.CallbackQuery, feedID string) {
	kb := intervalsKeyboard(feedID)
	b.editMessage(ctx, bot, cb.From.ID, cb.Message.Message.ID, text.MsgSelectInterval, kb)
}

// handleSelectInterval обрабатывает выбор интервала или специальные действия.
func (b *Bot) handleSelectInterval(ctx context.Context, bot *tg.Bot, updateID int64, cb *models.CallbackQuery, feedID, interval string) {
	// Парсим интервал
	seconds, err := strconv.Atoi(interval)
	if err != nil {
		b.logger.Warn(fmt.Sprintf(text.LogWarnUnknownCommand, updateID))
		msg := fmt.Sprintf(text.MsgCallbackError, cb.Data)
		b.editMessage(ctx, bot, cb.From.ID, cb.Message.Message.ID, msg, nil)
		return
	}

	if seconds < 0 {
		b.handleShowFeeds(ctx, bot, cb.From.ID, cb.Message.Message.ID)
		return
	}

	sub := entity.Subscription{
		ChatID:      cb.From.ID,
		FeedID:      feedID,
		IntervalSec: seconds,
		CreatedAt:   time.Now(),
	}

	// Создаём функцию, которая обновляет текущее сообщение.
	updateMessageFunc, err1 := b.createSendFunc(ctx, bot, sub, cb.Message.Message.ID)
	if err1 != nil {
		b.logger.Error(text.LogErrSubsCreate,
			text.LogFieldChatID, sub.ChatID,
			text.LogFieldFeedID, sub.FeedID,
			text.LogFieldIntervalSec, sub.IntervalSec,
			text.LogFieldError, err1)
		msg := fmt.Sprintf(text.MsgCallbackError, cb.Data)
		b.editMessage(ctx, bot, cb.From.ID, cb.Message.Message.ID, msg, nil)
		return
	}

	// Если интервал больше нуля, создаём подписку.
	if seconds > 0 {
		send, err2 := b.createSendFunc(ctx, bot, sub, 0)
		if err2 != nil {
			b.logger.Error(text.LogErrSubsCreate,
				text.LogFieldChatID, sub.ChatID,
				text.LogFieldFeedID, sub.FeedID,
				text.LogFieldIntervalSec, sub.IntervalSec,
				text.LogFieldError, err2)
			msg := fmt.Sprintf(text.MsgCallbackError, cb.Data)
			b.editMessage(ctx, bot, cb.From.ID, cb.Message.Message.ID, msg, nil)
			return
		}
		if err3 := b.scheduler.Subscribe(ctx, sub, send); err3 != nil {
			b.logger.Error(text.LogErrSubsCreate,
				text.LogFieldChatID, sub.ChatID,
				text.LogFieldFeedID, sub.FeedID,
				text.LogFieldIntervalSec, sub.IntervalSec,
				text.LogFieldError, err3)
			msg := fmt.Sprintf(text.MsgCallbackError, cb.Data)
			b.editMessage(ctx, bot, cb.From.ID, cb.Message.Message.ID, msg, nil)
			return
		}

		// Логируем успешное создание подписки
		b.logger.Info(text.LogInfoSubsCreated,
			text.LogFieldChatID, sub.ChatID,
			text.LogFieldFeedID, sub.FeedID,
			text.LogFieldIntervalSec, sub.IntervalSec)
	}

	// Только после успешного создания подписки
	// вызываем функцию обновления текущего сообщения.
	updateMessageFunc()
}

// handleMyChatMember обрабатывает изменения статуса бота в чате.
func (b *Bot) handleMyChatMember(update *models.ChatMemberUpdated) {
	// Если пользователь заблокировал у себя бота, прекращаем рассылку.
	if update.NewChatMember.Type == models.ChatMemberTypeBanned {
		b.logger.Info(text.LogInfoUserBlocked, text.LogFieldChatID, update.Chat.ID)
		b.scheduler.Unsubscribe(update.Chat.ID)
	}
}

// HandleError функция-декоратор, которая получает ошибки,
// возникающие в процессе фоновой работы клиента Telegram,
// и передаёт их дальше, в обработчик.
func (b *Bot) HandleError(err error) {
	b.handleError(err, 0)
}

// handleError обрабатывает все ошибки взаимодействия с API Telegram.
func (b *Bot) handleError(err error, chatID int64) {
	if err == nil {
		return
	}

	switch {
	case errors.Is(err, tg.ErrorForbidden):
		// Пользователь заблокировал бота.
		b.logger.Warn(text.LogWarnAccessDenied, text.LogFieldError, err)
		if chatID > 0 {
			b.scheduler.Unsubscribe(chatID)
		}

	case errors.Is(err, tg.ErrorTooManyRequests):
		// Превышен лимит частоты отправки, вводим задержку.
		b.logger.Warn(fmt.Sprintf(text.LogWarnRateLimit, b.retryDelay), text.LogFieldError, err)
		time.Sleep(b.retryDelay)

	default:
		// Ошибки, на которые не понятно, как реагировать.
		b.logger.Error(text.LogErrCritical, text.LogFieldChatID, chatID, text.LogFieldError, err)
		b.StopError = err
		b.stop()
	}
}
