package telebot

import (
	"time"
)

type Configs struct {
	ExternalURL           string        `env:"EXTERNAL_URL"`
	TelegramSecretToken   string        `env:"TELEGRAM_SECRET_TOKEN"`
	TelegramStoragePrefix string        `env:"TELEGRAM_STORAGE_PREFIX" envDefault:"telegram"`
	TelegramPullInterval  time.Duration `env:"TELEGRAM_PULL_INTERVAL" envDefault:"30s"`
	TelegramConcurrency   uint8         `env:"TELEGRAM_CONCURRENCY" envDefault:"10"`
}
