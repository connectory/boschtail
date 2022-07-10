package schema

import "entgo.io/ent"

// Mixer holds the schema definition for the Mixer entity.
type Mixer struct {
	ent.Schema
}

// Fields of the Mixer.
func (Mixer) Fields() []ent.Field {
	return nil
}

// Edges of the Mixer.
func (Mixer) Edges() []ent.Edge {
	return nil
}
