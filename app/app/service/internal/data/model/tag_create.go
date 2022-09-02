// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/app/service/internal/data/model/tag"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TagCreate is the builder for creating a Tag entity.
type TagCreate struct {
	config
	mutation *TagMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (tc *TagCreate) SetCreateTime(t time.Time) *TagCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *TagCreate) SetNillableCreateTime(t *time.Time) *TagCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUpdateTime sets the "update_time" field.
func (tc *TagCreate) SetUpdateTime(t time.Time) *TagCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tc *TagCreate) SetNillableUpdateTime(t *time.Time) *TagCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetDeleteTime sets the "delete_time" field.
func (tc *TagCreate) SetDeleteTime(t time.Time) *TagCreate {
	tc.mutation.SetDeleteTime(t)
	return tc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (tc *TagCreate) SetNillableDeleteTime(t *time.Time) *TagCreate {
	if t != nil {
		tc.SetDeleteTime(*t)
	}
	return tc
}

// SetTitle sets the "title" field.
func (tc *TagCreate) SetTitle(s string) *TagCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TagCreate) SetDescription(s string) *TagCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetHighlight sets the "highlight" field.
func (tc *TagCreate) SetHighlight(i int) *TagCreate {
	tc.mutation.SetHighlight(i)
	return tc
}

// SetType sets the "type" field.
func (tc *TagCreate) SetType(i int) *TagCreate {
	tc.mutation.SetType(i)
	return tc
}

// Mutation returns the TagMutation object of the builder.
func (tc *TagCreate) Mutation() *TagMutation {
	return tc.mutation
}

// Save creates the Tag in the database.
func (tc *TagCreate) Save(ctx context.Context) (*Tag, error) {
	var (
		err  error
		node *Tag
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tag)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TagMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TagCreate) SaveX(ctx context.Context) *Tag {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TagCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TagCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TagCreate) defaults() {
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := tag.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := tag.DefaultUpdateTime()
		tc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TagCreate) check() error {
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "Tag.create_time"`)}
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "Tag.update_time"`)}
	}
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`model: missing required field "Tag.title"`)}
	}
	if _, ok := tc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`model: missing required field "Tag.description"`)}
	}
	if _, ok := tc.mutation.Highlight(); !ok {
		return &ValidationError{Name: "highlight", err: errors.New(`model: missing required field "Tag.highlight"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`model: missing required field "Tag.type"`)}
	}
	return nil
}

func (tc *TagCreate) sqlSave(ctx context.Context) (*Tag, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (tc *TagCreate) createSpec() (*Tag, *sqlgraph.CreateSpec) {
	var (
		_node = &Tag{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: tag.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := tc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := tc.mutation.Highlight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldHighlight,
		})
		_node.Highlight = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldType,
		})
		_node.Type = value
	}
	return _node, _spec
}

// TagCreateBulk is the builder for creating many Tag entities in bulk.
type TagCreateBulk struct {
	config
	builders []*TagCreate
}

// Save creates the Tag entities in the database.
func (tcb *TagCreateBulk) Save(ctx context.Context) ([]*Tag, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tag, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TagMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TagCreateBulk) SaveX(ctx context.Context) []*Tag {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TagCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TagCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
