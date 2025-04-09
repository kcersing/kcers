// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	entorder "kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/dal/db/mysql/ent/orderpay"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderPayQuery is the builder for querying OrderPay entities.
type OrderPayQuery struct {
	config
	ctx        *QueryContext
	order      []orderpay.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderPay
	withOrder  *OrderQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderPayQuery builder.
func (opq *OrderPayQuery) Where(ps ...predicate.OrderPay) *OrderPayQuery {
	opq.predicates = append(opq.predicates, ps...)
	return opq
}

// Limit the number of records to be returned by this query.
func (opq *OrderPayQuery) Limit(limit int) *OrderPayQuery {
	opq.ctx.Limit = &limit
	return opq
}

// Offset to start from.
func (opq *OrderPayQuery) Offset(offset int) *OrderPayQuery {
	opq.ctx.Offset = &offset
	return opq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (opq *OrderPayQuery) Unique(unique bool) *OrderPayQuery {
	opq.ctx.Unique = &unique
	return opq
}

// Order specifies how the records should be ordered.
func (opq *OrderPayQuery) Order(o ...orderpay.OrderOption) *OrderPayQuery {
	opq.order = append(opq.order, o...)
	return opq
}

// QueryOrder chains the current query on the "order" edge.
func (opq *OrderPayQuery) QueryOrder() *OrderQuery {
	query := (&OrderClient{config: opq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := opq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderpay.Table, orderpay.FieldID, selector),
			sqlgraph.To(entorder.Table, entorder.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderpay.OrderTable, orderpay.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(opq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderPay entity from the query.
// Returns a *NotFoundError when no OrderPay was found.
func (opq *OrderPayQuery) First(ctx context.Context) (*OrderPay, error) {
	nodes, err := opq.Limit(1).All(setContextOp(ctx, opq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderpay.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (opq *OrderPayQuery) FirstX(ctx context.Context) *OrderPay {
	node, err := opq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderPay ID from the query.
// Returns a *NotFoundError when no OrderPay ID was found.
func (opq *OrderPayQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = opq.Limit(1).IDs(setContextOp(ctx, opq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderpay.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (opq *OrderPayQuery) FirstIDX(ctx context.Context) int64 {
	id, err := opq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderPay entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderPay entity is found.
// Returns a *NotFoundError when no OrderPay entities are found.
func (opq *OrderPayQuery) Only(ctx context.Context) (*OrderPay, error) {
	nodes, err := opq.Limit(2).All(setContextOp(ctx, opq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderpay.Label}
	default:
		return nil, &NotSingularError{orderpay.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (opq *OrderPayQuery) OnlyX(ctx context.Context) *OrderPay {
	node, err := opq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderPay ID in the query.
// Returns a *NotSingularError when more than one OrderPay ID is found.
// Returns a *NotFoundError when no entities are found.
func (opq *OrderPayQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = opq.Limit(2).IDs(setContextOp(ctx, opq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderpay.Label}
	default:
		err = &NotSingularError{orderpay.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (opq *OrderPayQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := opq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderPays.
func (opq *OrderPayQuery) All(ctx context.Context) ([]*OrderPay, error) {
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryAll)
	if err := opq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderPay, *OrderPayQuery]()
	return withInterceptors[[]*OrderPay](ctx, opq, qr, opq.inters)
}

// AllX is like All, but panics if an error occurs.
func (opq *OrderPayQuery) AllX(ctx context.Context) []*OrderPay {
	nodes, err := opq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderPay IDs.
func (opq *OrderPayQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if opq.ctx.Unique == nil && opq.path != nil {
		opq.Unique(true)
	}
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryIDs)
	if err = opq.Select(orderpay.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (opq *OrderPayQuery) IDsX(ctx context.Context) []int64 {
	ids, err := opq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (opq *OrderPayQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryCount)
	if err := opq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, opq, querierCount[*OrderPayQuery](), opq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (opq *OrderPayQuery) CountX(ctx context.Context) int {
	count, err := opq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (opq *OrderPayQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, opq.ctx, ent.OpQueryExist)
	switch _, err := opq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (opq *OrderPayQuery) ExistX(ctx context.Context) bool {
	exist, err := opq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderPayQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (opq *OrderPayQuery) Clone() *OrderPayQuery {
	if opq == nil {
		return nil
	}
	return &OrderPayQuery{
		config:     opq.config,
		ctx:        opq.ctx.Clone(),
		order:      append([]orderpay.OrderOption{}, opq.order...),
		inters:     append([]Interceptor{}, opq.inters...),
		predicates: append([]predicate.OrderPay{}, opq.predicates...),
		withOrder:  opq.withOrder.Clone(),
		// clone intermediate query.
		sql:       opq.sql.Clone(),
		path:      opq.path,
		modifiers: append([]func(*sql.Selector){}, opq.modifiers...),
	}
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (opq *OrderPayQuery) WithOrder(opts ...func(*OrderQuery)) *OrderPayQuery {
	query := (&OrderClient{config: opq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	opq.withOrder = query
	return opq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderPay.Query().
//		GroupBy(orderpay.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (opq *OrderPayQuery) GroupBy(field string, fields ...string) *OrderPayGroupBy {
	opq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderPayGroupBy{build: opq}
	grbuild.flds = &opq.ctx.Fields
	grbuild.label = orderpay.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.OrderPay.Query().
//		Select(orderpay.FieldCreatedAt).
//		Scan(ctx, &v)
func (opq *OrderPayQuery) Select(fields ...string) *OrderPaySelect {
	opq.ctx.Fields = append(opq.ctx.Fields, fields...)
	sbuild := &OrderPaySelect{OrderPayQuery: opq}
	sbuild.label = orderpay.Label
	sbuild.flds, sbuild.scan = &opq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderPaySelect configured with the given aggregations.
func (opq *OrderPayQuery) Aggregate(fns ...AggregateFunc) *OrderPaySelect {
	return opq.Select().Aggregate(fns...)
}

func (opq *OrderPayQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range opq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, opq); err != nil {
				return err
			}
		}
	}
	for _, f := range opq.ctx.Fields {
		if !orderpay.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if opq.path != nil {
		prev, err := opq.path(ctx)
		if err != nil {
			return err
		}
		opq.sql = prev
	}
	return nil
}

func (opq *OrderPayQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderPay, error) {
	var (
		nodes       = []*OrderPay{}
		_spec       = opq.querySpec()
		loadedTypes = [1]bool{
			opq.withOrder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderPay).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderPay{config: opq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(opq.modifiers) > 0 {
		_spec.Modifiers = opq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, opq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := opq.withOrder; query != nil {
		if err := opq.loadOrder(ctx, query, nodes, nil,
			func(n *OrderPay, e *Order) { n.Edges.Order = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (opq *OrderPayQuery) loadOrder(ctx context.Context, query *OrderQuery, nodes []*OrderPay, init func(*OrderPay), assign func(*OrderPay, *Order)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*OrderPay)
	for i := range nodes {
		fk := nodes[i].OrderID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(entorder.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (opq *OrderPayQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := opq.querySpec()
	if len(opq.modifiers) > 0 {
		_spec.Modifiers = opq.modifiers
	}
	_spec.Node.Columns = opq.ctx.Fields
	if len(opq.ctx.Fields) > 0 {
		_spec.Unique = opq.ctx.Unique != nil && *opq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, opq.driver, _spec)
}

func (opq *OrderPayQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderpay.Table, orderpay.Columns, sqlgraph.NewFieldSpec(orderpay.FieldID, field.TypeInt64))
	_spec.From = opq.sql
	if unique := opq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if opq.path != nil {
		_spec.Unique = true
	}
	if fields := opq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderpay.FieldID)
		for i := range fields {
			if fields[i] != orderpay.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if opq.withOrder != nil {
			_spec.Node.AddColumnOnce(orderpay.FieldOrderID)
		}
	}
	if ps := opq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := opq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := opq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := opq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (opq *OrderPayQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(opq.driver.Dialect())
	t1 := builder.Table(orderpay.Table)
	columns := opq.ctx.Fields
	if len(columns) == 0 {
		columns = orderpay.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if opq.sql != nil {
		selector = opq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if opq.ctx.Unique != nil && *opq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range opq.modifiers {
		m(selector)
	}
	for _, p := range opq.predicates {
		p(selector)
	}
	for _, p := range opq.order {
		p(selector)
	}
	if offset := opq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := opq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (opq *OrderPayQuery) Modify(modifiers ...func(s *sql.Selector)) *OrderPaySelect {
	opq.modifiers = append(opq.modifiers, modifiers...)
	return opq.Select()
}

// OrderPayGroupBy is the group-by builder for OrderPay entities.
type OrderPayGroupBy struct {
	selector
	build *OrderPayQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (opgb *OrderPayGroupBy) Aggregate(fns ...AggregateFunc) *OrderPayGroupBy {
	opgb.fns = append(opgb.fns, fns...)
	return opgb
}

// Scan applies the selector query and scans the result into the given value.
func (opgb *OrderPayGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, opgb.build.ctx, ent.OpQueryGroupBy)
	if err := opgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderPayQuery, *OrderPayGroupBy](ctx, opgb.build, opgb, opgb.build.inters, v)
}

func (opgb *OrderPayGroupBy) sqlScan(ctx context.Context, root *OrderPayQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(opgb.fns))
	for _, fn := range opgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*opgb.flds)+len(opgb.fns))
		for _, f := range *opgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*opgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := opgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderPaySelect is the builder for selecting fields of OrderPay entities.
type OrderPaySelect struct {
	*OrderPayQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ops *OrderPaySelect) Aggregate(fns ...AggregateFunc) *OrderPaySelect {
	ops.fns = append(ops.fns, fns...)
	return ops
}

// Scan applies the selector query and scans the result into the given value.
func (ops *OrderPaySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ops.ctx, ent.OpQuerySelect)
	if err := ops.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderPayQuery, *OrderPaySelect](ctx, ops.OrderPayQuery, ops, ops.inters, v)
}

func (ops *OrderPaySelect) sqlScan(ctx context.Context, root *OrderPayQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ops.fns))
	for _, fn := range ops.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ops.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ops.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ops *OrderPaySelect) Modify(modifiers ...func(s *sql.Selector)) *OrderPaySelect {
	ops.modifiers = append(ops.modifiers, modifiers...)
	return ops
}
