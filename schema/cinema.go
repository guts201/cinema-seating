package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Cinema struct {
	ent.Schema
}

func (Cinema) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("num_row").Positive(),
		field.Uint32("num_column").Positive(),
		field.String("name").
			NotEmpty(),
		field.Uint32("min_distance").Min(0),
	}
}

func (Cinema) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}

func (Cinema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("screenings", Screening.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
