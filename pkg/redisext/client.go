package redisext

import (
	"context"
	"crypto/tls"
	"encoding/json"
	goerrors "errors"
	"fmt"
	"github.com/adjust/rmq/v5"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis_rate/v10"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"time"
)

func Connect(ctx context.Context, configs Configs) (*Client, error) {
	cli := &Client{
		Configs: configs,
		errChan: make(chan error),
	}
	{
		opts := &redis.Options{
			Addr:     configs.RedisAddress,
			Username: configs.RedisUsername,
			Password: configs.RedisPassword,
			DB:       configs.RedisDatabase,
		}
		if configs.RedisTLSRequired {
			opts.TLSConfig = &tls.Config{}
		}
		cli.DB = redis.NewClient(opts)
		if err := cli.DB.Ping(ctx).Err(); err != nil {
			return nil, err
		}
	}
	{
		var err error
		cli.Queue, err = rmq.OpenConnectionWithRedisClient(configs.RedisQueueApp, cli.DB, cli.errChan)
		if err != nil {
			return nil, err
		}
		go cli.cleaningQueue()
		go cli.loggingQueue()
	}
	{
		cli.Atomic = redislock.New(cli.DB)
	}
	{
		cli.Lock = redsync.New(goredis.NewPool(cli.DB))
	}
	{
		cli.Limiter = redis_rate.NewLimiter(cli.DB)
	}
	return cli, nil
}

type Client struct {
	Configs
	DB      *redis.Client
	Queue   rmq.Connection
	Lock    *redsync.Redsync
	Limiter *redis_rate.Limiter
	errChan chan error
	Atomic  *redislock.Client
}

func (c *Client) Close() error {
	return c.DB.Close()
}

func (c *Client) cleaningQueue() {
	cleaner := rmq.NewCleaner(c.Queue)
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		if _, err := cleaner.Clean(); err != nil && !goerrors.Is(err, rmq.ErrorNotFound) {
			c.errChan <- err
		}
	}
}

func (c *Client) loggingQueue() {
	for err := range c.errChan {
		log.WithError(err).Error("redis queue error")
	}
}

func (c *Client) Set(ctx context.Context, ns, key string, value any, exp time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.DB.Set(ctx, fmt.Sprintf("%v:%v", ns, key), bytes, exp).Err()
}

func (c *Client) Get(ctx context.Context, ns, key string, value any) error {
	bytes, err := c.DB.Get(ctx, fmt.Sprintf("%v:%v", ns, key)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return errors.New(codes.NotFound).
				WithDetails(err.Error())
		}
		return errors.New(codes.Unknown).
			WithDetails(err.Error())
	}
	if err = json.Unmarshal(bytes, value); err != nil {
		return errors.Wrap(err, codes.InvalidArgument)
	}
	return nil
}

func (c *Client) Del(ctx context.Context, ns, key string) error {
	if err := c.DB.Del(ctx, fmt.Sprintf("%v:%v", ns, key)).Err(); err != nil {
		if err == redis.Nil {
			return errors.New(codes.NotFound).
				WithDetails(err.Error())
		}
		return errors.New(codes.Unknown).
			WithDetails(err.Error())
	}
	return nil
}

func (c *Client) Exists(ctx context.Context, ns, key string) error {
	if err := c.DB.Get(ctx, fmt.Sprintf("%v:%v", ns, key)).Err(); err != nil {
		if err == redis.Nil {
			return errors.New(codes.NotFound).
				WithDetails(err.Error())
		}
		return errors.New(codes.Unknown).
			WithDetails(err.Error())
	}
	return nil
}

func (c *Client) Expire(ctx context.Context, ns, key string, exp time.Duration) error {
	if err := c.DB.Expire(ctx, fmt.Sprintf("%v:%v", ns, key), exp).Err(); err != nil {
		if err == redis.Nil {
			return errors.New(codes.NotFound).
				WithDetails(err.Error())
		}
		return errors.New(codes.Unknown).
			WithDetails(err.Error())
	}
	return nil
}

func (c *Client) OpenQueue(name string) (*Queue, error) {
	queue, err := c.Queue.OpenQueue(name)
	if err != nil {
		return nil, err
	}
	queue.SetPushQueue(queue)
	return &Queue{
		cli:  c,
		base: queue,
	}, nil
}

func (c *Client) NewMutex(ns, key string, exp time.Duration) *redsync.Mutex {
	return c.Lock.NewMutex(fmt.Sprintf("%v:%v", ns, key), redsync.WithExpiry(exp))
}

func (c *Client) ObtainAtomic(ctx context.Context, ns, key string, exp time.Duration) (*redislock.Lock, error) {
	lock, err := c.Atomic.Obtain(ctx, fmt.Sprintf("%v:%v", ns, key), exp, nil)
	if goerrors.Is(err, redislock.ErrNotObtained) {
		return nil, errors.Wrap(err, codes.AlreadyExists)
	} else if err != nil {
		return nil, err
	}
	return lock, nil
}

func (c *Client) Limit(ctx context.Context, ns, key string, limit redis_rate.Limit) (*redis_rate.Result, error) {
	res, err := c.Limiter.Allow(ctx, fmt.Sprintf("%v:%v", ns, key), limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) LimitReport(ctx context.Context, ns, key string, limit redis_rate.Limit) (*redis_rate.Result, error) {
	res, err := c.Limiter.AllowAtMost(ctx, fmt.Sprintf("%v:%v", ns, key), limit, 0)
	if err != nil {
		return nil, err
	}
	return res, nil
}
