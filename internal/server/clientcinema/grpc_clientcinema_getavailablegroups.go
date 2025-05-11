package clientcinema

import (
	"context"

	"cinema/api"
)

func (s *clientCinemaServer) GetAvailableGroups(ctx context.Context, request *cinema.GetAvailableGroupsRequest) (*cinema.GetAvailableGroupsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &cinema.GetAvailableGroupsResponse{}, nil
}
