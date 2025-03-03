package main

import (
	"context"
	"github.com/akelbikhanov/garantex_bot/internal/service"
	"log"

	"github.com/akelbikhanov/garantex_bot/internal/config"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Создаём бота с обработкой ошибки
	b, err := bot.New(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	// Обработчик команды /start
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		chatID := update.Message.Chat.ID
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Привет! Я буду отправлять тебе курс USDT/RUB.",
		})
	})

	// Получаем курс USDT/RUB перед запуском бота
	buy, sell, err := service.GetUSDTPrice()
	if err != nil {
		log.Fatalf("Ошибка получения курса USDT/RUB: %v", err)
	}
	log.Printf("Курс USDT/RUB: покупка %s ₽, продажа %s ₽", buy, sell)

	// Запускаем бота
	b.Start(context.Background())
}
