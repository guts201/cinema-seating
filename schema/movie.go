package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Movie struct {
	ent.Schema
}

func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.Uint64("duration"),
	}
}

func (Movie) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
