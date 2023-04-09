package redisext

import "time"

type Configs struct {
	RedisAddress       string        `env:"REDIS_ADDRESS,required"`
	RedisPassword      string        `env:"REDIS_PASSWORD"`
	RedisDatabase      int           `env:"REDIS_DATABASE"`
	RedisQueueCleaning time.Duration `env:"REDIS_QUEUE_CLEANING" envDefault:"1m"`
	RedisQueuePool     int64         `env:"REDIS_QUEUE_POOL" envDefault:"10"`
	RedisQueueApp      string        `env:"REDIS_QUEUE_APP" envDefault:"app"`
}
