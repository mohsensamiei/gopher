package response

import (
	"context"
	"fmt"
	"github.com/mohsensamiei/gopher/v2/pkg/redisext"
	"time"
)

func NewRedisAdapter(rdb *redisext.Client) *RedisAdapter {
	return &RedisAdapter{
		rdb: rdb,
	}
}

type RedisAdapter struct {
	rdb *redisext.Client
}

func (a RedisAdapter) Get(ctx context.Context, key uint64) ([]byte, bool) {
	data, err := a.rdb.DB.Get(ctx, fmt.Sprint("responses:", key)).Bytes()
	return data, err == nil
}

func (a RedisAdapter) Set(ctx context.Context, key uint64, response []byte, expiration time.Time) {
	_ = a.rdb.DB.Set(ctx, fmt.Sprint("responses:", key), response, expiration.Sub(time.Now())).Err()
}

func (a RedisAdapter) Release(ctx context.Context, key uint64) {
	_ = a.rdb.DB.Del(ctx, fmt.Sprint("responses:", key)).Err()
}
