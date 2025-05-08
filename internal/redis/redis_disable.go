package redis

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"

	"cinema/pkg/logging"
)

type disabled struct{}

func Disabled() Redis {
	logging.Logger(context.Background()).Warn("Redis was created with Disabled implement")
	return &disabled{}
}

func (s disabled) Get(ctx context.Context, key string, dis proto.Message) (proto.Message, error) {

	return nil, nil
}
func (s disabled) Set(ctx context.Context, key string, expireTime time.Duration, t proto.Message) {

}

func (s disabled) Delete(ctx context.Context, key ...string) {

}
