// Code generated by entc, DO NOT EDIT.

package ent

import (
	"boschtail/ent/mixerconfig"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MixerConfigCreate is the builder for creating a MixerConfig entity.
type MixerConfigCreate struct {
	config
	mutation *MixerConfigMutation
	hooks    []Hook
}

// Mutation returns the MixerConfigMutation object of the builder.
func (mcc *MixerConfigCreate) Mutation() *MixerConfigMutation {
	return mcc.mutation
}

// Save creates the MixerConfig in the database.
func (mcc *MixerConfigCreate) Save(ctx context.Context) (*MixerConfig, error) {
	var (
		err  error
		node *MixerConfig
	)
	if len(mcc.hooks) == 0 {
		if err = mcc.check(); err != nil {
			return nil, err
		}
		node, err = mcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MixerConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mcc.check(); err != nil {
				return nil, err
			}
			mcc.mutation = mutation
			if node, err = mcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mcc.hooks) - 1; i >= 0; i-- {
			if mcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mcc *MixerConfigCreate) SaveX(ctx context.Context) *MixerConfig {
	v, err := mcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcc *MixerConfigCreate) Exec(ctx context.Context) error {
	_, err := mcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcc *MixerConfigCreate) ExecX(ctx context.Context) {
	if err := mcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcc *MixerConfigCreate) check() error {
	return nil
}

func (mcc *MixerConfigCreate) sqlSave(ctx context.Context) (*MixerConfig, error) {
	_node, _spec := mcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mcc *MixerConfigCreate) createSpec() (*MixerConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &MixerConfig{config: mcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mixerconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mixerconfig.FieldID,
			},
		}
	)
	return _node, _spec
}

// MixerConfigCreateBulk is the builder for creating many MixerConfig entities in bulk.
type MixerConfigCreateBulk struct {
	config
	builders []*MixerConfigCreate
}

// Save creates the MixerConfig entities in the database.
func (mccb *MixerConfigCreateBulk) Save(ctx context.Context) ([]*MixerConfig, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mccb.builders))
	nodes := make([]*MixerConfig, len(mccb.builders))
	mutators := make([]Mutator, len(mccb.builders))
	for i := range mccb.builders {
		func(i int, root context.Context) {
			builder := mccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MixerConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mccb *MixerConfigCreateBulk) SaveX(ctx context.Context) []*MixerConfig {
	v, err := mccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mccb *MixerConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := mccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mccb *MixerConfigCreateBulk) ExecX(ctx context.Context) {
	if err := mccb.Exec(ctx); err != nil {
		panic(err)
	}
}
