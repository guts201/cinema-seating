package backofficecinema

import (
	"context"
	"fmt"

	cinema "cinema/api"
)

func (s *backofficeCinemaServer) CreateScreening(ctx context.Context, request *cinema.CreateScreeningRequest) (*cinema.CreateScreeningResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	res, err := s.usecase.ScreeningService.CreateScreening(ctx, randomString(10), int64(request.MovieId), int64(request.CinemaId), request.StartTime)
	if err != nil {
		return nil, err
	}

	return &cinema.CreateScreeningResponse{
		Id: int64(res.ID),
	}, nil
}

func randomString(length int) string {
	b := make([]byte, length+2)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
