package schema

import (
	"entgo.io/ent"
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
