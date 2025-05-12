package screening

import (
	"cinema/internal/entity"
	"cinema/internal/mapper"
	"cinema/pkg/ent"
	cinemaEnt "cinema/pkg/ent/cinema"
	"cinema/pkg/ent/screening"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScreeningRepository interface {
	GetScreening(ctx context.Context, id int) (entity.Screening, error)
	CreateScreening(ctx context.Context, title string, movieId, cinemaId int, startTime *timestamppb.Timestamp) (entity.Screening, error)
	ListScreening(ctx context.Context, cinemaId int) ([]entity.Screening, error)
}

type screeningRepository struct {
	client *ent.Client
}

func New(client *ent.Client) ScreeningRepository {
	return &screeningRepository{
		client: client,
	}
}

func (s screeningRepository) GetScreening(ctx context.Context, id int) (entity.Screening, error) {
	res, err := s.client.Screening.Query().Where(screening.IDEQ(int64(id))).WithCinema().WithMovie().Only(ctx)
	return mapper.FromRepoScreening(res), err
}

func (s screeningRepository) CreateScreening(ctx context.Context, title string, movieId, cinemaId int, startTime *timestamppb.Timestamp) (entity.Screening, error) {
	_screening, err := s.client.Screening.Create().
		SetTitle(title).
		SetStartTime(startTime.AsTime()).
		SetMovieID(int64(movieId)).
		SetCinemaID(int64(cinemaId)).
		Save(ctx)
	return mapper.FromRepoScreening(_screening), err
}

func (s screeningRepository) ListScreening(ctx context.Context, cinemaId int) ([]entity.Screening, error) {
	res, err := s.client.Screening.Query().
		Where(screening.HasCinemaWith(cinemaEnt.IDEQ(int64(cinemaId)))).
		WithCinema().
		WithMovie().
		All(ctx)

	return mapper.FromRepoScreenings(res), err
}
