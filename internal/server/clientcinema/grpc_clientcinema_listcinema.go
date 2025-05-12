package clientcinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/mapper"
)

func (s *clientCinemaServer) ListCinema(ctx context.Context, request *cinema.ListCinemaRequest) (*cinema.ListCinemaResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	_cinemas, err := s.usecase.CinemaService.GetCinemas(ctx, int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, err
	}

	cinemasRes := mapper.ToProtoCinemas(_cinemas)

	return &cinema.ListCinemaResponse{
		Cinemas: cinemasRes,
	}, nil
}
