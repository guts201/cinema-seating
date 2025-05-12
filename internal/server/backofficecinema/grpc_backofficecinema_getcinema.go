package backofficecinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/mapper"
)

func (s *backofficeCinemaServer) GetCinema(ctx context.Context, request *cinema.GetCinemaRequest) (*cinema.GetCinemaResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	res, err := s.usecase.CinemaService.GetCinema(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	_cinema := mapper.ToProtoCinema(*res)

	return &cinema.GetCinemaResponse{
		Cinema: _cinema,
	}, nil
}
