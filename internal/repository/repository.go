package repository

import (
	"cinema/internal/redis"
	"cinema/pkg/ent"
)

type Repository struct {
	redis  redis.Redis
	client *ent.Client
}

func NewMovieRepository(redis redis.Redis, client *ent.Client) *Repository {
	return &Repository{
		redis:  redis,
		client: client,
	}
}
