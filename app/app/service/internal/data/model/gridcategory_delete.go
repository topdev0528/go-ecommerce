// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/app/app/service/internal/data/model/gridcategory"
	"mall-go/app/app/service/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GridCategoryDelete is the builder for deleting a GridCategory entity.
type GridCategoryDelete struct {
	config
	hooks    []Hook
	mutation *GridCategoryMutation
}

// Where appends a list predicates to the GridCategoryDelete builder.
func (gcd *GridCategoryDelete) Where(ps ...predicate.GridCategory) *GridCategoryDelete {
	gcd.mutation.Where(ps...)
	return gcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gcd *GridCategoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gcd.hooks) == 0 {
		affected, err = gcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GridCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcd.mutation = mutation
			affected, err = gcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gcd.hooks) - 1; i >= 0; i-- {
			if gcd.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = gcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcd *GridCategoryDelete) ExecX(ctx context.Context) int {
	n, err := gcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gcd *GridCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: gridcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: gridcategory.FieldID,
			},
		},
	}
	if ps := gcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GridCategoryDeleteOne is the builder for deleting a single GridCategory entity.
type GridCategoryDeleteOne struct {
	gcd *GridCategoryDelete
}

// Exec executes the deletion query.
func (gcdo *GridCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := gcdo.gcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gridcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gcdo *GridCategoryDeleteOne) ExecX(ctx context.Context) {
	gcdo.gcd.ExecX(ctx)
}
