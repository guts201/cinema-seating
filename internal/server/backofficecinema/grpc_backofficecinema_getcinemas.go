package backofficecinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/mapper"
)

func (s *backofficeCinemaServer) GetCinemas(ctx context.Context, request *cinema.GetCinemasRequest) (*cinema.GetCinemasResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	_cinemas, err := s.usecase.CinemaService.GetCinemas(ctx, int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, err
	}

	cinemasRes := mapper.ToProtoCinemas(_cinemas)

	return &cinema.GetCinemasResponse{
		Cinemas: cinemasRes,
	}, nil
}
