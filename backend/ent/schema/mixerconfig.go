package schema

import "entgo.io/ent"

// MixerConfig holds the schema definition for the MixerConfig entity.
type MixerConfig struct {
	ent.Schema
}

// Fields of the MixerConfig.
func (MixerConfig) Fields() []ent.Field {
	return nil
}

// Edges of the MixerConfig.
func (MixerConfig) Edges() []ent.Edge {
	return nil
}
