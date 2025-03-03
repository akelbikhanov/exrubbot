package main

import (
	"context"
	"log"

	"github.com/azret/garantex_bot/internal/config"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Создаём бота
	b := bot.New(cfg.TelegramToken)

	// Обработчик команды /start
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		chatID := update.Message.Chat.ID
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Привет! Я буду отправлять тебе курс USDT/RUB.",
		})
	})

	// Запускаем бота
	b.Start(context.Background())
}
