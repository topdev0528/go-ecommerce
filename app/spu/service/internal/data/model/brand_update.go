// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/app/spu/service/internal/data/model/brand"
	"mall-go/app/spu/service/internal/data/model/predicate"
	"mall-go/app/spu/service/internal/data/model/spu"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BrandUpdate is the builder for updating Brand entities.
type BrandUpdate struct {
	config
	hooks    []Hook
	mutation *BrandMutation
}

// Where appends a list predicates to the BrandUpdate builder.
func (bu *BrandUpdate) Where(ps ...predicate.Brand) *BrandUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdateTime sets the "update_time" field.
func (bu *BrandUpdate) SetUpdateTime(t time.Time) *BrandUpdate {
	bu.mutation.SetUpdateTime(t)
	return bu
}

// SetDeleteTime sets the "delete_time" field.
func (bu *BrandUpdate) SetDeleteTime(t time.Time) *BrandUpdate {
	bu.mutation.SetDeleteTime(t)
	return bu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (bu *BrandUpdate) SetNillableDeleteTime(t *time.Time) *BrandUpdate {
	if t != nil {
		bu.SetDeleteTime(*t)
	}
	return bu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (bu *BrandUpdate) ClearDeleteTime() *BrandUpdate {
	bu.mutation.ClearDeleteTime()
	return bu
}

// SetName sets the "name" field.
func (bu *BrandUpdate) SetName(s string) *BrandUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetDescription sets the "description" field.
func (bu *BrandUpdate) SetDescription(s string) *BrandUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// AddSpuIDs adds the "spu" edge to the Spu entity by IDs.
func (bu *BrandUpdate) AddSpuIDs(ids ...int64) *BrandUpdate {
	bu.mutation.AddSpuIDs(ids...)
	return bu
}

// AddSpu adds the "spu" edges to the Spu entity.
func (bu *BrandUpdate) AddSpu(s ...*Spu) *BrandUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bu.AddSpuIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (bu *BrandUpdate) Mutation() *BrandMutation {
	return bu.mutation
}

// ClearSpu clears all "spu" edges to the Spu entity.
func (bu *BrandUpdate) ClearSpu() *BrandUpdate {
	bu.mutation.ClearSpu()
	return bu
}

// RemoveSpuIDs removes the "spu" edge to Spu entities by IDs.
func (bu *BrandUpdate) RemoveSpuIDs(ids ...int64) *BrandUpdate {
	bu.mutation.RemoveSpuIDs(ids...)
	return bu
}

// RemoveSpu removes "spu" edges to Spu entities.
func (bu *BrandUpdate) RemoveSpu(s ...*Spu) *BrandUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bu.RemoveSpuIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BrandUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	bu.defaults()
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BrandUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BrandUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BrandUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BrandUpdate) defaults() {
	if _, ok := bu.mutation.UpdateTime(); !ok {
		v := brand.UpdateDefaultUpdateTime()
		bu.mutation.SetUpdateTime(v)
	}
}

func (bu *BrandUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   brand.Table,
			Columns: brand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: brand.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldUpdateTime,
		})
	}
	if value, ok := bu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldDeleteTime,
		})
	}
	if bu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: brand.FieldDeleteTime,
		})
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldName,
		})
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldDescription,
		})
	}
	if bu.mutation.SpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.SpuTable,
			Columns: []string{brand.SpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedSpuIDs(); len(nodes) > 0 && !bu.mutation.SpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.SpuTable,
			Columns: []string{brand.SpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.SpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.SpuTable,
			Columns: []string{brand.SpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{brand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BrandUpdateOne is the builder for updating a single Brand entity.
type BrandUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BrandMutation
}

// SetUpdateTime sets the "update_time" field.
func (buo *BrandUpdateOne) SetUpdateTime(t time.Time) *BrandUpdateOne {
	buo.mutation.SetUpdateTime(t)
	return buo
}

// SetDeleteTime sets the "delete_time" field.
func (buo *BrandUpdateOne) SetDeleteTime(t time.Time) *BrandUpdateOne {
	buo.mutation.SetDeleteTime(t)
	return buo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (buo *BrandUpdateOne) SetNillableDeleteTime(t *time.Time) *BrandUpdateOne {
	if t != nil {
		buo.SetDeleteTime(*t)
	}
	return buo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (buo *BrandUpdateOne) ClearDeleteTime() *BrandUpdateOne {
	buo.mutation.ClearDeleteTime()
	return buo
}

// SetName sets the "name" field.
func (buo *BrandUpdateOne) SetName(s string) *BrandUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetDescription sets the "description" field.
func (buo *BrandUpdateOne) SetDescription(s string) *BrandUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// AddSpuIDs adds the "spu" edge to the Spu entity by IDs.
func (buo *BrandUpdateOne) AddSpuIDs(ids ...int64) *BrandUpdateOne {
	buo.mutation.AddSpuIDs(ids...)
	return buo
}

// AddSpu adds the "spu" edges to the Spu entity.
func (buo *BrandUpdateOne) AddSpu(s ...*Spu) *BrandUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return buo.AddSpuIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (buo *BrandUpdateOne) Mutation() *BrandMutation {
	return buo.mutation
}

// ClearSpu clears all "spu" edges to the Spu entity.
func (buo *BrandUpdateOne) ClearSpu() *BrandUpdateOne {
	buo.mutation.ClearSpu()
	return buo
}

// RemoveSpuIDs removes the "spu" edge to Spu entities by IDs.
func (buo *BrandUpdateOne) RemoveSpuIDs(ids ...int64) *BrandUpdateOne {
	buo.mutation.RemoveSpuIDs(ids...)
	return buo
}

// RemoveSpu removes "spu" edges to Spu entities.
func (buo *BrandUpdateOne) RemoveSpu(s ...*Spu) *BrandUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return buo.RemoveSpuIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BrandUpdateOne) Select(field string, fields ...string) *BrandUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Brand entity.
func (buo *BrandUpdateOne) Save(ctx context.Context) (*Brand, error) {
	var (
		err  error
		node *Brand
	)
	buo.defaults()
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BrandUpdateOne) SaveX(ctx context.Context) *Brand {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BrandUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BrandUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BrandUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdateTime(); !ok {
		v := brand.UpdateDefaultUpdateTime()
		buo.mutation.SetUpdateTime(v)
	}
}

func (buo *BrandUpdateOne) sqlSave(ctx context.Context) (_node *Brand, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   brand.Table,
			Columns: brand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: brand.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Brand.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, brand.FieldID)
		for _, f := range fields {
			if !brand.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != brand.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldUpdateTime,
		})
	}
	if value, ok := buo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldDeleteTime,
		})
	}
	if buo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: brand.FieldDeleteTime,
		})
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldName,
		})
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldDescription,
		})
	}
	if buo.mutation.SpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.SpuTable,
			Columns: []string{brand.SpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedSpuIDs(); len(nodes) > 0 && !buo.mutation.SpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.SpuTable,
			Columns: []string{brand.SpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.SpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.SpuTable,
			Columns: []string{brand.SpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Brand{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{brand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}