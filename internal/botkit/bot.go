package botkit

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"runtime/debug"
	"time"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func New(api *tgbotapi.BotAPI) *Bot {
	return &Bot{api: api}
}

func (b *Bot) Run(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for {
		select {
		case update := <-updates:
			updateCtx, updateCancel := context.WithTimeout(context.Background(), 5*time.Minute)
			b.handleUpdate(updateCtx, update)
			updateCancel()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (b *Bot) handleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("[ERROR] panic recovered: %v\n%s", p, string(debug.Stack()))
		}
	}()
}
