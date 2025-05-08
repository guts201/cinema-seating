package server

import (
	mykit "cinema/pkg/genkit/pkg/api"
	"google.golang.org/grpc/reflection"

	pb0 "cinema/api"
	"cinema/internal/server/cinema"
	config "cinema/pkg/config"
)

// Serve ...
func Serve(cfg *config.Config) {
	service := newService(cfg, []mykit.Option{}...)

	server := service.Server()
	pb0.RegisterCinemaServer(server, cinema.NewServer())

	// Register reflection service on gRPC server.
	// Please remove if you it's not necessary for your service
	reflection.Register(server)

	service.Serve()
}
