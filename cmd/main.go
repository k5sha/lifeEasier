package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"github.com/k5sha/lifeEasier/internal/botkit"
	"github.com/k5sha/lifeEasier/internal/config"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	botAPI, err := tgbotapi.NewBotAPI(config.Get().TelegramBotToken)
	if err != nil {
		log.Printf("[ERROR] failed to create botAPI: %v", err)
		return
	}

	db, err := sqlx.Connect("postgres", config.Get().DatabaseDSN)
	if err != nil {
		log.Printf("[ERROR] failed to connect to db: %v", err)
		return
	}
	defer db.Close()

	easyBot := botkit.New(botAPI)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := easyBot.Run(ctx); err != nil {
		log.Printf("[ERROR] failed to run botkit: %v", err)
	}

}
