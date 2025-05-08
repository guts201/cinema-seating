package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Cinema struct {
	ent.Schema
}

func (Cinema) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("num_row").Positive(),
		field.Int64("num_column").Positive(),
		field.String("name").
			NotEmpty(),
		field.String("address").
			NotEmpty(),
	}
}

func (Cinema) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
