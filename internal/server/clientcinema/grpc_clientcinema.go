package clientcinema

import (
	cinema "cinema/api"
	"cinema/internal/redis"
	"cinema/internal/repository"
	"cinema/internal/usecase"
	"cinema/pkg/ent"
)

func NewServer(entClient *ent.Client, redis redis.Redis) cinema.ClientCinemaServer {
	return &clientCinemaServer{
		usecase: *usecase.New(entClient),
		repo:    *repository.NewMovieRepository(redis, entClient),
	}
}

type clientCinemaServer struct {
	usecase usecase.UseCase
	repo    repository.Repository
	cinema.UnimplementedClientCinemaServer
}
