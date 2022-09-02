// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/app/service/internal/data/model/gridcategory"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GridCategoryCreate is the builder for creating a GridCategory entity.
type GridCategoryCreate struct {
	config
	mutation *GridCategoryMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (gcc *GridCategoryCreate) SetCreateTime(t time.Time) *GridCategoryCreate {
	gcc.mutation.SetCreateTime(t)
	return gcc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (gcc *GridCategoryCreate) SetNillableCreateTime(t *time.Time) *GridCategoryCreate {
	if t != nil {
		gcc.SetCreateTime(*t)
	}
	return gcc
}

// SetUpdateTime sets the "update_time" field.
func (gcc *GridCategoryCreate) SetUpdateTime(t time.Time) *GridCategoryCreate {
	gcc.mutation.SetUpdateTime(t)
	return gcc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (gcc *GridCategoryCreate) SetNillableUpdateTime(t *time.Time) *GridCategoryCreate {
	if t != nil {
		gcc.SetUpdateTime(*t)
	}
	return gcc
}

// SetDeleteTime sets the "delete_time" field.
func (gcc *GridCategoryCreate) SetDeleteTime(t time.Time) *GridCategoryCreate {
	gcc.mutation.SetDeleteTime(t)
	return gcc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (gcc *GridCategoryCreate) SetNillableDeleteTime(t *time.Time) *GridCategoryCreate {
	if t != nil {
		gcc.SetDeleteTime(*t)
	}
	return gcc
}

// SetTitle sets the "title" field.
func (gcc *GridCategoryCreate) SetTitle(s string) *GridCategoryCreate {
	gcc.mutation.SetTitle(s)
	return gcc
}

// SetImg sets the "img" field.
func (gcc *GridCategoryCreate) SetImg(s string) *GridCategoryCreate {
	gcc.mutation.SetImg(s)
	return gcc
}

// SetName sets the "name" field.
func (gcc *GridCategoryCreate) SetName(s string) *GridCategoryCreate {
	gcc.mutation.SetName(s)
	return gcc
}

// SetCategoryID sets the "category_id" field.
func (gcc *GridCategoryCreate) SetCategoryID(i int) *GridCategoryCreate {
	gcc.mutation.SetCategoryID(i)
	return gcc
}

// SetRootCategoryID sets the "root_category_id" field.
func (gcc *GridCategoryCreate) SetRootCategoryID(i int) *GridCategoryCreate {
	gcc.mutation.SetRootCategoryID(i)
	return gcc
}

// Mutation returns the GridCategoryMutation object of the builder.
func (gcc *GridCategoryCreate) Mutation() *GridCategoryMutation {
	return gcc.mutation
}

// Save creates the GridCategory in the database.
func (gcc *GridCategoryCreate) Save(ctx context.Context) (*GridCategory, error) {
	var (
		err  error
		node *GridCategory
	)
	gcc.defaults()
	if len(gcc.hooks) == 0 {
		if err = gcc.check(); err != nil {
			return nil, err
		}
		node, err = gcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GridCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gcc.check(); err != nil {
				return nil, err
			}
			gcc.mutation = mutation
			if node, err = gcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gcc.hooks) - 1; i >= 0; i-- {
			if gcc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = gcc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gcc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GridCategory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GridCategoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gcc *GridCategoryCreate) SaveX(ctx context.Context) *GridCategory {
	v, err := gcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcc *GridCategoryCreate) Exec(ctx context.Context) error {
	_, err := gcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcc *GridCategoryCreate) ExecX(ctx context.Context) {
	if err := gcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gcc *GridCategoryCreate) defaults() {
	if _, ok := gcc.mutation.CreateTime(); !ok {
		v := gridcategory.DefaultCreateTime()
		gcc.mutation.SetCreateTime(v)
	}
	if _, ok := gcc.mutation.UpdateTime(); !ok {
		v := gridcategory.DefaultUpdateTime()
		gcc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gcc *GridCategoryCreate) check() error {
	if _, ok := gcc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "GridCategory.create_time"`)}
	}
	if _, ok := gcc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "GridCategory.update_time"`)}
	}
	if _, ok := gcc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`model: missing required field "GridCategory.title"`)}
	}
	if _, ok := gcc.mutation.Img(); !ok {
		return &ValidationError{Name: "img", err: errors.New(`model: missing required field "GridCategory.img"`)}
	}
	if _, ok := gcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "GridCategory.name"`)}
	}
	if _, ok := gcc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`model: missing required field "GridCategory.category_id"`)}
	}
	if _, ok := gcc.mutation.RootCategoryID(); !ok {
		return &ValidationError{Name: "root_category_id", err: errors.New(`model: missing required field "GridCategory.root_category_id"`)}
	}
	return nil
}

func (gcc *GridCategoryCreate) sqlSave(ctx context.Context) (*GridCategory, error) {
	_node, _spec := gcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (gcc *GridCategoryCreate) createSpec() (*GridCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &GridCategory{config: gcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gridcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: gridcategory.FieldID,
			},
		}
	)
	if value, ok := gcc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gridcategory.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := gcc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gridcategory.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := gcc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gridcategory.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := gcc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gridcategory.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := gcc.mutation.Img(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gridcategory.FieldImg,
		})
		_node.Img = value
	}
	if value, ok := gcc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gridcategory.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gcc.mutation.CategoryID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldCategoryID,
		})
		_node.CategoryID = value
	}
	if value, ok := gcc.mutation.RootCategoryID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: gridcategory.FieldRootCategoryID,
		})
		_node.RootCategoryID = value
	}
	return _node, _spec
}

// GridCategoryCreateBulk is the builder for creating many GridCategory entities in bulk.
type GridCategoryCreateBulk struct {
	config
	builders []*GridCategoryCreate
}

// Save creates the GridCategory entities in the database.
func (gccb *GridCategoryCreateBulk) Save(ctx context.Context) ([]*GridCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gccb.builders))
	nodes := make([]*GridCategory, len(gccb.builders))
	mutators := make([]Mutator, len(gccb.builders))
	for i := range gccb.builders {
		func(i int, root context.Context) {
			builder := gccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GridCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, gccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gccb *GridCategoryCreateBulk) SaveX(ctx context.Context) []*GridCategory {
	v, err := gccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gccb *GridCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := gccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gccb *GridCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := gccb.Exec(ctx); err != nil {
		panic(err)
	}
}
