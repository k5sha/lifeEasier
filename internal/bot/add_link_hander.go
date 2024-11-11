package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/k5sha/lifeEasier/internal/model"
	"log"
	"regexp"
	"strings"
)

type LinkStorage interface {
	Add(ctx context.Context, link model.Link) (int64, error)
}

func parseMessage(msg tgbotapi.Message) (model.Link, string, error) {
	re := regexp.MustCompile(`https?://[^\s]+`)

	matches := re.FindString(msg.Text)
	if matches == "" {
		return model.Link{}, "", fmt.Errorf("no URL found in the message")
	}
	modifiedText := strings.Replace(msg.Text, matches, "", -1)

	link := model.Link{
		Message: modifiedText,
		Link:    matches,
		ChatId:  msg.Chat.ID,
	}

	return link, matches, nil
}

func AddLinkHandler(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update, linkStorage LinkStorage) error {
	if update.Message == nil {
		return fmt.Errorf("no message found in the update")
	}

	linkModel, extractedURL, err := parseMessage(*update.Message)
	if err != nil || extractedURL == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Link not find in message"))
		_, err = bot.Send(msg)
		if err != nil {
			return fmt.Errorf("failed to send error message: %v", err)
		}
		return fmt.Errorf("failed to parse message: %v", err)
	}

	linkID, err := linkStorage.Add(ctx, linkModel)
	if err != nil {
		return fmt.Errorf("failed to store the linkModel: %v", err)
	}

	log.Printf("[DEBUG] Successful add new linkModel with ID: %d ", linkID)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Link added successfully ✅ Wait, I'll remind you later (1d-7d)")
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send confirmation message: %v", err)
	}

	return nil
}
