package notifier

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/k5sha/lifeEasier/internal/model"
	"log"
	"time"
)

type LinkProvider interface {
	AllNotPosted(ctx context.Context, limit uint64) ([]model.Link, error)
	MarkAsPosted(ctx context.Context, id int64) error
}

type Notifier struct {
	links        LinkProvider
	bot          *tgbotapi.BotAPI
	sendInterval time.Duration
}

func New(
	linkProvider LinkProvider,
	bot *tgbotapi.BotAPI,
	sendInterval time.Duration,
) *Notifier {
	return &Notifier{
		links:        linkProvider,
		bot:          bot,
		sendInterval: sendInterval,
	}
}

func (n *Notifier) Start(ctx context.Context) error {
	ticker := time.NewTicker(n.sendInterval)
	defer ticker.Stop()

	if err := n.SelectAndSendLink(ctx); err != nil {
		log.Printf("[ERROR] %s", err)
		return err
	}

	for {
		select {
		case <-ticker.C:
			if err := n.SelectAndSendLink(ctx); err != nil {
				log.Printf("[ERROR] %s", err)
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// TODO: Yaml config
func (n *Notifier) SelectAndSendLink(ctx context.Context) error {
	links, err := n.links.AllNotPosted(ctx, 3)
	if err != nil {
		return err
	}
	if len(links) == 0 {
		return nil
	}

	for _, link := range links {
		if err := n.sendArticle(link); err != nil {
			return err
		}
		if err := n.links.MarkAsPosted(ctx, link.Id); err != nil {
			return err
		}
	}

	return nil
}

func (n *Notifier) sendArticle(link model.Link) error {
	msgText := "*📢 Time for do it!*\n\n"

	if link.Message != "" {
		msgText += fmt.Sprintf("*📦 Your message:* %s\n\n", link.Message)
	}

	msgText += fmt.Sprintf("*🔗 Link:* [Click](%s)", link.Link)

	msg := tgbotapi.NewMessage(link.ChatId, msgText)
	msg.ParseMode = tgbotapi.ModeMarkdown

	_, err := n.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
