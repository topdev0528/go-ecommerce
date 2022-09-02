// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/app/user/service/internal/data/model/predicate"
	"mall-go/app/user/service/internal/data/model/userpoint"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPointDelete is the builder for deleting a UserPoint entity.
type UserPointDelete struct {
	config
	hooks    []Hook
	mutation *UserPointMutation
}

// Where appends a list predicates to the UserPointDelete builder.
func (upd *UserPointDelete) Where(ps ...predicate.UserPoint) *UserPointDelete {
	upd.mutation.Where(ps...)
	return upd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (upd *UserPointDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(upd.hooks) == 0 {
		affected, err = upd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserPointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			upd.mutation = mutation
			affected, err = upd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upd.hooks) - 1; i >= 0; i-- {
			if upd.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = upd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (upd *UserPointDelete) ExecX(ctx context.Context) int {
	n, err := upd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (upd *UserPointDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userpoint.FieldID,
			},
		},
	}
	if ps := upd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, upd.driver, _spec)
}

// UserPointDeleteOne is the builder for deleting a single UserPoint entity.
type UserPointDeleteOne struct {
	upd *UserPointDelete
}

// Exec executes the deletion query.
func (updo *UserPointDeleteOne) Exec(ctx context.Context) error {
	n, err := updo.upd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userpoint.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (updo *UserPointDeleteOne) ExecX(ctx context.Context) {
	updo.upd.ExecX(ctx)
}
