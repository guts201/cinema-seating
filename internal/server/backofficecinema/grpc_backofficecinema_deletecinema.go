package backofficecinema

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	cinema "cinema/api"
)

func (s *backofficeCinemaServer) DeleteCinema(ctx context.Context, request *cinema.DeleteCinemaRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	err := s.usecase.CinemaService.DeleteCinema(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
