package redis

import (
	"context"
	"encoding/json"
	"time"

	conf "cinema/pkg/config"
	"cinema/pkg/logging"
	redis1 "cinema/pkg/redis"

	redis2 "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis interface {
	Get(ctx context.Context, slug string, dis interface{}) (interface{}, error)
	Set(ctx context.Context, slug string, expireTime time.Duration, t interface{})
	Delete(ctx context.Context, key ...string)
}
type redis struct {
	redis  redis2.Client
	logger *zap.Logger
}

func New(cfg *conf.Config, logger *zap.Logger, enable bool) Redis {
	if !enable {
		return Disabled()
	}
	redisClient, err := redis1.New(cfg.GetRedis())
	if err != nil {
		logger.Fatal("can not init redis", zap.Error(err))
	}
	return &redis{
		redis:  *redisClient,
		logger: logger,
	}
}

func (r *redis) Get(ctx context.Context, key string, dis interface{}) (interface{}, error) {
	cmd := r.redis.Get(ctx, key)
	bytes := []byte(cmd.Val())
	if cmd.Val() == "" {
		return nil, nil
	}
	err := json.Unmarshal(bytes, dis)
	if err != nil {
		return nil, err
	}
	r.logger.Debug("[RedisCache][GetCache] Get" + key)
	return dis, nil
}

func (r *redis) Delete(ctx context.Context, keys ...string) {
	r.redis.Del(ctx, keys...)
	r.logger.Debug("[RedisCache][RemoveCache]")
}

func (r *redis) Set(ctx context.Context, key string, expireTime time.Duration, t interface{}) {
	logging.Logger(ctx).Debug("[RedisCache][SetCache] Set" + key)
	bytes, err := json.Marshal(t)
	if err != nil {
		return
	}
	r.redis.Set(ctx, key, bytes, expireTime)
}
