package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Seat struct {
	ent.Schema
}

func (Seat) Fields() []ent.Field {
	return []ent.Field{
		field.Int16("row").Positive(),
		field.Int16("column").Positive(),
	}
}

func (Seat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Seat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cinema", Cinema.Type).Ref("seats").Unique(),
		edge.To("seat_reservations", SeatReservation.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
