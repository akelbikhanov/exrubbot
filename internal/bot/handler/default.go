package handler

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/text"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	updateFieldNames []string
	cacheOnce        sync.Once
)

// DefaultHandler обработчик по-умолчанию.
func (h *Handler) DefaultHandler() bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		start := time.Now()
		defer func() {
			elapsed := time.Since(start)
			h.logg.Info(fmt.Sprintf(text.InfoUpdateProcessingTime, update.ID, elapsed))
		}()

		switch {
		case update.Message != nil:
			h.handlerMessage(ctx, b, update.Message)
		case update.CallbackQuery != nil:
			h.handlerCallback(ctx, b, update.CallbackQuery)
		case update.MyChatMember != nil:
			h.handlerMyChatMember(update.MyChatMember)
		default:
			h.logg.Info(fmt.Sprintf(text.InfoUpdateProcessingSkip, getUpdateType(update)))
		}
	}
}

// GetUpdateType возвращает имя первого непустого поля типа *models.Update.
// Это позволяет определить, какой тип апдейта пришёл от Telegram.
func getUpdateType(update *models.Update) string {
	if update == nil {
		return "nil"
	}

	v := reflect.ValueOf(update).Elem()

	// Один раз кешируем имена всех полей структуры Update
	cacheOnce.Do(func() {
		t := v.Type()
		updateFieldNames = make([]string, 0, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			// Берём только указательные поля (все события приходят как *Type)
			if t.Field(i).Type.Kind() == reflect.Ptr {
				updateFieldNames = append(updateFieldNames, t.Field(i).Name)
			}
		}
	})

	// Ищем первое не-nil указательное поле
	for _, fieldName := range updateFieldNames {
		field := v.FieldByName(fieldName)
		if field.IsValid() && !field.IsNil() {
			return fieldName
		}
	}

	return text.InfoUpdateUnknownType
}
