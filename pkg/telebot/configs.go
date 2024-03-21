package telebot

import (
	"time"
)

type Configs struct {
	TelegramPullInterval time.Duration `env:"TELEGRAM_PULL_INTERVAL" envDefault:"30s"`
	TelegramConcurrency  uint8         `env:"TELEGRAM_CONCURRENCY" envDefault:"10"`
}
