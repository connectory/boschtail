// Code generated by entc, DO NOT EDIT.

package ent

import (
	"boschtail/ent/recipeingredient"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecipeIngredientCreate is the builder for creating a RecipeIngredient entity.
type RecipeIngredientCreate struct {
	config
	mutation *RecipeIngredientMutation
	hooks    []Hook
}

// Mutation returns the RecipeIngredientMutation object of the builder.
func (ric *RecipeIngredientCreate) Mutation() *RecipeIngredientMutation {
	return ric.mutation
}

// Save creates the RecipeIngredient in the database.
func (ric *RecipeIngredientCreate) Save(ctx context.Context) (*RecipeIngredient, error) {
	var (
		err  error
		node *RecipeIngredient
	)
	if len(ric.hooks) == 0 {
		if err = ric.check(); err != nil {
			return nil, err
		}
		node, err = ric.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipeIngredientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ric.check(); err != nil {
				return nil, err
			}
			ric.mutation = mutation
			if node, err = ric.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ric.hooks) - 1; i >= 0; i-- {
			if ric.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ric.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ric.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ric *RecipeIngredientCreate) SaveX(ctx context.Context) *RecipeIngredient {
	v, err := ric.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ric *RecipeIngredientCreate) Exec(ctx context.Context) error {
	_, err := ric.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ric *RecipeIngredientCreate) ExecX(ctx context.Context) {
	if err := ric.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ric *RecipeIngredientCreate) check() error {
	return nil
}

func (ric *RecipeIngredientCreate) sqlSave(ctx context.Context) (*RecipeIngredient, error) {
	_node, _spec := ric.createSpec()
	if err := sqlgraph.CreateNode(ctx, ric.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ric *RecipeIngredientCreate) createSpec() (*RecipeIngredient, *sqlgraph.CreateSpec) {
	var (
		_node = &RecipeIngredient{config: ric.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recipeingredient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recipeingredient.FieldID,
			},
		}
	)
	return _node, _spec
}

// RecipeIngredientCreateBulk is the builder for creating many RecipeIngredient entities in bulk.
type RecipeIngredientCreateBulk struct {
	config
	builders []*RecipeIngredientCreate
}

// Save creates the RecipeIngredient entities in the database.
func (ricb *RecipeIngredientCreateBulk) Save(ctx context.Context) ([]*RecipeIngredient, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ricb.builders))
	nodes := make([]*RecipeIngredient, len(ricb.builders))
	mutators := make([]Mutator, len(ricb.builders))
	for i := range ricb.builders {
		func(i int, root context.Context) {
			builder := ricb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecipeIngredientMutation)
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
					_, err = mutators[i+1].Mutate(root, ricb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ricb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ricb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ricb *RecipeIngredientCreateBulk) SaveX(ctx context.Context) []*RecipeIngredient {
	v, err := ricb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ricb *RecipeIngredientCreateBulk) Exec(ctx context.Context) error {
	_, err := ricb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ricb *RecipeIngredientCreateBulk) ExecX(ctx context.Context) {
	if err := ricb.Exec(ctx); err != nil {
		panic(err)
	}
}
