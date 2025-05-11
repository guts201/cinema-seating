package seatreservation

import (
	"cinema/internal/entity"
	"cinema/pkg/ent"
	"context"

	"github.com/google/uuid"
)

type seatReservationRepository struct {
	client *ent.Client
}

func New(client *ent.Client) *seatReservationRepository {
	return &seatReservationRepository{
		client: client,
	}
}

func (r seatReservationRepository) CreateSeatReservation(ctx context.Context, groupId uuid.UUID, screeningId int, seatReservationIds []int) (*entity.SeatGroup, error) {
	creates := make([]*ent.SeatReservationCreate, len(seatReservationIds))
	for i, s := range seatReservationIds {
		creates[i] = r.client.SeatReservation.Create().
			SetScreeningID(int64(screeningId)).
			SetGroupID(groupId).
			SetStatus("reserved").
			SetSeatID(int64(s))
	}
	_, err := r.client.SeatReservation.CreateBulk(creates...).Save(ctx)
	return nil, err
}

func (r seatReservationRepository) GetSeatReservation(ctx context.Context, status string, screeningId int64) ([]*entity.SeatGroup, error) {
	// seatReservations, err := r.client.SeatReservation.Query().
	// 	Where(seatreservation.StatusEQ(status)).
	// 	Where(seatreservation.HasScreeningWith(screening.IDEQ(screeningId))).
	// 	All(ctx)

	return nil, nil
}

func (r seatReservationRepository) DeleteSeatReservation(ctx context.Context, id int) error {
	return nil
}
