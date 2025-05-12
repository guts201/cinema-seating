package backofficecinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/mapper"
)

func (s *backofficeCinemaServer) CreateCinema(ctx context.Context, request *cinema.CreateCinemaRequest) (*cinema.CreateCinemaResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	_cinema := mapper.FromProtoReqCinema(request)

	res, err := s.usecase.CinemaService.CreateCinema(ctx, _cinema)
	if err != nil {
		return nil, err
	}

	return &cinema.CreateCinemaResponse{
		Id: int64(res.ID),
	}, nil
}
