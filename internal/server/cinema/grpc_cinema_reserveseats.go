package cinema

import (
	"context"

	"cinema/api"
)

func (s *cinemaServer) ReserveSeats(ctx context.Context, request *cinema.ReserveSeatsRequest) (*cinema.ReserveSeatsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &cinema.ReserveSeatsResponse{}, nil
}
