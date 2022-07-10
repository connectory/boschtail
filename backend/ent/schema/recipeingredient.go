package schema

import "entgo.io/ent"

// RecipeIngredient holds the schema definition for the RecipeIngredient entity.
type RecipeIngredient struct {
	ent.Schema
}

// Fields of the RecipeIngredient.
func (RecipeIngredient) Fields() []ent.Field {
	return nil
}

// Edges of the RecipeIngredient.
func (RecipeIngredient) Edges() []ent.Edge {
	return nil
}
