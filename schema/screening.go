package schema

import (
	"entgo.io/ent"
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
		field.Int32("min_distance").Min(0),
	}
}

func (Screening) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
