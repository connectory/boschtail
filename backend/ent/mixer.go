// Code generated by entc, DO NOT EDIT.

package ent

import (
	"boschtail/ent/mixer"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Mixer is the model entity for the Mixer schema.
type Mixer struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mixer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mixer.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Mixer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mixer fields.
func (m *Mixer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mixer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Mixer.
// Note that you need to call Mixer.Unwrap() before calling this method if this Mixer
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mixer) Update() *MixerUpdateOne {
	return (&MixerClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Mixer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mixer) Unwrap() *Mixer {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mixer is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mixer) String() string {
	var builder strings.Builder
	builder.WriteString("Mixer(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Mixers is a parsable slice of Mixer.
type Mixers []*Mixer

func (m Mixers) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}