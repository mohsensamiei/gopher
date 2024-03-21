package response

import (
	"github.com/mohsensamiei/gopher/v2/pkg/cache"
	"github.com/mohsensamiei/gopher/v2/pkg/redisext"
)

func New(configs Configs, rdb *redisext.Client) (*cache.Client, error) {
	return cache.NewClient(
		cache.ClientWithAdapter(NewRedisAdapter(rdb)),
		cache.ClientWithTTL(configs.ResponseCacheDuration),
		cache.ClientWithRefreshKey("fresh"),
	)
}
