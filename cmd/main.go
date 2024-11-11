package main

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"github.com/k5sha/lifeEasier/internal/botkit"
	"github.com/k5sha/lifeEasier/internal/config"
	"github.com/k5sha/lifeEasier/internal/notifier"
	"github.com/k5sha/lifeEasier/internal/storage"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Printf("[INFO] Bot starting...")
	botAPI, err := tgbotapi.NewBotAPI(config.Get().TelegramBotToken)
	if err != nil {
		log.Printf("[ERROR] failed to create botAPI: %v", err)
		return
	}
	log.Printf("[INFO] Authorized on account %s", botAPI.Self.UserName)

	db, err := sqlx.Connect("postgres", config.Get().DatabaseDSN)
	if err != nil {
		log.Printf("[ERROR] failed to connect to db: %v", err)
		return
	}
	defer db.Close()
	log.Printf("[INFO] Connected to database successful")

	easyBot := botkit.New(botAPI)

	linkStorage := storage.NewLinkStorage(db)

	notifier := notifier.New(
		linkStorage,
		botAPI,
		config.Get().SendInterval,
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func(ctx context.Context) {
		if err := notifier.Start(ctx); err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Printf("[ERROR] failed to run notifier: %v", err)
				return
			}

			log.Printf("[INFO] notifier stopped")
		}
	}(ctx)

	if err := easyBot.Run(ctx, linkStorage); err != nil {
		log.Printf("[ERROR] failed to run botkit: %v", err)
	}

}
