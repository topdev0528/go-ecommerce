// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/order/service/internal/data/model/order"
	"mall-go/app/order/service/internal/data/model/ordersnap"
	"mall-go/app/order/service/internal/data/model/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderSnapQuery is the builder for querying OrderSnap entities.
type OrderSnapQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderSnap
	// eager-loading edges.
	withOrder *OrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderSnapQuery builder.
func (osq *OrderSnapQuery) Where(ps ...predicate.OrderSnap) *OrderSnapQuery {
	osq.predicates = append(osq.predicates, ps...)
	return osq
}

// Limit adds a limit step to the query.
func (osq *OrderSnapQuery) Limit(limit int) *OrderSnapQuery {
	osq.limit = &limit
	return osq
}

// Offset adds an offset step to the query.
func (osq *OrderSnapQuery) Offset(offset int) *OrderSnapQuery {
	osq.offset = &offset
	return osq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osq *OrderSnapQuery) Unique(unique bool) *OrderSnapQuery {
	osq.unique = &unique
	return osq
}

// Order adds an order step to the query.
func (osq *OrderSnapQuery) Order(o ...OrderFunc) *OrderSnapQuery {
	osq.order = append(osq.order, o...)
	return osq
}

// QueryOrder chains the current query on the "order" edge.
func (osq *OrderSnapQuery) QueryOrder() *OrderQuery {
	query := &OrderQuery{config: osq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ordersnap.Table, ordersnap.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ordersnap.OrderTable, ordersnap.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderSnap entity from the query.
// Returns a *NotFoundError when no OrderSnap was found.
func (osq *OrderSnapQuery) First(ctx context.Context) (*OrderSnap, error) {
	nodes, err := osq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ordersnap.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osq *OrderSnapQuery) FirstX(ctx context.Context) *OrderSnap {
	node, err := osq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderSnap ID from the query.
// Returns a *NotFoundError when no OrderSnap ID was found.
func (osq *OrderSnapQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = osq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ordersnap.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osq *OrderSnapQuery) FirstIDX(ctx context.Context) int64 {
	id, err := osq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderSnap entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderSnap entity is found.
// Returns a *NotFoundError when no OrderSnap entities are found.
func (osq *OrderSnapQuery) Only(ctx context.Context) (*OrderSnap, error) {
	nodes, err := osq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ordersnap.Label}
	default:
		return nil, &NotSingularError{ordersnap.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osq *OrderSnapQuery) OnlyX(ctx context.Context) *OrderSnap {
	node, err := osq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderSnap ID in the query.
// Returns a *NotSingularError when more than one OrderSnap ID is found.
// Returns a *NotFoundError when no entities are found.
func (osq *OrderSnapQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = osq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ordersnap.Label}
	default:
		err = &NotSingularError{ordersnap.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osq *OrderSnapQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := osq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderSnaps.
func (osq *OrderSnapQuery) All(ctx context.Context) ([]*OrderSnap, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return osq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (osq *OrderSnapQuery) AllX(ctx context.Context) []*OrderSnap {
	nodes, err := osq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderSnap IDs.
func (osq *OrderSnapQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := osq.Select(ordersnap.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osq *OrderSnapQuery) IDsX(ctx context.Context) []int64 {
	ids, err := osq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osq *OrderSnapQuery) Count(ctx context.Context) (int, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return osq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (osq *OrderSnapQuery) CountX(ctx context.Context) int {
	count, err := osq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osq *OrderSnapQuery) Exist(ctx context.Context) (bool, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return osq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (osq *OrderSnapQuery) ExistX(ctx context.Context) bool {
	exist, err := osq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderSnapQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osq *OrderSnapQuery) Clone() *OrderSnapQuery {
	if osq == nil {
		return nil
	}
	return &OrderSnapQuery{
		config:     osq.config,
		limit:      osq.limit,
		offset:     osq.offset,
		order:      append([]OrderFunc{}, osq.order...),
		predicates: append([]predicate.OrderSnap{}, osq.predicates...),
		withOrder:  osq.withOrder.Clone(),
		// clone intermediate query.
		sql:    osq.sql.Clone(),
		path:   osq.path,
		unique: osq.unique,
	}
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrderSnapQuery) WithOrder(opts ...func(*OrderQuery)) *OrderSnapQuery {
	query := &OrderQuery{config: osq.config}
	for _, opt := range opts {
		opt(query)
	}
	osq.withOrder = query
	return osq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SnapImg string `json:"snap_img,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderSnap.Query().
//		GroupBy(ordersnap.FieldSnapImg).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (osq *OrderSnapQuery) GroupBy(field string, fields ...string) *OrderSnapGroupBy {
	group := &OrderSnapGroupBy{config: osq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return osq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SnapImg string `json:"snap_img,omitempty"`
//	}
//
//	client.OrderSnap.Query().
//		Select(ordersnap.FieldSnapImg).
//		Scan(ctx, &v)
//
func (osq *OrderSnapQuery) Select(fields ...string) *OrderSnapSelect {
	osq.fields = append(osq.fields, fields...)
	return &OrderSnapSelect{OrderSnapQuery: osq}
}

func (osq *OrderSnapQuery) prepareQuery(ctx context.Context) error {
	for _, f := range osq.fields {
		if !ordersnap.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if osq.path != nil {
		prev, err := osq.path(ctx)
		if err != nil {
			return err
		}
		osq.sql = prev
	}
	return nil
}

func (osq *OrderSnapQuery) sqlAll(ctx context.Context) ([]*OrderSnap, error) {
	var (
		nodes       = []*OrderSnap{}
		_spec       = osq.querySpec()
		loadedTypes = [1]bool{
			osq.withOrder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrderSnap{config: osq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("model: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, osq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := osq.withOrder; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*OrderSnap)
		for i := range nodes {
			fk := nodes[i].OrderID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(order.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Order = n
			}
		}
	}

	return nodes, nil
}

func (osq *OrderSnapQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osq.querySpec()
	_spec.Node.Columns = osq.fields
	if len(osq.fields) > 0 {
		_spec.Unique = osq.unique != nil && *osq.unique
	}
	return sqlgraph.CountNodes(ctx, osq.driver, _spec)
}

func (osq *OrderSnapQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := osq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (osq *OrderSnapQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordersnap.Table,
			Columns: ordersnap.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: ordersnap.FieldID,
			},
		},
		From:   osq.sql,
		Unique: true,
	}
	if unique := osq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := osq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ordersnap.FieldID)
		for i := range fields {
			if fields[i] != ordersnap.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := osq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osq *OrderSnapQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osq.driver.Dialect())
	t1 := builder.Table(ordersnap.Table)
	columns := osq.fields
	if len(columns) == 0 {
		columns = ordersnap.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osq.sql != nil {
		selector = osq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if osq.unique != nil && *osq.unique {
		selector.Distinct()
	}
	for _, p := range osq.predicates {
		p(selector)
	}
	for _, p := range osq.order {
		p(selector)
	}
	if offset := osq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderSnapGroupBy is the group-by builder for OrderSnap entities.
type OrderSnapGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osgb *OrderSnapGroupBy) Aggregate(fns ...AggregateFunc) *OrderSnapGroupBy {
	osgb.fns = append(osgb.fns, fns...)
	return osgb
}

// Scan applies the group-by query and scans the result into the given value.
func (osgb *OrderSnapGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := osgb.path(ctx)
	if err != nil {
		return err
	}
	osgb.sql = query
	return osgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (osgb *OrderSnapGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := osgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderSnapGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("model: OrderSnapGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (osgb *OrderSnapGroupBy) StringsX(ctx context.Context) []string {
	v, err := osgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderSnapGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = osgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordersnap.Label}
	default:
		err = fmt.Errorf("model: OrderSnapGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (osgb *OrderSnapGroupBy) StringX(ctx context.Context) string {
	v, err := osgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderSnapGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("model: OrderSnapGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (osgb *OrderSnapGroupBy) IntsX(ctx context.Context) []int {
	v, err := osgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderSnapGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = osgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordersnap.Label}
	default:
		err = fmt.Errorf("model: OrderSnapGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (osgb *OrderSnapGroupBy) IntX(ctx context.Context) int {
	v, err := osgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderSnapGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("model: OrderSnapGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (osgb *OrderSnapGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := osgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderSnapGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = osgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordersnap.Label}
	default:
		err = fmt.Errorf("model: OrderSnapGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (osgb *OrderSnapGroupBy) Float64X(ctx context.Context) float64 {
	v, err := osgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderSnapGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("model: OrderSnapGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (osgb *OrderSnapGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := osgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderSnapGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = osgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordersnap.Label}
	default:
		err = fmt.Errorf("model: OrderSnapGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (osgb *OrderSnapGroupBy) BoolX(ctx context.Context) bool {
	v, err := osgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (osgb *OrderSnapGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range osgb.fields {
		if !ordersnap.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := osgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (osgb *OrderSnapGroupBy) sqlQuery() *sql.Selector {
	selector := osgb.sql.Select()
	aggregation := make([]string, 0, len(osgb.fns))
	for _, fn := range osgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(osgb.fields)+len(osgb.fns))
		for _, f := range osgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(osgb.fields...)...)
}

// OrderSnapSelect is the builder for selecting fields of OrderSnap entities.
type OrderSnapSelect struct {
	*OrderSnapQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oss *OrderSnapSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oss.prepareQuery(ctx); err != nil {
		return err
	}
	oss.sql = oss.OrderSnapQuery.sqlQuery(ctx)
	return oss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oss *OrderSnapSelect) ScanX(ctx context.Context, v interface{}) {
	if err := oss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (oss *OrderSnapSelect) Strings(ctx context.Context) ([]string, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("model: OrderSnapSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oss *OrderSnapSelect) StringsX(ctx context.Context) []string {
	v, err := oss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (oss *OrderSnapSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordersnap.Label}
	default:
		err = fmt.Errorf("model: OrderSnapSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oss *OrderSnapSelect) StringX(ctx context.Context) string {
	v, err := oss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (oss *OrderSnapSelect) Ints(ctx context.Context) ([]int, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("model: OrderSnapSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oss *OrderSnapSelect) IntsX(ctx context.Context) []int {
	v, err := oss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (oss *OrderSnapSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordersnap.Label}
	default:
		err = fmt.Errorf("model: OrderSnapSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oss *OrderSnapSelect) IntX(ctx context.Context) int {
	v, err := oss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (oss *OrderSnapSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("model: OrderSnapSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oss *OrderSnapSelect) Float64sX(ctx context.Context) []float64 {
	v, err := oss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (oss *OrderSnapSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordersnap.Label}
	default:
		err = fmt.Errorf("model: OrderSnapSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oss *OrderSnapSelect) Float64X(ctx context.Context) float64 {
	v, err := oss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (oss *OrderSnapSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("model: OrderSnapSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oss *OrderSnapSelect) BoolsX(ctx context.Context) []bool {
	v, err := oss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (oss *OrderSnapSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordersnap.Label}
	default:
		err = fmt.Errorf("model: OrderSnapSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oss *OrderSnapSelect) BoolX(ctx context.Context) bool {
	v, err := oss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oss *OrderSnapSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oss.sql.Query()
	if err := oss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
