package redisext

import "time"

type Configs struct {
	RedisAddress       string        `env:"REDIS_ADDRESS,required"`
	RedisUsername      string        `env:"REDIS_USERNAME"`
	RedisPassword      string        `env:"REDIS_PASSWORD"`
	RedisDatabase      int           `env:"REDIS_DATABASE"`
	RedisTLSRequired   bool          `env:"REDIS_TLS_REQUIRED"`
	RedisQueueCleaning time.Duration `env:"REDIS_QUEUE_CLEANING" envDefault:"1m"`
	RedisQueuePool     int64         `env:"REDIS_QUEUE_POOL" envDefault:"10"`
	RedisQueueApp      string        `env:"REDIS_QUEUE_APP" envDefault:"app"`
}
