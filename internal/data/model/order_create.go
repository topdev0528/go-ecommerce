// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/internal/data/model/order"
	"mall-go/internal/data/model/ordersnap"
	"mall-go/internal/data/model/ordersub"
	"mall-go/internal/data/model/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (oc *OrderCreate) SetCreateTime(t time.Time) *OrderCreate {
	oc.mutation.SetCreateTime(t)
	return oc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreateTime(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreateTime(*t)
	}
	return oc
}

// SetUpdateTime sets the "update_time" field.
func (oc *OrderCreate) SetUpdateTime(t time.Time) *OrderCreate {
	oc.mutation.SetUpdateTime(t)
	return oc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdateTime(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdateTime(*t)
	}
	return oc
}

// SetDeleteTime sets the "delete_time" field.
func (oc *OrderCreate) SetDeleteTime(t time.Time) *OrderCreate {
	oc.mutation.SetDeleteTime(t)
	return oc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (oc *OrderCreate) SetNillableDeleteTime(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetDeleteTime(*t)
	}
	return oc
}

// SetOrderNo sets the "order_no" field.
func (oc *OrderCreate) SetOrderNo(s string) *OrderCreate {
	oc.mutation.SetOrderNo(s)
	return oc
}

// SetTransactionID sets the "transaction_id" field.
func (oc *OrderCreate) SetTransactionID(s string) *OrderCreate {
	oc.mutation.SetTransactionID(s)
	return oc
}

// SetUserID sets the "user_id" field.
func (oc *OrderCreate) SetUserID(i int64) *OrderCreate {
	oc.mutation.SetUserID(i)
	return oc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUserID(i *int64) *OrderCreate {
	if i != nil {
		oc.SetUserID(*i)
	}
	return oc
}

// SetTotalPrice sets the "total_price" field.
func (oc *OrderCreate) SetTotalPrice(f float64) *OrderCreate {
	oc.mutation.SetTotalPrice(f)
	return oc
}

// SetTotalCount sets the "total_count" field.
func (oc *OrderCreate) SetTotalCount(i int) *OrderCreate {
	oc.mutation.SetTotalCount(i)
	return oc
}

// SetFinalTotalPrice sets the "final_total_price" field.
func (oc *OrderCreate) SetFinalTotalPrice(f float64) *OrderCreate {
	oc.mutation.SetFinalTotalPrice(f)
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(i int) *OrderCreate {
	oc.mutation.SetStatus(i)
	return oc
}

// SetUser sets the "user" edge to the User entity.
func (oc *OrderCreate) SetUser(u *User) *OrderCreate {
	return oc.SetUserID(u.ID)
}

// AddOrderSnapIDs adds the "order_snap" edge to the OrderSnap entity by IDs.
func (oc *OrderCreate) AddOrderSnapIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddOrderSnapIDs(ids...)
	return oc
}

// AddOrderSnap adds the "order_snap" edges to the OrderSnap entity.
func (oc *OrderCreate) AddOrderSnap(o ...*OrderSnap) *OrderCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddOrderSnapIDs(ids...)
}

// AddOrderSubIDs adds the "order_sub" edge to the OrderSub entity by IDs.
func (oc *OrderCreate) AddOrderSubIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddOrderSubIDs(ids...)
	return oc
}

// AddOrderSub adds the "order_sub" edges to the OrderSub entity.
func (oc *OrderCreate) AddOrderSub(o ...*OrderSub) *OrderCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddOrderSubIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	oc.defaults()
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.CreateTime(); !ok {
		v := order.DefaultCreateTime()
		oc.mutation.SetCreateTime(v)
	}
	if _, ok := oc.mutation.UpdateTime(); !ok {
		v := order.DefaultUpdateTime()
		oc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := oc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := oc.mutation.OrderNo(); !ok {
		return &ValidationError{Name: "order_no", err: errors.New(`model: missing required field "order_no"`)}
	}
	if v, ok := oc.mutation.OrderNo(); ok {
		if err := order.OrderNoValidator(v); err != nil {
			return &ValidationError{Name: "order_no", err: fmt.Errorf(`model: validator failed for field "order_no": %w`, err)}
		}
	}
	if _, ok := oc.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "transaction_id", err: errors.New(`model: missing required field "transaction_id"`)}
	}
	if _, ok := oc.mutation.TotalPrice(); !ok {
		return &ValidationError{Name: "total_price", err: errors.New(`model: missing required field "total_price"`)}
	}
	if _, ok := oc.mutation.TotalCount(); !ok {
		return &ValidationError{Name: "total_count", err: errors.New(`model: missing required field "total_count"`)}
	}
	if _, ok := oc.mutation.FinalTotalPrice(); !ok {
		return &ValidationError{Name: "final_total_price", err: errors.New(`model: missing required field "final_total_price"`)}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`model: missing required field "status"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: order.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: order.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := oc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := oc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := oc.mutation.OrderNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldOrderNo,
		})
		_node.OrderNo = value
	}
	if value, ok := oc.mutation.TransactionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldTransactionID,
		})
		_node.TransactionID = value
	}
	if value, ok := oc.mutation.TotalPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotalPrice,
		})
		_node.TotalPrice = value
	}
	if value, ok := oc.mutation.TotalCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldTotalCount,
		})
		_node.TotalCount = value
	}
	if value, ok := oc.mutation.FinalTotalPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldFinalTotalPrice,
		})
		_node.FinalTotalPrice = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := oc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrderSnapIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSnapTable,
			Columns: []string{order.OrderSnapColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersnap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrderSubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderSubTable,
			Columns: []string{order.OrderSubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: ordersub.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
