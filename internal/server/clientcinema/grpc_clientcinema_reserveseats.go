package clientcinema

import (
	"context"

	"cinema/api"
)

func (s *clientCinemaServer) ReserveSeats(ctx context.Context, request *cinema.ReserveSeatsRequest) (*cinema.ReserveSeatsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &cinema.ReserveSeatsResponse{}, nil
}
