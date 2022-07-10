package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			NamedValues(
				"Ordered", "ORDERED",
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Annotations(
				entgql.OrderField("STATUS"),
			).
			Default("ORDERED"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
