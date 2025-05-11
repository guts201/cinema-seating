package movieservice

import (
	"cinema/internal/entity"
	"cinema/internal/repository/movie"
	"cinema/pkg/ent"
	"context"
)

type MovieUseCase interface {
	CreateMovie(ctx context.Context, movie entity.Movie) (*entity.Movie, error)
}
type movieService struct {
	movieRepo movie.Movie
}

func NewMovieService(entClient *ent.Client) MovieUseCase {
	movieRepo := movie.New(entClient)
	return &movieService{
		movieRepo: movieRepo,
	}
}

func (m movieService) CreateMovie(ctx context.Context, movie entity.Movie) (*entity.Movie, error) {
	return m.movieRepo.CreateMovie(ctx, movie)
}
