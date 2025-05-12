package screeningservice

import (
	"cinema/internal/entity"
	"cinema/internal/repository/screening"
	"cinema/pkg/ent"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScreeningUseCase interface {
	CreateScreening(ctx context.Context, title string, movieID, cinemaId int64, startTime *timestamppb.Timestamp) (entity.Screening, error)
	ListScreening(ctx context.Context, cinemaId int) ([]entity.Screening, error)
}
type screeningService struct {
	screeningRepo screening.ScreeningRepository
}

func NewScreeningService(entClient *ent.Client) ScreeningUseCase {
	screeningRepo := screening.New(entClient)
	return &screeningService{
		screeningRepo: screeningRepo,
	}
}

func (s screeningService) CreateScreening(ctx context.Context, title string, movieID, cinemaId int64, startTime *timestamppb.Timestamp) (entity.Screening, error) {
	return s.screeningRepo.CreateScreening(ctx, title, int(movieID), int(cinemaId), startTime)
}

func (s screeningService) ListScreening(ctx context.Context, cinemaId int) ([]entity.Screening, error) {
	return s.screeningRepo.ListScreening(ctx, cinemaId)
}