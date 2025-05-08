package redis

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"

	c "github.com/go-redis/cache/v9"
	r "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis interface {
	Get(ctx context.Context, slug string, dis proto.Message) (proto.Message, error)
	Set(ctx context.Context, slug string, expireTime time.Duration, t proto.Message)
	Delete(ctx context.Context, key ...string)
}

type redis struct {
	redis  r.Cmdable
	cache  *c.TinyLFU
	logger *zap.Logger
	Redis
}

func New(origin *r.Client, logger *zap.Logger) Redis {
	if origin == nil {
		return Disabled()
	}
	return &redis{
		cache:  c.NewTinyLFU(cacheLocalMaxCount, cacheLocalTime),
		logger: logger,
		redis:  origin,
	}
}
