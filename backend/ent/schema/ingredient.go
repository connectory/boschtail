package schema

import "entgo.io/ent"

// Ingredient holds the schema definition for the Ingredient entity.
type Ingredient struct {
	ent.Schema
}

// Fields of the Ingredient.
func (Ingredient) Fields() []ent.Field {
	return nil
}

// Edges of the Ingredient.
func (Ingredient) Edges() []ent.Edge {
	return nil
}
