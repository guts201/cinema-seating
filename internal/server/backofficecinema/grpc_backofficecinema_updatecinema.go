package backofficecinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/entity"
)

func (s *backofficeCinemaServer) UpdateCinema(ctx context.Context, request *cinema.UpdateCinemaRequest) (*cinema.UpdateCinemaResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	_cinema := entity.Cinema{
		Name:        request.Name,
		MinDistance: int(request.MinDistance),
		Rows:        int(request.Rows),
		Columns:     int(request.Columns),
	}

	_, err := s.usecase.CinemaService.UpdateCinema(ctx, int(request.Id), _cinema)
	if err != nil {
		return nil, err
	}

	return &cinema.UpdateCinemaResponse{}, nil
}
