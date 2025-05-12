package redis

import (
	"context"
	"time"
)

type disabledRedis struct{}

func Disabled() Redis {
	return &disabledRedis{}
}

func (d *disabledRedis) Get(ctx context.Context, key string, dis interface{}) (interface{}, error) {
	return nil, nil
}

func (d *disabledRedis) Set(ctx context.Context, key string, expireTime time.Duration, t interface{}) {
}

func (d *disabledRedis) Delete(ctx context.Context, key ...string) {
}
