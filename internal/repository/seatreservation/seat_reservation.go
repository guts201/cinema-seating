package seatreservation

import (
	cinema "cinema/api"
	"cinema/internal/entity"
	"cinema/internal/mapper"
	"cinema/internal/redis"
	"cinema/pkg/ent"
	"cinema/pkg/ent/screening"
	"cinema/pkg/ent/seatreservation"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	seatReservationKey = "seat-reservation-%s-%d"
)

type SeatReservationRepository interface {
	CreateSeatReservation(ctx context.Context, groupId uuid.UUID, screeningId int, seatReservations []entity.Seat) ([]entity.SeatWithId, error)
	GetSeatReservation(ctx context.Context, status cinema.SeatReservationStatus, screeningId int64) ([]entity.Seat, error)
	DeleteSeatReservation(ctx context.Context, screenId int, seatReservationIds []int64) error
}

type seatReservationRepository struct {
	client *ent.Client
	redis  redis.Redis
}

func New(client *ent.Client, redis redis.Redis) SeatReservationRepository {
	return &seatReservationRepository{
		client: client,
		redis:  redis,
	}
}

func (r seatReservationRepository) CreateSeatReservation(ctx context.Context, groupId uuid.UUID, screeningId int, seatReservations []entity.Seat) ([]entity.SeatWithId, error) {
	creates := make([]*ent.SeatReservationCreate, len(seatReservations))
	for i, s := range seatReservations {
		creates[i] = r.client.SeatReservation.Create().
			SetScreeningID(int64(screeningId)).
			SetGroupID(groupId).
			SetStatus(cinema.SeatReservationStatus_RESERVED).
			SetRowNum(uint32(s.Row)).
			SetColumnNum(uint32(s.Column))
	}
	tpm, err := r.client.SeatReservation.CreateBulk(creates...).Save(ctx)

	if err != nil {
		return nil, err
	}
	if len(tpm) > 0 {
		r.redis.Delete(ctx, fmt.Sprintf(seatReservationKey, cinema.SeatReservationStatus_RESERVED, int64(screeningId)))
	}

	var seatWithIds []entity.SeatWithId
	for _, s := range tpm {
		seatWithIds = append(seatWithIds, mapper.FromRepoSeatWithId(s))
	}

	return seatWithIds, err
}

func (r seatReservationRepository) GetSeatReservation(ctx context.Context, status cinema.SeatReservationStatus, screeningId int64) ([]entity.Seat, error) {

	key := fmt.Sprintf(seatReservationKey, status, screeningId)
	seatsFromCache, err := r.redis.Get(ctx, key, nil)
	if err == nil && seatsFromCache != nil {
		return seatsFromCache.([]entity.Seat), nil
	}

	seatReservations, err := r.client.SeatReservation.Query().
		Where(seatreservation.StatusEQ(status)).
		Where(seatreservation.HasScreeningWith(screening.IDEQ(screeningId))).
		All(ctx)

	if err != nil {
		return nil, err
	}

	seats := make([]entity.Seat, len(seatReservations))
	for i, s := range seatReservations {
		seats[i] = mapper.FromRepoSeat(s)
	}

	if len(seats) > 0 {
		r.redis.Set(ctx, key, time.Minute*1, seats)
	}

	return seats, nil
}

func (r seatReservationRepository) DeleteSeatReservation(ctx context.Context, screenId int, seatReservationIds []int64) error {
	eff, err := r.client.SeatReservation.Delete().
		Where(seatreservation.HasScreeningWith(screening.IDEQ(int64(screenId)))).
		Where(seatreservation.IDIn(seatReservationIds...)).
		Exec(ctx)

	if eff > 0 {
		r.redis.Delete(ctx, fmt.Sprintf(seatReservationKey, cinema.SeatReservationStatus_RESERVED, int64(screenId)))
	}
	return err
}
