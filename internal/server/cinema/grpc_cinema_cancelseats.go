package cinema

import (
	"context"

	"cinema/api"
)

func (s *cinemaServer) CancelSeats(ctx context.Context, request *cinema.CancelSeatsRequest) (*cinema.CancelSeatsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &cinema.CancelSeatsResponse{}, nil
}
