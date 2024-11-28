package telebot

import (
	"time"
)

type Configs struct {
	TelegramStoragePrefix string        `env:"TELEGRAM_STORAGE_PREFIX" envDefault:"telegram"`
	TelegramPullInterval  time.Duration `env:"TELEGRAM_PULL_INTERVAL" envDefault:"30s"`
	TelegramConcurrency   uint8         `env:"TELEGRAM_CONCURRENCY" envDefault:"10"`
}
