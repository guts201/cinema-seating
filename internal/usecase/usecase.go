package usecase

import (
	cinemaservice "cinema/internal/usecase/cinema"
	movieservice "cinema/internal/usecase/movie"
	seatingservice "cinema/internal/usecase/seating"
	"cinema/pkg/ent"
)

type UseCase struct {
	CinemaService  cinemaservice.CinemaUseCase
	SeatingService *seatingservice.SeatingService
	MovieService   movieservice.MovieUseCase
}

func New(entClient *ent.Client) *UseCase {
	return &UseCase{
		CinemaService:  cinemaservice.NewCinemaService(entClient),
		SeatingService: seatingservice.NewSeatingService(),
		MovieService:   movieservice.NewMovieService(entClient),
	}
}
