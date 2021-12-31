// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/internal/data/ent/schema"
	"mall-go/internal/data/model/predicate"
	"mall-go/internal/data/model/sku"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SkuUpdate is the builder for updating Sku entities.
type SkuUpdate struct {
	config
	hooks    []Hook
	mutation *SkuMutation
}

// Where appends a list predicates to the SkuUpdate builder.
func (su *SkuUpdate) Where(ps ...predicate.Sku) *SkuUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SkuUpdate) SetUpdateTime(t time.Time) *SkuUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetDeleteTime sets the "delete_time" field.
func (su *SkuUpdate) SetDeleteTime(t time.Time) *SkuUpdate {
	su.mutation.SetDeleteTime(t)
	return su
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (su *SkuUpdate) SetNillableDeleteTime(t *time.Time) *SkuUpdate {
	if t != nil {
		su.SetDeleteTime(*t)
	}
	return su
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (su *SkuUpdate) ClearDeleteTime() *SkuUpdate {
	su.mutation.ClearDeleteTime()
	return su
}

// SetPrice sets the "price" field.
func (su *SkuUpdate) SetPrice(f float64) *SkuUpdate {
	su.mutation.ResetPrice()
	su.mutation.SetPrice(f)
	return su
}

// AddPrice adds f to the "price" field.
func (su *SkuUpdate) AddPrice(f float64) *SkuUpdate {
	su.mutation.AddPrice(f)
	return su
}

// SetDiscountPrice sets the "discount_price" field.
func (su *SkuUpdate) SetDiscountPrice(f float64) *SkuUpdate {
	su.mutation.ResetDiscountPrice()
	su.mutation.SetDiscountPrice(f)
	return su
}

// AddDiscountPrice adds f to the "discount_price" field.
func (su *SkuUpdate) AddDiscountPrice(f float64) *SkuUpdate {
	su.mutation.AddDiscountPrice(f)
	return su
}

// SetOnline sets the "online" field.
func (su *SkuUpdate) SetOnline(i int8) *SkuUpdate {
	su.mutation.ResetOnline()
	su.mutation.SetOnline(i)
	return su
}

// AddOnline adds i to the "online" field.
func (su *SkuUpdate) AddOnline(i int8) *SkuUpdate {
	su.mutation.AddOnline(i)
	return su
}

// SetImg sets the "img" field.
func (su *SkuUpdate) SetImg(s string) *SkuUpdate {
	su.mutation.SetImg(s)
	return su
}

// SetTitle sets the "title" field.
func (su *SkuUpdate) SetTitle(s string) *SkuUpdate {
	su.mutation.SetTitle(s)
	return su
}

// SetSpuID sets the "spu_id" field.
func (su *SkuUpdate) SetSpuID(i int64) *SkuUpdate {
	su.mutation.ResetSpuID()
	su.mutation.SetSpuID(i)
	return su
}

// SetNillableSpuID sets the "spu_id" field if the given value is not nil.
func (su *SkuUpdate) SetNillableSpuID(i *int64) *SkuUpdate {
	if i != nil {
		su.SetSpuID(*i)
	}
	return su
}

// AddSpuID adds i to the "spu_id" field.
func (su *SkuUpdate) AddSpuID(i int64) *SkuUpdate {
	su.mutation.AddSpuID(i)
	return su
}

// ClearSpuID clears the value of the "spu_id" field.
func (su *SkuUpdate) ClearSpuID() *SkuUpdate {
	su.mutation.ClearSpuID()
	return su
}

// SetSpecs sets the "specs" field.
func (su *SkuUpdate) SetSpecs(s []schema.Spec) *SkuUpdate {
	su.mutation.SetSpecs(s)
	return su
}

// SetCode sets the "code" field.
func (su *SkuUpdate) SetCode(s string) *SkuUpdate {
	su.mutation.SetCode(s)
	return su
}

// SetStock sets the "stock" field.
func (su *SkuUpdate) SetStock(i int) *SkuUpdate {
	su.mutation.ResetStock()
	su.mutation.SetStock(i)
	return su
}

// AddStock adds i to the "stock" field.
func (su *SkuUpdate) AddStock(i int) *SkuUpdate {
	su.mutation.AddStock(i)
	return su
}

// SetCategoryID sets the "category_id" field.
func (su *SkuUpdate) SetCategoryID(i int64) *SkuUpdate {
	su.mutation.ResetCategoryID()
	su.mutation.SetCategoryID(i)
	return su
}

// AddCategoryID adds i to the "category_id" field.
func (su *SkuUpdate) AddCategoryID(i int64) *SkuUpdate {
	su.mutation.AddCategoryID(i)
	return su
}

// SetRootCategoryID sets the "root_category_id" field.
func (su *SkuUpdate) SetRootCategoryID(i int64) *SkuUpdate {
	su.mutation.ResetRootCategoryID()
	su.mutation.SetRootCategoryID(i)
	return su
}

// AddRootCategoryID adds i to the "root_category_id" field.
func (su *SkuUpdate) AddRootCategoryID(i int64) *SkuUpdate {
	su.mutation.AddRootCategoryID(i)
	return su
}

// Mutation returns the SkuMutation object of the builder.
func (su *SkuUpdate) Mutation() *SkuMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SkuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SkuUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SkuUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SkuUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SkuUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := sku.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

func (su *SkuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sku.Table,
			Columns: sku.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sku.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sku.FieldUpdateTime,
		})
	}
	if value, ok := su.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sku.FieldDeleteTime,
		})
	}
	if su.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sku.FieldDeleteTime,
		})
	}
	if value, ok := su.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldPrice,
		})
	}
	if value, ok := su.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldPrice,
		})
	}
	if value, ok := su.mutation.DiscountPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldDiscountPrice,
		})
	}
	if value, ok := su.mutation.AddedDiscountPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldDiscountPrice,
		})
	}
	if value, ok := su.mutation.Online(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: sku.FieldOnline,
		})
	}
	if value, ok := su.mutation.AddedOnline(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: sku.FieldOnline,
		})
	}
	if value, ok := su.mutation.Img(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sku.FieldImg,
		})
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sku.FieldTitle,
		})
	}
	if value, ok := su.mutation.SpuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldSpuID,
		})
	}
	if value, ok := su.mutation.AddedSpuID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldSpuID,
		})
	}
	if su.mutation.SpuIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: sku.FieldSpuID,
		})
	}
	if value, ok := su.mutation.Specs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: sku.FieldSpecs,
		})
	}
	if value, ok := su.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sku.FieldCode,
		})
	}
	if value, ok := su.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sku.FieldStock,
		})
	}
	if value, ok := su.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sku.FieldStock,
		})
	}
	if value, ok := su.mutation.CategoryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldCategoryID,
		})
	}
	if value, ok := su.mutation.AddedCategoryID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldCategoryID,
		})
	}
	if value, ok := su.mutation.RootCategoryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldRootCategoryID,
		})
	}
	if value, ok := su.mutation.AddedRootCategoryID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldRootCategoryID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sku.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SkuUpdateOne is the builder for updating a single Sku entity.
type SkuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SkuMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *SkuUpdateOne) SetUpdateTime(t time.Time) *SkuUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetDeleteTime sets the "delete_time" field.
func (suo *SkuUpdateOne) SetDeleteTime(t time.Time) *SkuUpdateOne {
	suo.mutation.SetDeleteTime(t)
	return suo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (suo *SkuUpdateOne) SetNillableDeleteTime(t *time.Time) *SkuUpdateOne {
	if t != nil {
		suo.SetDeleteTime(*t)
	}
	return suo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (suo *SkuUpdateOne) ClearDeleteTime() *SkuUpdateOne {
	suo.mutation.ClearDeleteTime()
	return suo
}

// SetPrice sets the "price" field.
func (suo *SkuUpdateOne) SetPrice(f float64) *SkuUpdateOne {
	suo.mutation.ResetPrice()
	suo.mutation.SetPrice(f)
	return suo
}

// AddPrice adds f to the "price" field.
func (suo *SkuUpdateOne) AddPrice(f float64) *SkuUpdateOne {
	suo.mutation.AddPrice(f)
	return suo
}

// SetDiscountPrice sets the "discount_price" field.
func (suo *SkuUpdateOne) SetDiscountPrice(f float64) *SkuUpdateOne {
	suo.mutation.ResetDiscountPrice()
	suo.mutation.SetDiscountPrice(f)
	return suo
}

// AddDiscountPrice adds f to the "discount_price" field.
func (suo *SkuUpdateOne) AddDiscountPrice(f float64) *SkuUpdateOne {
	suo.mutation.AddDiscountPrice(f)
	return suo
}

// SetOnline sets the "online" field.
func (suo *SkuUpdateOne) SetOnline(i int8) *SkuUpdateOne {
	suo.mutation.ResetOnline()
	suo.mutation.SetOnline(i)
	return suo
}

// AddOnline adds i to the "online" field.
func (suo *SkuUpdateOne) AddOnline(i int8) *SkuUpdateOne {
	suo.mutation.AddOnline(i)
	return suo
}

// SetImg sets the "img" field.
func (suo *SkuUpdateOne) SetImg(s string) *SkuUpdateOne {
	suo.mutation.SetImg(s)
	return suo
}

// SetTitle sets the "title" field.
func (suo *SkuUpdateOne) SetTitle(s string) *SkuUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// SetSpuID sets the "spu_id" field.
func (suo *SkuUpdateOne) SetSpuID(i int64) *SkuUpdateOne {
	suo.mutation.ResetSpuID()
	suo.mutation.SetSpuID(i)
	return suo
}

// SetNillableSpuID sets the "spu_id" field if the given value is not nil.
func (suo *SkuUpdateOne) SetNillableSpuID(i *int64) *SkuUpdateOne {
	if i != nil {
		suo.SetSpuID(*i)
	}
	return suo
}

// AddSpuID adds i to the "spu_id" field.
func (suo *SkuUpdateOne) AddSpuID(i int64) *SkuUpdateOne {
	suo.mutation.AddSpuID(i)
	return suo
}

// ClearSpuID clears the value of the "spu_id" field.
func (suo *SkuUpdateOne) ClearSpuID() *SkuUpdateOne {
	suo.mutation.ClearSpuID()
	return suo
}

// SetSpecs sets the "specs" field.
func (suo *SkuUpdateOne) SetSpecs(s []schema.Spec) *SkuUpdateOne {
	suo.mutation.SetSpecs(s)
	return suo
}

// SetCode sets the "code" field.
func (suo *SkuUpdateOne) SetCode(s string) *SkuUpdateOne {
	suo.mutation.SetCode(s)
	return suo
}

// SetStock sets the "stock" field.
func (suo *SkuUpdateOne) SetStock(i int) *SkuUpdateOne {
	suo.mutation.ResetStock()
	suo.mutation.SetStock(i)
	return suo
}

// AddStock adds i to the "stock" field.
func (suo *SkuUpdateOne) AddStock(i int) *SkuUpdateOne {
	suo.mutation.AddStock(i)
	return suo
}

// SetCategoryID sets the "category_id" field.
func (suo *SkuUpdateOne) SetCategoryID(i int64) *SkuUpdateOne {
	suo.mutation.ResetCategoryID()
	suo.mutation.SetCategoryID(i)
	return suo
}

// AddCategoryID adds i to the "category_id" field.
func (suo *SkuUpdateOne) AddCategoryID(i int64) *SkuUpdateOne {
	suo.mutation.AddCategoryID(i)
	return suo
}

// SetRootCategoryID sets the "root_category_id" field.
func (suo *SkuUpdateOne) SetRootCategoryID(i int64) *SkuUpdateOne {
	suo.mutation.ResetRootCategoryID()
	suo.mutation.SetRootCategoryID(i)
	return suo
}

// AddRootCategoryID adds i to the "root_category_id" field.
func (suo *SkuUpdateOne) AddRootCategoryID(i int64) *SkuUpdateOne {
	suo.mutation.AddRootCategoryID(i)
	return suo
}

// Mutation returns the SkuMutation object of the builder.
func (suo *SkuUpdateOne) Mutation() *SkuMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SkuUpdateOne) Select(field string, fields ...string) *SkuUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Sku entity.
func (suo *SkuUpdateOne) Save(ctx context.Context) (*Sku, error) {
	var (
		err  error
		node *Sku
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SkuUpdateOne) SaveX(ctx context.Context) *Sku {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SkuUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SkuUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SkuUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := sku.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

func (suo *SkuUpdateOne) sqlSave(ctx context.Context) (_node *Sku, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sku.Table,
			Columns: sku.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sku.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Sku.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sku.FieldID)
		for _, f := range fields {
			if !sku.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != sku.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sku.FieldUpdateTime,
		})
	}
	if value, ok := suo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sku.FieldDeleteTime,
		})
	}
	if suo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sku.FieldDeleteTime,
		})
	}
	if value, ok := suo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldPrice,
		})
	}
	if value, ok := suo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldPrice,
		})
	}
	if value, ok := suo.mutation.DiscountPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldDiscountPrice,
		})
	}
	if value, ok := suo.mutation.AddedDiscountPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sku.FieldDiscountPrice,
		})
	}
	if value, ok := suo.mutation.Online(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: sku.FieldOnline,
		})
	}
	if value, ok := suo.mutation.AddedOnline(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: sku.FieldOnline,
		})
	}
	if value, ok := suo.mutation.Img(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sku.FieldImg,
		})
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sku.FieldTitle,
		})
	}
	if value, ok := suo.mutation.SpuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldSpuID,
		})
	}
	if value, ok := suo.mutation.AddedSpuID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldSpuID,
		})
	}
	if suo.mutation.SpuIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: sku.FieldSpuID,
		})
	}
	if value, ok := suo.mutation.Specs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: sku.FieldSpecs,
		})
	}
	if value, ok := suo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sku.FieldCode,
		})
	}
	if value, ok := suo.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sku.FieldStock,
		})
	}
	if value, ok := suo.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sku.FieldStock,
		})
	}
	if value, ok := suo.mutation.CategoryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldCategoryID,
		})
	}
	if value, ok := suo.mutation.AddedCategoryID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldCategoryID,
		})
	}
	if value, ok := suo.mutation.RootCategoryID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldRootCategoryID,
		})
	}
	if value, ok := suo.mutation.AddedRootCategoryID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sku.FieldRootCategoryID,
		})
	}
	_node = &Sku{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sku.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
