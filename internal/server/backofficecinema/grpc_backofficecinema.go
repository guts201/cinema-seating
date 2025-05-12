package backofficecinema

import (
	cinema "cinema/api"
	"cinema/internal/usecase"
	"cinema/pkg/ent"
)

func NewServer(entClient *ent.Client) cinema.BackofficeCinemaServer {
	return &backofficeCinemaServer{
		usecase: usecase.New(entClient),
	}
}

type backofficeCinemaServer struct {
	usecase *usecase.UseCase
	cinema.UnimplementedBackofficeCinemaServer
}
