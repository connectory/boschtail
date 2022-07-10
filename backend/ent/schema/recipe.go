package schema

import "entgo.io/ent"

// Recipe holds the schema definition for the Recipe entity.
type Recipe struct {
	ent.Schema
}

// Fields of the Recipe.
func (Recipe) Fields() []ent.Field {
	return nil
}

// Edges of the Recipe.
func (Recipe) Edges() []ent.Edge {
	return nil
}
