package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/akelbikhanov/exrubbot/internal/text"
	tg "github.com/go-telegram/bot"
)

// RestoreSubscriptions восстанавливает подписки из файла.
func (b *Bot) RestoreSubscriptions(ctx context.Context, bot *tg.Bot) {
	data, err := b.storage.Load()
	if err != nil {
		b.logger.Error(text.LogErrSubsLoad, text.LogFieldError, err)
		return
	}

	restored := 0
	skipped := 0

	for _, item := range data {
		send, err1 := b.createSendFunc(ctx, bot, item, 0)
		if err1 != nil {
			b.logger.Error(text.LogErrSubsRestore,
				text.LogFieldChatID, item.ChatID,
				text.LogFieldFeedID, item.FeedID,
				text.LogFieldIntervalSec, item.IntervalSec,
				text.LogFieldError, err1)
			skipped++
			continue
		}
		if err2 := b.scheduler.Subscribe(ctx, item, send); err2 != nil {
			b.logger.Error(text.LogErrSubsRestore,
				text.LogFieldChatID, item.ChatID,
				text.LogFieldFeedID, item.FeedID,
				text.LogFieldIntervalSec, item.IntervalSec,
				text.LogFieldError, err1)
			skipped++
			continue
		}
		restored++
	}

	b.logger.Info(fmt.Sprintf(text.LogInfoSubsRestored, restored, skipped))
}

func (b *Bot) SaveSubscriptions(ctx context.Context) {
	const saveInterval = 5 * time.Minute
	ticker := time.NewTicker(saveInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			b.saveSubscriptions()
			return
		case <-ticker.C:
			b.saveSubscriptions()
		}
	}
}

func (b *Bot) saveSubscriptions() {
	subs := b.scheduler.GetSubscriptions()
	if err := b.storage.Save(subs); err != nil {
		b.logger.Error(text.LogErrSubsSave, text.LogFieldError, err)
		return
	}

	b.logger.Info(fmt.Sprintf(text.LogInfoSubsSaved, len(subs)))
}
