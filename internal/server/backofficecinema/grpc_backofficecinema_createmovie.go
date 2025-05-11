package backofficecinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/entity"
)

func (s *backofficeCinemaServer) CreateMovie(ctx context.Context, request *cinema.CreateMovieRequest) (*cinema.CreateMovieResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	_movie := entity.Movie{
		Title:           request.Title,
		DurationMinutes: int(request.DurationMinutes),
	}

	res, err := s.usecase.MovieService.CreateMovie(ctx, _movie)
	if err != nil {
		return nil, err
	}

	return &cinema.CreateMovieResponse{
		Id: int64(res.ID),
	}, nil
}
