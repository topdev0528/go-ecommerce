// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/user/service/internal/data/model/predicate"
	"mall-go/app/user/service/internal/data/model/userwalletdetail"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserWalletDetailUpdate is the builder for updating UserWalletDetail entities.
type UserWalletDetailUpdate struct {
	config
	hooks    []Hook
	mutation *UserWalletDetailMutation
}

// Where appends a list predicates to the UserWalletDetailUpdate builder.
func (uwdu *UserWalletDetailUpdate) Where(ps ...predicate.UserWalletDetail) *UserWalletDetailUpdate {
	uwdu.mutation.Where(ps...)
	return uwdu
}

// SetUpdateTime sets the "update_time" field.
func (uwdu *UserWalletDetailUpdate) SetUpdateTime(t time.Time) *UserWalletDetailUpdate {
	uwdu.mutation.SetUpdateTime(t)
	return uwdu
}

// SetDeleteTime sets the "delete_time" field.
func (uwdu *UserWalletDetailUpdate) SetDeleteTime(t time.Time) *UserWalletDetailUpdate {
	uwdu.mutation.SetDeleteTime(t)
	return uwdu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uwdu *UserWalletDetailUpdate) SetNillableDeleteTime(t *time.Time) *UserWalletDetailUpdate {
	if t != nil {
		uwdu.SetDeleteTime(*t)
	}
	return uwdu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (uwdu *UserWalletDetailUpdate) ClearDeleteTime() *UserWalletDetailUpdate {
	uwdu.mutation.ClearDeleteTime()
	return uwdu
}

// SetDescription sets the "description" field.
func (uwdu *UserWalletDetailUpdate) SetDescription(s string) *UserWalletDetailUpdate {
	uwdu.mutation.SetDescription(s)
	return uwdu
}

// SetOp sets the "op" field.
func (uwdu *UserWalletDetailUpdate) SetOp(i int) *UserWalletDetailUpdate {
	uwdu.mutation.ResetOp()
	uwdu.mutation.SetOp(i)
	return uwdu
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (uwdu *UserWalletDetailUpdate) SetNillableOp(i *int) *UserWalletDetailUpdate {
	if i != nil {
		uwdu.SetOp(*i)
	}
	return uwdu
}

// AddOp adds i to the "op" field.
func (uwdu *UserWalletDetailUpdate) AddOp(i int) *UserWalletDetailUpdate {
	uwdu.mutation.AddOp(i)
	return uwdu
}

// SetCurrent sets the "current" field.
func (uwdu *UserWalletDetailUpdate) SetCurrent(i int) *UserWalletDetailUpdate {
	uwdu.mutation.ResetCurrent()
	uwdu.mutation.SetCurrent(i)
	return uwdu
}

// AddCurrent adds i to the "current" field.
func (uwdu *UserWalletDetailUpdate) AddCurrent(i int) *UserWalletDetailUpdate {
	uwdu.mutation.AddCurrent(i)
	return uwdu
}

// SetValue sets the "value" field.
func (uwdu *UserWalletDetailUpdate) SetValue(i int) *UserWalletDetailUpdate {
	uwdu.mutation.ResetValue()
	uwdu.mutation.SetValue(i)
	return uwdu
}

// AddValue adds i to the "value" field.
func (uwdu *UserWalletDetailUpdate) AddValue(i int) *UserWalletDetailUpdate {
	uwdu.mutation.AddValue(i)
	return uwdu
}

// SetType sets the "type" field.
func (uwdu *UserWalletDetailUpdate) SetType(i int) *UserWalletDetailUpdate {
	uwdu.mutation.ResetType()
	uwdu.mutation.SetType(i)
	return uwdu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (uwdu *UserWalletDetailUpdate) SetNillableType(i *int) *UserWalletDetailUpdate {
	if i != nil {
		uwdu.SetType(*i)
	}
	return uwdu
}

// AddType adds i to the "type" field.
func (uwdu *UserWalletDetailUpdate) AddType(i int) *UserWalletDetailUpdate {
	uwdu.mutation.AddType(i)
	return uwdu
}

// Mutation returns the UserWalletDetailMutation object of the builder.
func (uwdu *UserWalletDetailUpdate) Mutation() *UserWalletDetailMutation {
	return uwdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uwdu *UserWalletDetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uwdu.defaults()
	if len(uwdu.hooks) == 0 {
		affected, err = uwdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWalletDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwdu.mutation = mutation
			affected, err = uwdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uwdu.hooks) - 1; i >= 0; i-- {
			if uwdu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = uwdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uwdu *UserWalletDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := uwdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uwdu *UserWalletDetailUpdate) Exec(ctx context.Context) error {
	_, err := uwdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwdu *UserWalletDetailUpdate) ExecX(ctx context.Context) {
	if err := uwdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwdu *UserWalletDetailUpdate) defaults() {
	if _, ok := uwdu.mutation.UpdateTime(); !ok {
		v := userwalletdetail.UpdateDefaultUpdateTime()
		uwdu.mutation.SetUpdateTime(v)
	}
}

func (uwdu *UserWalletDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwalletdetail.Table,
			Columns: userwalletdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userwalletdetail.FieldID,
			},
		},
	}
	if ps := uwdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uwdu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userwalletdetail.FieldUpdateTime,
		})
	}
	if value, ok := uwdu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userwalletdetail.FieldDeleteTime,
		})
	}
	if uwdu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userwalletdetail.FieldDeleteTime,
		})
	}
	if value, ok := uwdu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userwalletdetail.FieldDescription,
		})
	}
	if value, ok := uwdu.mutation.GetOp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldOp,
		})
	}
	if value, ok := uwdu.mutation.AddedOp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldOp,
		})
	}
	if value, ok := uwdu.mutation.Current(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldCurrent,
		})
	}
	if value, ok := uwdu.mutation.AddedCurrent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldCurrent,
		})
	}
	if value, ok := uwdu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldValue,
		})
	}
	if value, ok := uwdu.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldValue,
		})
	}
	if value, ok := uwdu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldType,
		})
	}
	if value, ok := uwdu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldType,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uwdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userwalletdetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserWalletDetailUpdateOne is the builder for updating a single UserWalletDetail entity.
type UserWalletDetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserWalletDetailMutation
}

// SetUpdateTime sets the "update_time" field.
func (uwduo *UserWalletDetailUpdateOne) SetUpdateTime(t time.Time) *UserWalletDetailUpdateOne {
	uwduo.mutation.SetUpdateTime(t)
	return uwduo
}

// SetDeleteTime sets the "delete_time" field.
func (uwduo *UserWalletDetailUpdateOne) SetDeleteTime(t time.Time) *UserWalletDetailUpdateOne {
	uwduo.mutation.SetDeleteTime(t)
	return uwduo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uwduo *UserWalletDetailUpdateOne) SetNillableDeleteTime(t *time.Time) *UserWalletDetailUpdateOne {
	if t != nil {
		uwduo.SetDeleteTime(*t)
	}
	return uwduo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (uwduo *UserWalletDetailUpdateOne) ClearDeleteTime() *UserWalletDetailUpdateOne {
	uwduo.mutation.ClearDeleteTime()
	return uwduo
}

// SetDescription sets the "description" field.
func (uwduo *UserWalletDetailUpdateOne) SetDescription(s string) *UserWalletDetailUpdateOne {
	uwduo.mutation.SetDescription(s)
	return uwduo
}

// SetOp sets the "op" field.
func (uwduo *UserWalletDetailUpdateOne) SetOp(i int) *UserWalletDetailUpdateOne {
	uwduo.mutation.ResetOp()
	uwduo.mutation.SetOp(i)
	return uwduo
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (uwduo *UserWalletDetailUpdateOne) SetNillableOp(i *int) *UserWalletDetailUpdateOne {
	if i != nil {
		uwduo.SetOp(*i)
	}
	return uwduo
}

// AddOp adds i to the "op" field.
func (uwduo *UserWalletDetailUpdateOne) AddOp(i int) *UserWalletDetailUpdateOne {
	uwduo.mutation.AddOp(i)
	return uwduo
}

// SetCurrent sets the "current" field.
func (uwduo *UserWalletDetailUpdateOne) SetCurrent(i int) *UserWalletDetailUpdateOne {
	uwduo.mutation.ResetCurrent()
	uwduo.mutation.SetCurrent(i)
	return uwduo
}

// AddCurrent adds i to the "current" field.
func (uwduo *UserWalletDetailUpdateOne) AddCurrent(i int) *UserWalletDetailUpdateOne {
	uwduo.mutation.AddCurrent(i)
	return uwduo
}

// SetValue sets the "value" field.
func (uwduo *UserWalletDetailUpdateOne) SetValue(i int) *UserWalletDetailUpdateOne {
	uwduo.mutation.ResetValue()
	uwduo.mutation.SetValue(i)
	return uwduo
}

// AddValue adds i to the "value" field.
func (uwduo *UserWalletDetailUpdateOne) AddValue(i int) *UserWalletDetailUpdateOne {
	uwduo.mutation.AddValue(i)
	return uwduo
}

// SetType sets the "type" field.
func (uwduo *UserWalletDetailUpdateOne) SetType(i int) *UserWalletDetailUpdateOne {
	uwduo.mutation.ResetType()
	uwduo.mutation.SetType(i)
	return uwduo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (uwduo *UserWalletDetailUpdateOne) SetNillableType(i *int) *UserWalletDetailUpdateOne {
	if i != nil {
		uwduo.SetType(*i)
	}
	return uwduo
}

// AddType adds i to the "type" field.
func (uwduo *UserWalletDetailUpdateOne) AddType(i int) *UserWalletDetailUpdateOne {
	uwduo.mutation.AddType(i)
	return uwduo
}

// Mutation returns the UserWalletDetailMutation object of the builder.
func (uwduo *UserWalletDetailUpdateOne) Mutation() *UserWalletDetailMutation {
	return uwduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uwduo *UserWalletDetailUpdateOne) Select(field string, fields ...string) *UserWalletDetailUpdateOne {
	uwduo.fields = append([]string{field}, fields...)
	return uwduo
}

// Save executes the query and returns the updated UserWalletDetail entity.
func (uwduo *UserWalletDetailUpdateOne) Save(ctx context.Context) (*UserWalletDetail, error) {
	var (
		err  error
		node *UserWalletDetail
	)
	uwduo.defaults()
	if len(uwduo.hooks) == 0 {
		node, err = uwduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWalletDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwduo.mutation = mutation
			node, err = uwduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uwduo.hooks) - 1; i >= 0; i-- {
			if uwduo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = uwduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uwduo *UserWalletDetailUpdateOne) SaveX(ctx context.Context) *UserWalletDetail {
	node, err := uwduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uwduo *UserWalletDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := uwduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwduo *UserWalletDetailUpdateOne) ExecX(ctx context.Context) {
	if err := uwduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwduo *UserWalletDetailUpdateOne) defaults() {
	if _, ok := uwduo.mutation.UpdateTime(); !ok {
		v := userwalletdetail.UpdateDefaultUpdateTime()
		uwduo.mutation.SetUpdateTime(v)
	}
}

func (uwduo *UserWalletDetailUpdateOne) sqlSave(ctx context.Context) (_node *UserWalletDetail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwalletdetail.Table,
			Columns: userwalletdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userwalletdetail.FieldID,
			},
		},
	}
	id, ok := uwduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "UserWalletDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uwduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userwalletdetail.FieldID)
		for _, f := range fields {
			if !userwalletdetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != userwalletdetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uwduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uwduo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userwalletdetail.FieldUpdateTime,
		})
	}
	if value, ok := uwduo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userwalletdetail.FieldDeleteTime,
		})
	}
	if uwduo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userwalletdetail.FieldDeleteTime,
		})
	}
	if value, ok := uwduo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userwalletdetail.FieldDescription,
		})
	}
	if value, ok := uwduo.mutation.GetOp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldOp,
		})
	}
	if value, ok := uwduo.mutation.AddedOp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldOp,
		})
	}
	if value, ok := uwduo.mutation.Current(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldCurrent,
		})
	}
	if value, ok := uwduo.mutation.AddedCurrent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldCurrent,
		})
	}
	if value, ok := uwduo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldValue,
		})
	}
	if value, ok := uwduo.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldValue,
		})
	}
	if value, ok := uwduo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldType,
		})
	}
	if value, ok := uwduo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userwalletdetail.FieldType,
		})
	}
	_node = &UserWalletDetail{config: uwduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uwduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userwalletdetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
