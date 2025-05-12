package cinema

import (
	"cinema/internal/entity"
	"cinema/internal/mapper"
	"cinema/pkg/ent"
	"context"
)

type Cinema interface {
	CreateCinema(ctx context.Context, cinema entity.Cinema) (*entity.Cinema, error)
	GetCinema(ctx context.Context, id int) (*entity.Cinema, error)
	GetCinemas(ctx context.Context, limit, offset int) ([]entity.Cinema, error)
	UpdateCinema(ctx context.Context, id int, cinema entity.Cinema) (*entity.Cinema, error)
	DeleteCinema(ctx context.Context, id int) error
}

type cinemaRepo struct {
	client *ent.Client
}

func New(client *ent.Client) Cinema {
	return &cinemaRepo{
		client: client,
	}
}

func (c cinemaRepo) CreateCinema(ctx context.Context, cinema entity.Cinema) (*entity.Cinema, error) {
	_cinema, err := c.client.Cinema.Create().
		SetName(cinema.Name).
		SetMinDistance(uint32(cinema.MinDistance)).
		SetNumRow(uint32(cinema.Rows)).
		SetNumColumn(uint32(cinema.Columns)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromRepoCinema(_cinema), nil
}

func (c cinemaRepo) GetCinema(ctx context.Context, id int) (*entity.Cinema, error) {
	_cinema, err := c.client.Cinema.Get(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	return mapper.FromRepoCinema(_cinema), nil
}

func (c cinemaRepo) GetCinemas(ctx context.Context, limit, offset int) ([]entity.Cinema, error) {
	_cinemas, err := c.client.Cinema.Query().
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromRepoCinemas(_cinemas), nil
}

func (c cinemaRepo) UpdateCinema(ctx context.Context, id int, cinema entity.Cinema) (*entity.Cinema, error) {
	_cinema, err := c.client.Cinema.UpdateOneID(int64(id)).
		SetName(cinema.Name).
		SetMinDistance(uint32(cinema.MinDistance)).
		SetNumRow(uint32(cinema.Rows)).
		SetNumColumn(uint32(cinema.Columns)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromRepoCinema(_cinema), nil
}

func (c cinemaRepo) DeleteCinema(ctx context.Context, id int) error {
	err := c.client.Cinema.DeleteOneID(int64(id)).Exec(ctx)
	return err
}
