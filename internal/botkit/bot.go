package botkit

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/k5sha/lifeEasier/internal/bot"
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

func (b *Bot) Run(ctx context.Context, linkStorage bot.LinkStorage) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for {
		select {
		case update := <-updates:
			updateCtx, updateCancel := context.WithTimeout(context.Background(), 5*time.Minute)
			b.handleUpdate(updateCtx, update, linkStorage)
			updateCancel()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (b *Bot) handleUpdate(ctx context.Context, update tgbotapi.Update, linkStorage bot.LinkStorage) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("[ERROR] panic recovered: %v\n%s", p, string(debug.Stack()))
		}
	}()

	if update.Message == nil {
		return
	}

	// TODO: add handler to command
	if update.Message.IsCommand() {
		return
	}

	err := bot.AddLinkHandler(ctx, b.api, update, linkStorage)
	if err != nil {
		log.Printf("[%s:%d] %s", update.Message.From.UserName, update.Message.Chat.ID, err)
	}
}
