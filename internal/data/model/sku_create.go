// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/internal/data/ent/schema"
	"mall-go/internal/data/model/sku"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SkuCreate is the builder for creating a Sku entity.
type SkuCreate struct {
	config
	mutation *SkuMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sc *SkuCreate) SetCreateTime(t time.Time) *SkuCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *SkuCreate) SetNillableCreateTime(t *time.Time) *SkuCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *SkuCreate) SetUpdateTime(t time.Time) *SkuCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *SkuCreate) SetNillableUpdateTime(t *time.Time) *SkuCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetDeleteTime sets the "delete_time" field.
func (sc *SkuCreate) SetDeleteTime(t time.Time) *SkuCreate {
	sc.mutation.SetDeleteTime(t)
	return sc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (sc *SkuCreate) SetNillableDeleteTime(t *time.Time) *SkuCreate {
	if t != nil {
		sc.SetDeleteTime(*t)
	}
	return sc
}

// SetPrice sets the "price" field.
func (sc *SkuCreate) SetPrice(f float64) *SkuCreate {
	sc.mutation.SetPrice(f)
	return sc
}

// SetDiscountPrice sets the "discount_price" field.
func (sc *SkuCreate) SetDiscountPrice(f float64) *SkuCreate {
	sc.mutation.SetDiscountPrice(f)
	return sc
}

// SetOnline sets the "online" field.
func (sc *SkuCreate) SetOnline(i int) *SkuCreate {
	sc.mutation.SetOnline(i)
	return sc
}

// SetImg sets the "img" field.
func (sc *SkuCreate) SetImg(s string) *SkuCreate {
	sc.mutation.SetImg(s)
	return sc
}

// SetTitle sets the "title" field.
func (sc *SkuCreate) SetTitle(s string) *SkuCreate {
	sc.mutation.SetTitle(s)
	return sc
}

// SetSpuID sets the "spu_id" field.
func (sc *SkuCreate) SetSpuID(i int64) *SkuCreate {
	sc.mutation.SetSpuID(i)
	return sc
}

// SetNillableSpuID sets the "spu_id" field if the given value is not nil.
func (sc *SkuCreate) SetNillableSpuID(i *int64) *SkuCreate {
	if i != nil {
		sc.SetSpuID(*i)
	}
	return sc
}

// SetSpecs sets the "specs" field.
func (sc *SkuCreate) SetSpecs(s []schema.Spec) *SkuCreate {
	sc.mutation.SetSpecs(s)
	return sc
}

// SetCode sets the "code" field.
func (sc *SkuCreate) SetCode(s string) *SkuCreate {
	sc.mutation.SetCode(s)
	return sc
}

// SetStock sets the "stock" field.
func (sc *SkuCreate) SetStock(i int) *SkuCreate {
	sc.mutation.SetStock(i)
	return sc
}

// SetCategoryID sets the "category_id" field.
func (sc *SkuCreate) SetCategoryID(i int64) *SkuCreate {
	sc.mutation.SetCategoryID(i)
	return sc
}

// SetRootCategoryID sets the "root_category_id" field.
func (sc *SkuCreate) SetRootCategoryID(i int64) *SkuCreate {
	sc.mutation.SetRootCategoryID(i)
	return sc
}

// Mutation returns the SkuMutation object of the builder.
func (sc *SkuCreate) Mutation() *SkuMutation {
	return sc.mutation
}

// Save creates the Sku in the database.
func (sc *SkuCreate) Save(ctx context.Context) (*Sku, error) {
	var (
		err  error
		node *Sku
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SkuCreate) SaveX(ctx context.Context) *Sku {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SkuCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SkuCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SkuCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := sku.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := sku.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SkuCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := sc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`model: missing required field "price"`)}
	}
	if _, ok := sc.mutation.DiscountPrice(); !ok {
		return &ValidationError{Name: "discount_price", err: errors.New(`model: missing required field "discount_price"`)}
	}
	if _, ok := sc.mutation.Online(); !ok {
		return &ValidationError{Name: "online", err: errors.New(`model: missing required field "online"`)}
	}
	if _, ok := sc.mutation.Img(); !ok {
		return &ValidationError{Name: "img", err: errors.New(`model: missing required field "img"`)}
	}
	if _, ok := sc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`model: missing required field "title"`)}
	}
	if _, ok := sc.mutation.Specs(); !ok {
		return &ValidationError{Name: "specs", err: errors.New(`model: missing required field "specs"`)}
	}
	if _, ok := sc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`model: missing required field "code"`)}
	}
	if _, ok := sc.mutation.Stock(); !ok {
		return &ValidationError{Name: "stock", err: errors.New(`model: missing required field "stock"`)}
	}
	if _, ok := sc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`model: missing required field "category_id"`)}
	}
	if _, ok := sc.mutation.RootCategoryID(); !ok {
		return &ValidationError{Name: "root_category_id", err: errors.New(`model: missing required field "root_category_id"`)}
	}
	return nil
}

func (sc *SkuCreate) sqlSave(ctx context.Context) (*Sku, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (sc *SkuCreate) createSpec() (*Sku, *sqlgraph.CreateSpec) {
	var (
		_node = &Sku{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sku.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sku.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sku.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sku.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sku.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := sc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := sc.mutation.DiscountPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldDiscountPrice,
		})
		_node.DiscountPrice = value
	}
	if value, ok := sc.mutation.Online(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sku.FieldOnline,
		})
		_node.Online = value
	}
	if value, ok := sc.mutation.Img(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sku.FieldImg,
		})
		_node.Img = value
	}
	if value, ok := sc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sku.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := sc.mutation.SpuID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldSpuID,
		})
		_node.SpuID = value
	}
	if value, ok := sc.mutation.Specs(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: sku.FieldSpecs,
		})
		_node.Specs = value
	}
	if value, ok := sc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sku.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := sc.mutation.Stock(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sku.FieldStock,
		})
		_node.Stock = value
	}
	if value, ok := sc.mutation.CategoryID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldCategoryID,
		})
		_node.CategoryID = value
	}
	if value, ok := sc.mutation.RootCategoryID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldRootCategoryID,
		})
		_node.RootCategoryID = value
	}
	return _node, _spec
}

// SkuCreateBulk is the builder for creating many Sku entities in bulk.
type SkuCreateBulk struct {
	config
	builders []*SkuCreate
}

// Save creates the Sku entities in the database.
func (scb *SkuCreateBulk) Save(ctx context.Context) ([]*Sku, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sku, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkuMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SkuCreateBulk) SaveX(ctx context.Context) []*Sku {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SkuCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SkuCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
