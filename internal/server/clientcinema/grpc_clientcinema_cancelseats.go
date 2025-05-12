package clientcinema

import (
	"context"

	cinema "cinema/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *clientCinemaServer) CancelSeats(ctx context.Context, request *cinema.CancelSeatsRequest) (*cinema.CancelSeatsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	_, err := s.repo.Screening.GetScreening(ctx, int(request.ScreeningId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "screening not found")
	}


	err = s.repo.SeatReservation.DeleteSeatReservation(ctx, int(request.ScreeningId), request.GetSeatIds())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "reservation failed")
	}
	return &cinema.CancelSeatsResponse{}, nil
}
