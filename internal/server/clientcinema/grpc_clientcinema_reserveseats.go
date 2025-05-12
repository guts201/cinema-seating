package clientcinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/mapper"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *clientCinemaServer) ReserveSeats(ctx context.Context, request *cinema.ReserveSeatsRequest) (*cinema.ReserveSeatsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// 1. Load screening from DB
	screening, err := s.repo.Screening.GetScreening(ctx, int(request.ScreeningId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "screening not found")
	}

	// 2. Load reserved seats from Redis/DB
	reserved, err := s.repo.SeatReservation.GetSeatReservation(ctx, cinema.SeatReservationStatus_RESERVED, int64(request.ScreeningId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to load reserved seats from cache")
	}

	// 3. Validate business rules
	group := mapper.FromProtoSeatGroup(request.GetGroup())
	if !s.usecase.SeatingService.IsValidSeatGroup(screening.Row, screening.Column, group.Seats, reserved, screening.MinDistance) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid seat group")
	}

	// 4. Reserve seats (DB + Redis)
	res, err := s.repo.SeatReservation.CreateSeatReservation(ctx, uuid.New(), int(request.ScreeningId), group.Seats)
	if err != nil {
		return nil, err
	}

	return &cinema.ReserveSeatsResponse{
		Seats: mapper.ToProtoSeats(res),
	}, nil
}
