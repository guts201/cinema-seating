package clientcinema

import (
	"context"

	cinema "cinema/api"
	"cinema/internal/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *clientCinemaServer) GetAvailableGroups(ctx context.Context, request *cinema.GetAvailableGroupsRequest) (*cinema.GetAvailableGroupsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	screening, err := s.repo.Screening.GetScreening(ctx, int(request.ScreeningId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "screening not found")
	}

	// 2. Load reserved seats from Redis/DB
	reserved, err := s.repo.SeatReservation.GetSeatReservation(ctx, cinema.SeatReservationStatus_RESERVED, int64(request.ScreeningId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to load reserved seats from cache")
	}

	// 4. Reserve seats (DB + Redis)
	seats := s.usecase.SeatingService.GetAvailableSeats(screening.Row, screening.Column, reserved, screening.MinDistance)

	return &cinema.GetAvailableGroupsResponse{
		Groups:         mapper.ToProtoSeatGroup(seats),
		RemainingSeats: int32(len(seats)),
	}, nil
}
