package clientcinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/mapper"
)

func (s *clientCinemaServer) ListScreening(ctx context.Context, request *cinema.ListScreeningRequest) (*cinema.ListScreeningResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.Screening.ListScreening(ctx, int(request.CinemaId))
	if err != nil {
		return nil, err
	}

	_screenings := mapper.ToProtoScreenings(res)
	return &cinema.ListScreeningResponse{
		Screenings: _screenings,
	}, nil

}
