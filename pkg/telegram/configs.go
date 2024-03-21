package telegram

import "time"

type Configs struct {
	TelegramToken            string           `env:"TELEGRAM_TOKEN,required"`
	TelegramTimeout          time.Duration    `env:"TELEGRAM_TIMOUT" envDefault:"0"`
	TelegramDefaultParseMode FormattingOption `env:"TELEGRAM_DEFAULT_PARSE_MODE" envDefault:"HTML"`
}
