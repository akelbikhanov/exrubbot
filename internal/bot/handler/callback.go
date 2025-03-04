package handler

import (
	"context"
	"github.com/akelbikhanov/garantex_bot/internal/common"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// callbackHandler
func callbackHandler(ctx context.Context, b *bot.Bot, callbackQuery *models.CallbackQuery) {
	defer answerCallbackQuery(ctx, b, callbackQuery.ID)

}

// answering callback query first to let Telegram know that we received the callback query,
// and we're handling it. Otherwise, Telegram might retry sending the update repetitively
// as it thinks the callback query doesn't reach to our application. learn more by
// reading the footnote of the https://core.telegram.org/bots/api#callbackquery type.
func answerCallbackQuery(ctx context.Context, b *bot.Bot, callbackID string) {
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callbackID,
		Text:            "Готово",
		ShowAlert:       false,
	})
	if err != nil {
		common.LogError(err)
	}
}
