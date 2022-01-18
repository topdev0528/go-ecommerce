// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/sku/service/internal/data/model/orderdetail"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderDetailCreate is the builder for creating a OrderDetail entity.
type OrderDetailCreate struct {
	config
	mutation *OrderDetailMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (odc *OrderDetailCreate) SetCreateTime(t time.Time) *OrderDetailCreate {
	odc.mutation.SetCreateTime(t)
	return odc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillableCreateTime(t *time.Time) *OrderDetailCreate {
	if t != nil {
		odc.SetCreateTime(*t)
	}
	return odc
}

// SetUpdateTime sets the "update_time" field.
func (odc *OrderDetailCreate) SetUpdateTime(t time.Time) *OrderDetailCreate {
	odc.mutation.SetUpdateTime(t)
	return odc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillableUpdateTime(t *time.Time) *OrderDetailCreate {
	if t != nil {
		odc.SetUpdateTime(*t)
	}
	return odc
}

// SetDeleteTime sets the "delete_time" field.
func (odc *OrderDetailCreate) SetDeleteTime(t time.Time) *OrderDetailCreate {
	odc.mutation.SetDeleteTime(t)
	return odc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillableDeleteTime(t *time.Time) *OrderDetailCreate {
	if t != nil {
		odc.SetDeleteTime(*t)
	}
	return odc
}

// SetUserID sets the "user_id" field.
func (odc *OrderDetailCreate) SetUserID(i int64) *OrderDetailCreate {
	odc.mutation.SetUserID(i)
	return odc
}

// SetPayWay sets the "pay_way" field.
func (odc *OrderDetailCreate) SetPayWay(i int) *OrderDetailCreate {
	odc.mutation.SetPayWay(i)
	return odc
}

// SetNillablePayWay sets the "pay_way" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillablePayWay(i *int) *OrderDetailCreate {
	if i != nil {
		odc.SetPayWay(*i)
	}
	return odc
}

// SetClientType sets the "client_type" field.
func (odc *OrderDetailCreate) SetClientType(i int) *OrderDetailCreate {
	odc.mutation.SetClientType(i)
	return odc
}

// SetNillableClientType sets the "client_type" field if the given value is not nil.
func (odc *OrderDetailCreate) SetNillableClientType(i *int) *OrderDetailCreate {
	if i != nil {
		odc.SetClientType(*i)
	}
	return odc
}

// SetShipNo sets the "ship_no" field.
func (odc *OrderDetailCreate) SetShipNo(s string) *OrderDetailCreate {
	odc.mutation.SetShipNo(s)
	return odc
}

// Mutation returns the OrderDetailMutation object of the builder.
func (odc *OrderDetailCreate) Mutation() *OrderDetailMutation {
	return odc.mutation
}

// Save creates the OrderDetail in the database.
func (odc *OrderDetailCreate) Save(ctx context.Context) (*OrderDetail, error) {
	var (
		err  error
		node *OrderDetail
	)
	odc.defaults()
	if len(odc.hooks) == 0 {
		if err = odc.check(); err != nil {
			return nil, err
		}
		node, err = odc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = odc.check(); err != nil {
				return nil, err
			}
			odc.mutation = mutation
			if node, err = odc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(odc.hooks) - 1; i >= 0; i-- {
			if odc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = odc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (odc *OrderDetailCreate) SaveX(ctx context.Context) *OrderDetail {
	v, err := odc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (odc *OrderDetailCreate) Exec(ctx context.Context) error {
	_, err := odc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odc *OrderDetailCreate) ExecX(ctx context.Context) {
	if err := odc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (odc *OrderDetailCreate) defaults() {
	if _, ok := odc.mutation.CreateTime(); !ok {
		v := orderdetail.DefaultCreateTime()
		odc.mutation.SetCreateTime(v)
	}
	if _, ok := odc.mutation.UpdateTime(); !ok {
		v := orderdetail.DefaultUpdateTime()
		odc.mutation.SetUpdateTime(v)
	}
	if _, ok := odc.mutation.PayWay(); !ok {
		v := orderdetail.DefaultPayWay
		odc.mutation.SetPayWay(v)
	}
	if _, ok := odc.mutation.ClientType(); !ok {
		v := orderdetail.DefaultClientType
		odc.mutation.SetClientType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (odc *OrderDetailCreate) check() error {
	if _, ok := odc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := odc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := odc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`model: missing required field "user_id"`)}
	}
	if _, ok := odc.mutation.PayWay(); !ok {
		return &ValidationError{Name: "pay_way", err: errors.New(`model: missing required field "pay_way"`)}
	}
	if _, ok := odc.mutation.ClientType(); !ok {
		return &ValidationError{Name: "client_type", err: errors.New(`model: missing required field "client_type"`)}
	}
	if _, ok := odc.mutation.ShipNo(); !ok {
		return &ValidationError{Name: "ship_no", err: errors.New(`model: missing required field "ship_no"`)}
	}
	return nil
}

func (odc *OrderDetailCreate) sqlSave(ctx context.Context) (*OrderDetail, error) {
	_node, _spec := odc.createSpec()
	if err := sqlgraph.CreateNode(ctx, odc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (odc *OrderDetailCreate) createSpec() (*OrderDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderDetail{config: odc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: orderdetail.FieldID,
			},
		}
	)
	if value, ok := odc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := odc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := odc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderdetail.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := odc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderdetail.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := odc.mutation.PayWay(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldPayWay,
		})
		_node.PayWay = value
	}
	if value, ok := odc.mutation.ClientType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderdetail.FieldClientType,
		})
		_node.ClientType = value
	}
	if value, ok := odc.mutation.ShipNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderdetail.FieldShipNo,
		})
		_node.ShipNo = value
	}
	return _node, _spec
}

// OrderDetailCreateBulk is the builder for creating many OrderDetail entities in bulk.
type OrderDetailCreateBulk struct {
	config
	builders []*OrderDetailCreate
}

// Save creates the OrderDetail entities in the database.
func (odcb *OrderDetailCreateBulk) Save(ctx context.Context) ([]*OrderDetail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(odcb.builders))
	nodes := make([]*OrderDetail, len(odcb.builders))
	mutators := make([]Mutator, len(odcb.builders))
	for i := range odcb.builders {
		func(i int, root context.Context) {
			builder := odcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderDetailMutation)
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
					_, err = mutators[i+1].Mutate(root, odcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, odcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, odcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (odcb *OrderDetailCreateBulk) SaveX(ctx context.Context) []*OrderDetail {
	v, err := odcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (odcb *OrderDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := odcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odcb *OrderDetailCreateBulk) ExecX(ctx context.Context) {
	if err := odcb.Exec(ctx); err != nil {
		panic(err)
	}
}