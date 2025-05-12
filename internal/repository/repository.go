package repository

import (
	"cinema/internal/redis"
	"cinema/internal/repository/screening"
	"cinema/internal/repository/seatreservation"
	"cinema/pkg/ent"
)

type Repository struct {
	Screening       screening.ScreeningRepository
	SeatReservation seatreservation.SeatReservationRepository
}

func NewMovieRepository(redis redis.Redis, client *ent.Client) *Repository {
	return &Repository{
		Screening:       screening.New(client),
		SeatReservation: seatreservation.New(client, redis),
	}
}
