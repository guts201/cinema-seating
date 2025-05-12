package schema

import (
	cinema "cinema/api"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type SeatReservation struct {
	ent.Schema
}

func (SeatReservation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("reserved_at").Default(time.Now),
		field.UUID("group_id", uuid.UUID{}),
		field.Int32("status").GoType(cinema.SeatReservationStatus(0)).Default(0),
		// field.Time("start_time").Default(time.Now),
		// field.Time("end_time").Default(time.Now().Add(5 * time.Minute)),
		field.Uint32("row_num").Min(0),
		field.Uint32("column_num").Min(0),
	}
}

func (SeatReservation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (SeatReservation) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("screening", Screening.Type).Ref("seat_reservations").Unique(),
	}
}
