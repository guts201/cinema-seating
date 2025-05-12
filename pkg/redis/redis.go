package redis

import (
	"context"
	"os"

	redis "github.com/redis/go-redis/v9"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/redis/go-redis.v9"

	api "cinema/pkg/config"
)

func New(config *api.Redis, opts ...Option) (*redis.Client, error) {
	o := &Opt{
		Options: &redis.Options{
			Addr: config.GetAddress(),
		},
	}

	for _, o0 := range opts {
		o0.Apply(o)
	}

	client := redis.NewClient(o.Options)
	redistrace.WrapClient(client, redistrace.WithServiceName(os.Getenv("DD_SERVICE")))
	return client, client.Ping(context.Background()).Err()
}

type Opt struct {
	*redis.Options
}

type Option interface {
	Apply(o *Opt)
}

type OptionFunc func(*Opt)

func (f OptionFunc) Apply(o *Opt) {
	f(o)
}

// Limiter interface used to implemented circuit breaker or rate limiter.
func Limiter(limiter redis.Limiter) Option {
	return OptionFunc(func(o *Opt) {
		o.Limiter = limiter
	})
}
