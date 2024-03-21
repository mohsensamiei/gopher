package response

import "time"

type Configs struct {
	ResponseCacheDuration time.Duration `env:"RESPONSE_CACHE_DURATION" envDefault:"5m"`
}
