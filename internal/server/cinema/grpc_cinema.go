package cinema

import (
	cinema "cinema/api"
)

func NewServer() cinema.CinemaServer {
	return &cinemaServer{}
}

type cinemaServer struct {
	cinema.UnimplementedCinemaServer
}
