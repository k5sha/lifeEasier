package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
	"log"
	"sync"
)

type Config struct {
	TelegramBotToken string `yaml:"telegram_bot_token" env:"TELEGRAM_BOT_TOKEN" required:"true"`
	DatabaseDSN      string `yaml:"database_dsn" env:"DATABASE_DSN" default:"postgres://postgres:postgres@localhost:5432/life_easier_db?sslmode=disable"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		loader := aconfig.LoaderFor(&cfg, aconfig.Config{
			EnvPrefix: "NFB",
			Files:     []string{"./config.yaml", "./config.local.yaml", "$HOME/.config/lifeEasier/config.yaml"},
			FileDecoders: map[string]aconfig.FileDecoder{
				".yaml": aconfigyaml.New(),
			},
		})

		if err := loader.Load(); err != nil {
			log.Printf("[ERROR] failed to load config: %v", err)
		}
	})

	return cfg
}
