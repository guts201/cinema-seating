package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Screening struct {
	ent.Schema
}

func (Screening) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.Time("start_time"),
	}
}

func (Screening) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Screening) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("movie", Movie.Type).Ref("screenings").Unique(),
		edge.From("cinema", Cinema.Type).Ref("screenings").Unique(),
		edge.To("seat_reservations", SeatReservation.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
