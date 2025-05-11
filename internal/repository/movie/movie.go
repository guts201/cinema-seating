package movie

import (
	"cinema/internal/entity"
	"cinema/internal/mapper"
	"cinema/pkg/ent"
	"context"
)

type Movie interface {
	CreateMovie(ctx context.Context, movie entity.Movie) (*entity.Movie, error)
}

type movieRepo struct {
	client *ent.Client
}

func New(client *ent.Client) Movie {
	return &movieRepo{
		client: client,
	}
}

func (c movieRepo) CreateMovie(ctx context.Context, cinema entity.Movie) (*entity.Movie, error) {
	_movie, err := c.client.Movie.Create().
		SetTitle(cinema.Title).
		SetDuration(uint64(cinema.DurationMinutes)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromRepoMovie(_movie), nil
}
