package cinemaservice

import (
	"cinema/internal/entity"
	"cinema/internal/repository/cinema"
	"cinema/pkg/ent"
	"context"
)

type CinemaUseCase interface {
	CreateCinema(ctx context.Context, cinema entity.Cinema) (*entity.Cinema, error)
	GetCinema(ctx context.Context, id int) (*entity.Cinema, error)
	GetCinemas(ctx context.Context, limit, offset int) ([]entity.Cinema, error)
	UpdateCinema(ctx context.Context, id int, cinema entity.Cinema) (*entity.Cinema, error)
	DeleteCinema(ctx context.Context, id int) error
}
type cinemaService struct {
	cinemaRepo cinema.Cinema
}

func NewCinemaService(entClient *ent.Client) CinemaUseCase {
	return &cinemaService{
		cinemaRepo: cinema.New(entClient),
	}
}

func (c *cinemaService) CreateCinema(ctx context.Context, cinema entity.Cinema) (*entity.Cinema, error) {
	return c.cinemaRepo.CreateCinema(ctx, cinema)
}

func (c *cinemaService) GetCinema(ctx context.Context, id int) (*entity.Cinema, error) {
	return c.cinemaRepo.GetCinema(ctx, id)
}

func (c *cinemaService) GetCinemas(ctx context.Context, limit, offset int) ([]entity.Cinema, error) {
	return c.cinemaRepo.GetCinemas(ctx, limit, offset)
}

func (c *cinemaService) UpdateCinema(ctx context.Context, id int, cinema entity.Cinema) (*entity.Cinema, error) {
	return c.cinemaRepo.UpdateCinema(ctx, id, cinema)
}

func (c *cinemaService) DeleteCinema(ctx context.Context, id int) error {
	return c.cinemaRepo.DeleteCinema(ctx, id)
}
