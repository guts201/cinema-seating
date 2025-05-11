package backofficecinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/entity"
)

func (s *backofficeCinemaServer) CreateCinema(ctx context.Context, request *cinema.CreateCinemaRequest) (*cinema.CreateCinemaResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	_cinema := entity.Cinema{
		Name:        request.Name,
		MinDistance: int(request.MinDistance),
		Rows:        int(request.Rows),
		Columns:     int(request.Columns),
	}

	res, err := s.usecase.CinemaService.CreateCinema(ctx, _cinema)
	if err != nil {
		return nil, err
	}

	return &cinema.CreateCinemaResponse{
		Id: int64(res.ID),
	}, nil
}
