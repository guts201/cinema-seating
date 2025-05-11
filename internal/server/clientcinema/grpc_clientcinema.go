package clientcinema

import (
	"cinema/api"
)

func NewServer() cinema.ClientCinemaServer {
	return &clientCinemaServer{}
}

type clientCinemaServer struct {
	cinema.UnimplementedClientCinemaServer
}
