// Code generated by entc, DO NOT EDIT.

package ent

import (
	"boschtail/ent/predicate"
	"boschtail/ent/recipeingredient"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecipeIngredientUpdate is the builder for updating RecipeIngredient entities.
type RecipeIngredientUpdate struct {
	config
	hooks    []Hook
	mutation *RecipeIngredientMutation
}

// Where appends a list predicates to the RecipeIngredientUpdate builder.
func (riu *RecipeIngredientUpdate) Where(ps ...predicate.RecipeIngredient) *RecipeIngredientUpdate {
	riu.mutation.Where(ps...)
	return riu
}

// Mutation returns the RecipeIngredientMutation object of the builder.
func (riu *RecipeIngredientUpdate) Mutation() *RecipeIngredientMutation {
	return riu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (riu *RecipeIngredientUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(riu.hooks) == 0 {
		affected, err = riu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipeIngredientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			riu.mutation = mutation
			affected, err = riu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(riu.hooks) - 1; i >= 0; i-- {
			if riu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = riu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, riu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (riu *RecipeIngredientUpdate) SaveX(ctx context.Context) int {
	affected, err := riu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (riu *RecipeIngredientUpdate) Exec(ctx context.Context) error {
	_, err := riu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riu *RecipeIngredientUpdate) ExecX(ctx context.Context) {
	if err := riu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (riu *RecipeIngredientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recipeingredient.Table,
			Columns: recipeingredient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recipeingredient.FieldID,
			},
		},
	}
	if ps := riu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, riu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipeingredient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RecipeIngredientUpdateOne is the builder for updating a single RecipeIngredient entity.
type RecipeIngredientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecipeIngredientMutation
}

// Mutation returns the RecipeIngredientMutation object of the builder.
func (riuo *RecipeIngredientUpdateOne) Mutation() *RecipeIngredientMutation {
	return riuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (riuo *RecipeIngredientUpdateOne) Select(field string, fields ...string) *RecipeIngredientUpdateOne {
	riuo.fields = append([]string{field}, fields...)
	return riuo
}

// Save executes the query and returns the updated RecipeIngredient entity.
func (riuo *RecipeIngredientUpdateOne) Save(ctx context.Context) (*RecipeIngredient, error) {
	var (
		err  error
		node *RecipeIngredient
	)
	if len(riuo.hooks) == 0 {
		node, err = riuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipeIngredientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			riuo.mutation = mutation
			node, err = riuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(riuo.hooks) - 1; i >= 0; i-- {
			if riuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = riuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, riuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (riuo *RecipeIngredientUpdateOne) SaveX(ctx context.Context) *RecipeIngredient {
	node, err := riuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (riuo *RecipeIngredientUpdateOne) Exec(ctx context.Context) error {
	_, err := riuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riuo *RecipeIngredientUpdateOne) ExecX(ctx context.Context) {
	if err := riuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (riuo *RecipeIngredientUpdateOne) sqlSave(ctx context.Context) (_node *RecipeIngredient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recipeingredient.Table,
			Columns: recipeingredient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recipeingredient.FieldID,
			},
		},
	}
	id, ok := riuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RecipeIngredient.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := riuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recipeingredient.FieldID)
		for _, f := range fields {
			if !recipeingredient.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != recipeingredient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := riuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &RecipeIngredient{config: riuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, riuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipeingredient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}