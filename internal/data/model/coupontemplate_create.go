// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/internal/data/model/coupontemplate"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CouponTemplateCreate is the builder for creating a CouponTemplate entity.
type CouponTemplateCreate struct {
	config
	mutation *CouponTemplateMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ctc *CouponTemplateCreate) SetCreateTime(t time.Time) *CouponTemplateCreate {
	ctc.mutation.SetCreateTime(t)
	return ctc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ctc *CouponTemplateCreate) SetNillableCreateTime(t *time.Time) *CouponTemplateCreate {
	if t != nil {
		ctc.SetCreateTime(*t)
	}
	return ctc
}

// SetUpdateTime sets the "update_time" field.
func (ctc *CouponTemplateCreate) SetUpdateTime(t time.Time) *CouponTemplateCreate {
	ctc.mutation.SetUpdateTime(t)
	return ctc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ctc *CouponTemplateCreate) SetNillableUpdateTime(t *time.Time) *CouponTemplateCreate {
	if t != nil {
		ctc.SetUpdateTime(*t)
	}
	return ctc
}

// SetDeleteTime sets the "delete_time" field.
func (ctc *CouponTemplateCreate) SetDeleteTime(t time.Time) *CouponTemplateCreate {
	ctc.mutation.SetDeleteTime(t)
	return ctc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ctc *CouponTemplateCreate) SetNillableDeleteTime(t *time.Time) *CouponTemplateCreate {
	if t != nil {
		ctc.SetDeleteTime(*t)
	}
	return ctc
}

// SetTitle sets the "title" field.
func (ctc *CouponTemplateCreate) SetTitle(s string) *CouponTemplateCreate {
	ctc.mutation.SetTitle(s)
	return ctc
}

// SetDescription sets the "description" field.
func (ctc *CouponTemplateCreate) SetDescription(s string) *CouponTemplateCreate {
	ctc.mutation.SetDescription(s)
	return ctc
}

// SetFullMoney sets the "full_money" field.
func (ctc *CouponTemplateCreate) SetFullMoney(f float64) *CouponTemplateCreate {
	ctc.mutation.SetFullMoney(f)
	return ctc
}

// SetMinus sets the "minus" field.
func (ctc *CouponTemplateCreate) SetMinus(f float64) *CouponTemplateCreate {
	ctc.mutation.SetMinus(f)
	return ctc
}

// SetDiscount sets the "discount" field.
func (ctc *CouponTemplateCreate) SetDiscount(f float64) *CouponTemplateCreate {
	ctc.mutation.SetDiscount(f)
	return ctc
}

// SetType sets the "type" field.
func (ctc *CouponTemplateCreate) SetType(i int) *CouponTemplateCreate {
	ctc.mutation.SetType(i)
	return ctc
}

// Mutation returns the CouponTemplateMutation object of the builder.
func (ctc *CouponTemplateCreate) Mutation() *CouponTemplateMutation {
	return ctc.mutation
}

// Save creates the CouponTemplate in the database.
func (ctc *CouponTemplateCreate) Save(ctx context.Context) (*CouponTemplate, error) {
	var (
		err  error
		node *CouponTemplate
	)
	ctc.defaults()
	if len(ctc.hooks) == 0 {
		if err = ctc.check(); err != nil {
			return nil, err
		}
		node, err = ctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctc.check(); err != nil {
				return nil, err
			}
			ctc.mutation = mutation
			if node, err = ctc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ctc.hooks) - 1; i >= 0; i-- {
			if ctc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ctc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *CouponTemplateCreate) SaveX(ctx context.Context) *CouponTemplate {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *CouponTemplateCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *CouponTemplateCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctc *CouponTemplateCreate) defaults() {
	if _, ok := ctc.mutation.CreateTime(); !ok {
		v := coupontemplate.DefaultCreateTime()
		ctc.mutation.SetCreateTime(v)
	}
	if _, ok := ctc.mutation.UpdateTime(); !ok {
		v := coupontemplate.DefaultUpdateTime()
		ctc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *CouponTemplateCreate) check() error {
	if _, ok := ctc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := ctc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := ctc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`model: missing required field "title"`)}
	}
	if _, ok := ctc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`model: missing required field "description"`)}
	}
	if _, ok := ctc.mutation.FullMoney(); !ok {
		return &ValidationError{Name: "full_money", err: errors.New(`model: missing required field "full_money"`)}
	}
	if _, ok := ctc.mutation.Minus(); !ok {
		return &ValidationError{Name: "minus", err: errors.New(`model: missing required field "minus"`)}
	}
	if _, ok := ctc.mutation.Discount(); !ok {
		return &ValidationError{Name: "discount", err: errors.New(`model: missing required field "discount"`)}
	}
	if _, ok := ctc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`model: missing required field "type"`)}
	}
	return nil
}

func (ctc *CouponTemplateCreate) sqlSave(ctx context.Context) (*CouponTemplate, error) {
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (ctc *CouponTemplateCreate) createSpec() (*CouponTemplate, *sqlgraph.CreateSpec) {
	var (
		_node = &CouponTemplate{config: ctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coupontemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: coupontemplate.FieldID,
			},
		}
	)
	if value, ok := ctc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupontemplate.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ctc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupontemplate.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ctc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coupontemplate.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := ctc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontemplate.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ctc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coupontemplate.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := ctc.mutation.FullMoney(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldFullMoney,
		})
		_node.FullMoney = value
	}
	if value, ok := ctc.mutation.Minus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldMinus,
		})
		_node.Minus = value
	}
	if value, ok := ctc.mutation.Discount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: coupontemplate.FieldDiscount,
		})
		_node.Discount = value
	}
	if value, ok := ctc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coupontemplate.FieldType,
		})
		_node.Type = value
	}
	return _node, _spec
}

// CouponTemplateCreateBulk is the builder for creating many CouponTemplate entities in bulk.
type CouponTemplateCreateBulk struct {
	config
	builders []*CouponTemplateCreate
}

// Save creates the CouponTemplate entities in the database.
func (ctcb *CouponTemplateCreateBulk) Save(ctx context.Context) ([]*CouponTemplate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*CouponTemplate, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CouponTemplateMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *CouponTemplateCreateBulk) SaveX(ctx context.Context) []*CouponTemplate {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *CouponTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *CouponTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}
