package redis

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"
)

func (r *redis) Get(ctx context.Context, key string, dis proto.Message) (proto.Message, error) {
	cmd := r.redis.Get(ctx, key)
	bytes := []byte(cmd.Val())
	if cmd.Val() == "" {
		return nil, nil
	}
	err := proto.Unmarshal(bytes, dis)
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

func (r *redis) Set(ctx context.Context, key string, expireTime time.Duration, t proto.Message) {
	bytes, err := proto.Marshal(t)
	if err != nil {
		return
	}
	r.redis.Set(ctx, key, bytes, expireTime)
}
