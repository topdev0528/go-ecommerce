// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/sku/service/internal/data/model/orderdetail"
	"mall-go/app/sku/service/internal/data/model/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderDetailQuery is the builder for querying OrderDetail entities.
type OrderDetailQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderDetail
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderDetailQuery builder.
func (odq *OrderDetailQuery) Where(ps ...predicate.OrderDetail) *OrderDetailQuery {
	odq.predicates = append(odq.predicates, ps...)
	return odq
}

// Limit adds a limit step to the query.
func (odq *OrderDetailQuery) Limit(limit int) *OrderDetailQuery {
	odq.limit = &limit
	return odq
}

// Offset adds an offset step to the query.
func (odq *OrderDetailQuery) Offset(offset int) *OrderDetailQuery {
	odq.offset = &offset
	return odq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (odq *OrderDetailQuery) Unique(unique bool) *OrderDetailQuery {
	odq.unique = &unique
	return odq
}

// Order adds an order step to the query.
func (odq *OrderDetailQuery) Order(o ...OrderFunc) *OrderDetailQuery {
	odq.order = append(odq.order, o...)
	return odq
}

// First returns the first OrderDetail entity from the query.
// Returns a *NotFoundError when no OrderDetail was found.
func (odq *OrderDetailQuery) First(ctx context.Context) (*OrderDetail, error) {
	nodes, err := odq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderdetail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (odq *OrderDetailQuery) FirstX(ctx context.Context) *OrderDetail {
	node, err := odq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderDetail ID from the query.
// Returns a *NotFoundError when no OrderDetail ID was found.
func (odq *OrderDetailQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = odq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderdetail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (odq *OrderDetailQuery) FirstIDX(ctx context.Context) int64 {
	id, err := odq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Last returns the last OrderDetail entity from the query.
// Returns a *NotFoundError when no OrderDetail was found.
func (odq *OrderDetailQuery) Last(ctx context.Context) (*OrderDetail, error) {
	nodes, err := odq.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderdetail.Label}
	}
	return nodes[len(nodes)-1], nil
}

// LastX is like Last, but panics if an error occurs.
func (odq *OrderDetailQuery) LastX(ctx context.Context) *OrderDetail {
	node, err := odq.Last(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// LastID returns the last OrderDetail ID from the query.
// Returns a *NotFoundError when no OrderDetail ID was found.
func (odq *OrderDetailQuery) LastID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = odq.IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderdetail.Label}
		return
	}
	return ids[len(ids)-1], nil
}

// LastIDX is like LastID, but panics if an error occurs.
func (odq *OrderDetailQuery) LastIDX(ctx context.Context) int64 {
	id, err := odq.LastID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderDetail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OrderDetail entity is not found.
// Returns a *NotFoundError when no OrderDetail entities are found.
func (odq *OrderDetailQuery) Only(ctx context.Context) (*OrderDetail, error) {
	nodes, err := odq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderdetail.Label}
	default:
		return nil, &NotSingularError{orderdetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (odq *OrderDetailQuery) OnlyX(ctx context.Context) *OrderDetail {
	node, err := odq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderDetail ID in the query.
// Returns a *NotSingularError when exactly one OrderDetail ID is not found.
// Returns a *NotFoundError when no entities are found.
func (odq *OrderDetailQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = odq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = &NotSingularError{orderdetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (odq *OrderDetailQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := odq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderDetails.
func (odq *OrderDetailQuery) All(ctx context.Context) ([]*OrderDetail, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return odq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (odq *OrderDetailQuery) AllX(ctx context.Context) []*OrderDetail {
	nodes, err := odq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderDetail IDs.
func (odq *OrderDetailQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := odq.Select(orderdetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (odq *OrderDetailQuery) IDsX(ctx context.Context) []int64 {
	ids, err := odq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (odq *OrderDetailQuery) Count(ctx context.Context) (int, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return odq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (odq *OrderDetailQuery) CountX(ctx context.Context) int {
	count, err := odq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (odq *OrderDetailQuery) Exist(ctx context.Context) (bool, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return odq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (odq *OrderDetailQuery) ExistX(ctx context.Context) bool {
	exist, err := odq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderDetailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (odq *OrderDetailQuery) Clone() *OrderDetailQuery {
	if odq == nil {
		return nil
	}
	return &OrderDetailQuery{
		config:     odq.config,
		limit:      odq.limit,
		offset:     odq.offset,
		order:      append([]OrderFunc{}, odq.order...),
		predicates: append([]predicate.OrderDetail{}, odq.predicates...),
		// clone intermediate query.
		sql:  odq.sql.Clone(),
		path: odq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderDetail.Query().
//		GroupBy(orderdetail.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (odq *OrderDetailQuery) GroupBy(field string, fields ...string) *OrderDetailGroupBy {
	group := &OrderDetailGroupBy{config: odq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return odq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.OrderDetail.Query().
//		Select(orderdetail.FieldCreateTime).
//		Scan(ctx, &v)
//
func (odq *OrderDetailQuery) Select(fields ...string) *OrderDetailSelect {
	odq.fields = append(odq.fields, fields...)
	return &OrderDetailSelect{OrderDetailQuery: odq}
}

func (odq *OrderDetailQuery) prepareQuery(ctx context.Context) error {
	for _, f := range odq.fields {
		if !orderdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if odq.path != nil {
		prev, err := odq.path(ctx)
		if err != nil {
			return err
		}
		odq.sql = prev
	}
	return nil
}

func (odq *OrderDetailQuery) sqlAll(ctx context.Context) ([]*OrderDetail, error) {
	var (
		nodes = []*OrderDetail{}
		_spec = odq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrderDetail{config: odq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("model: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, odq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (odq *OrderDetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := odq.querySpec()
	return sqlgraph.CountNodes(ctx, odq.driver, _spec)
}

func (odq *OrderDetailQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := odq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (odq *OrderDetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderdetail.Table,
			Columns: orderdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: orderdetail.FieldID,
			},
		},
		From:   odq.sql,
		Unique: true,
	}
	if unique := odq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := odq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderdetail.FieldID)
		for i := range fields {
			if fields[i] != orderdetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := odq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := odq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := odq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := odq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (odq *OrderDetailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(odq.driver.Dialect())
	t1 := builder.Table(orderdetail.Table)
	columns := odq.fields
	if len(columns) == 0 {
		columns = orderdetail.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if odq.sql != nil {
		selector = odq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range odq.predicates {
		p(selector)
	}
	for _, p := range odq.order {
		p(selector)
	}
	if offset := odq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := odq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderDetailGroupBy is the group-by builder for OrderDetail entities.
type OrderDetailGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (odgb *OrderDetailGroupBy) Aggregate(fns ...AggregateFunc) *OrderDetailGroupBy {
	odgb.fns = append(odgb.fns, fns...)
	return odgb
}

// Scan applies the group-by query and scans the result into the given value.
func (odgb *OrderDetailGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := odgb.path(ctx)
	if err != nil {
		return err
	}
	odgb.sql = query
	return odgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (odgb *OrderDetailGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := odgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (odgb *OrderDetailGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(odgb.fields) > 1 {
		return nil, errors.New("model: OrderDetailGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := odgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (odgb *OrderDetailGroupBy) StringsX(ctx context.Context) []string {
	v, err := odgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (odgb *OrderDetailGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = odgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = fmt.Errorf("model: OrderDetailGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (odgb *OrderDetailGroupBy) StringX(ctx context.Context) string {
	v, err := odgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (odgb *OrderDetailGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(odgb.fields) > 1 {
		return nil, errors.New("model: OrderDetailGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := odgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (odgb *OrderDetailGroupBy) IntsX(ctx context.Context) []int {
	v, err := odgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (odgb *OrderDetailGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = odgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = fmt.Errorf("model: OrderDetailGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (odgb *OrderDetailGroupBy) IntX(ctx context.Context) int {
	v, err := odgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (odgb *OrderDetailGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(odgb.fields) > 1 {
		return nil, errors.New("model: OrderDetailGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := odgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (odgb *OrderDetailGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := odgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (odgb *OrderDetailGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = odgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = fmt.Errorf("model: OrderDetailGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (odgb *OrderDetailGroupBy) Float64X(ctx context.Context) float64 {
	v, err := odgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (odgb *OrderDetailGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(odgb.fields) > 1 {
		return nil, errors.New("model: OrderDetailGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := odgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (odgb *OrderDetailGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := odgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (odgb *OrderDetailGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = odgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = fmt.Errorf("model: OrderDetailGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (odgb *OrderDetailGroupBy) BoolX(ctx context.Context) bool {
	v, err := odgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (odgb *OrderDetailGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range odgb.fields {
		if !orderdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := odgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := odgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (odgb *OrderDetailGroupBy) sqlQuery() *sql.Selector {
	selector := odgb.sql.Select()
	aggregation := make([]string, 0, len(odgb.fns))
	for _, fn := range odgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(odgb.fields)+len(odgb.fns))
		for _, f := range odgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(odgb.fields...)...)
}

// OrderDetailSelect is the builder for selecting fields of OrderDetail entities.
type OrderDetailSelect struct {
	*OrderDetailQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ods *OrderDetailSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ods.prepareQuery(ctx); err != nil {
		return err
	}
	ods.sql = ods.OrderDetailQuery.sqlQuery(ctx)
	return ods.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ods *OrderDetailSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ods.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ods *OrderDetailSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ods.fields) > 1 {
		return nil, errors.New("model: OrderDetailSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ods *OrderDetailSelect) StringsX(ctx context.Context) []string {
	v, err := ods.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ods *OrderDetailSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ods.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = fmt.Errorf("model: OrderDetailSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ods *OrderDetailSelect) StringX(ctx context.Context) string {
	v, err := ods.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ods *OrderDetailSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ods.fields) > 1 {
		return nil, errors.New("model: OrderDetailSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ods *OrderDetailSelect) IntsX(ctx context.Context) []int {
	v, err := ods.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ods *OrderDetailSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ods.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = fmt.Errorf("model: OrderDetailSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ods *OrderDetailSelect) IntX(ctx context.Context) int {
	v, err := ods.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ods *OrderDetailSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ods.fields) > 1 {
		return nil, errors.New("model: OrderDetailSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ods *OrderDetailSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ods.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ods *OrderDetailSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ods.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = fmt.Errorf("model: OrderDetailSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ods *OrderDetailSelect) Float64X(ctx context.Context) float64 {
	v, err := ods.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ods *OrderDetailSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ods.fields) > 1 {
		return nil, errors.New("model: OrderDetailSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ods *OrderDetailSelect) BoolsX(ctx context.Context) []bool {
	v, err := ods.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ods *OrderDetailSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ods.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderdetail.Label}
	default:
		err = fmt.Errorf("model: OrderDetailSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ods *OrderDetailSelect) BoolX(ctx context.Context) bool {
	v, err := ods.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ods *OrderDetailSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ods.sql.Query()
	if err := ods.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}