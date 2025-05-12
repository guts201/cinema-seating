package usecase

import (
	cinemaservice "cinema/internal/usecase/cinema"
	movieservice "cinema/internal/usecase/movie"
	screeningservice "cinema/internal/usecase/screening"
	seatingservice "cinema/internal/usecase/seating"
	"cinema/pkg/ent"
)

type UseCase struct {
	CinemaService    cinemaservice.CinemaUseCase
	SeatingService   *seatingservice.SeatingService
	MovieService     movieservice.MovieUseCase
	ScreeningService screeningservice.ScreeningUseCase
}

func New(entClient *ent.Client) *UseCase {
	return &UseCase{
		CinemaService:    cinemaservice.NewCinemaService(entClient),
		SeatingService:   seatingservice.NewSeatingService(),
		MovieService:     movieservice.NewMovieService(entClient),
		ScreeningService: screeningservice.NewScreeningService(entClient),
	}
}
