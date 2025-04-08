// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/memberprofile"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberProfileQuery is the builder for querying MemberProfile entities.
type MemberProfileQuery struct {
	config
	ctx        *QueryContext
	order      []memberprofile.OrderOption
	inters     []Interceptor
	predicates []predicate.MemberProfile
	withMember *MemberQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MemberProfileQuery builder.
func (mpq *MemberProfileQuery) Where(ps ...predicate.MemberProfile) *MemberProfileQuery {
	mpq.predicates = append(mpq.predicates, ps...)
	return mpq
}

// Limit the number of records to be returned by this query.
func (mpq *MemberProfileQuery) Limit(limit int) *MemberProfileQuery {
	mpq.ctx.Limit = &limit
	return mpq
}

// Offset to start from.
func (mpq *MemberProfileQuery) Offset(offset int) *MemberProfileQuery {
	mpq.ctx.Offset = &offset
	return mpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mpq *MemberProfileQuery) Unique(unique bool) *MemberProfileQuery {
	mpq.ctx.Unique = &unique
	return mpq
}

// Order specifies how the records should be ordered.
func (mpq *MemberProfileQuery) Order(o ...memberprofile.OrderOption) *MemberProfileQuery {
	mpq.order = append(mpq.order, o...)
	return mpq
}

// QueryMember chains the current query on the "member" edge.
func (mpq *MemberProfileQuery) QueryMember() *MemberQuery {
	query := (&MemberClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(memberprofile.Table, memberprofile.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, memberprofile.MemberTable, memberprofile.MemberColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MemberProfile entity from the query.
// Returns a *NotFoundError when no MemberProfile was found.
func (mpq *MemberProfileQuery) First(ctx context.Context) (*MemberProfile, error) {
	nodes, err := mpq.Limit(1).All(setContextOp(ctx, mpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{memberprofile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mpq *MemberProfileQuery) FirstX(ctx context.Context) *MemberProfile {
	node, err := mpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MemberProfile ID from the query.
// Returns a *NotFoundError when no MemberProfile ID was found.
func (mpq *MemberProfileQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(1).IDs(setContextOp(ctx, mpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{memberprofile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mpq *MemberProfileQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MemberProfile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MemberProfile entity is found.
// Returns a *NotFoundError when no MemberProfile entities are found.
func (mpq *MemberProfileQuery) Only(ctx context.Context) (*MemberProfile, error) {
	nodes, err := mpq.Limit(2).All(setContextOp(ctx, mpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{memberprofile.Label}
	default:
		return nil, &NotSingularError{memberprofile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mpq *MemberProfileQuery) OnlyX(ctx context.Context) *MemberProfile {
	node, err := mpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MemberProfile ID in the query.
// Returns a *NotSingularError when more than one MemberProfile ID is found.
// Returns a *NotFoundError when no entities are found.
func (mpq *MemberProfileQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(2).IDs(setContextOp(ctx, mpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{memberprofile.Label}
	default:
		err = &NotSingularError{memberprofile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mpq *MemberProfileQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MemberProfiles.
func (mpq *MemberProfileQuery) All(ctx context.Context) ([]*MemberProfile, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryAll)
	if err := mpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MemberProfile, *MemberProfileQuery]()
	return withInterceptors[[]*MemberProfile](ctx, mpq, qr, mpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mpq *MemberProfileQuery) AllX(ctx context.Context) []*MemberProfile {
	nodes, err := mpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MemberProfile IDs.
func (mpq *MemberProfileQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mpq.ctx.Unique == nil && mpq.path != nil {
		mpq.Unique(true)
	}
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryIDs)
	if err = mpq.Select(memberprofile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mpq *MemberProfileQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mpq *MemberProfileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryCount)
	if err := mpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mpq, querierCount[*MemberProfileQuery](), mpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mpq *MemberProfileQuery) CountX(ctx context.Context) int {
	count, err := mpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mpq *MemberProfileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryExist)
	switch _, err := mpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mpq *MemberProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := mpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MemberProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mpq *MemberProfileQuery) Clone() *MemberProfileQuery {
	if mpq == nil {
		return nil
	}
	return &MemberProfileQuery{
		config:     mpq.config,
		ctx:        mpq.ctx.Clone(),
		order:      append([]memberprofile.OrderOption{}, mpq.order...),
		inters:     append([]Interceptor{}, mpq.inters...),
		predicates: append([]predicate.MemberProfile{}, mpq.predicates...),
		withMember: mpq.withMember.Clone(),
		// clone intermediate query.
		sql:       mpq.sql.Clone(),
		path:      mpq.path,
		modifiers: append([]func(*sql.Selector){}, mpq.modifiers...),
	}
}

// WithMember tells the query-builder to eager-load the nodes that are connected to
// the "member" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MemberProfileQuery) WithMember(opts ...func(*MemberQuery)) *MemberProfileQuery {
	query := (&MemberClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMember = query
	return mpq
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
//	client.MemberProfile.Query().
//		GroupBy(memberprofile.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mpq *MemberProfileQuery) GroupBy(field string, fields ...string) *MemberProfileGroupBy {
	mpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MemberProfileGroupBy{build: mpq}
	grbuild.flds = &mpq.ctx.Fields
	grbuild.label = memberprofile.Label
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
//	client.MemberProfile.Query().
//		Select(memberprofile.FieldCreatedAt).
//		Scan(ctx, &v)
func (mpq *MemberProfileQuery) Select(fields ...string) *MemberProfileSelect {
	mpq.ctx.Fields = append(mpq.ctx.Fields, fields...)
	sbuild := &MemberProfileSelect{MemberProfileQuery: mpq}
	sbuild.label = memberprofile.Label
	sbuild.flds, sbuild.scan = &mpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MemberProfileSelect configured with the given aggregations.
func (mpq *MemberProfileQuery) Aggregate(fns ...AggregateFunc) *MemberProfileSelect {
	return mpq.Select().Aggregate(fns...)
}

func (mpq *MemberProfileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mpq); err != nil {
				return err
			}
		}
	}
	for _, f := range mpq.ctx.Fields {
		if !memberprofile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mpq.path != nil {
		prev, err := mpq.path(ctx)
		if err != nil {
			return err
		}
		mpq.sql = prev
	}
	return nil
}

func (mpq *MemberProfileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MemberProfile, error) {
	var (
		nodes       = []*MemberProfile{}
		_spec       = mpq.querySpec()
		loadedTypes = [1]bool{
			mpq.withMember != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MemberProfile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MemberProfile{config: mpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mpq.modifiers) > 0 {
		_spec.Modifiers = mpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mpq.withMember; query != nil {
		if err := mpq.loadMember(ctx, query, nodes, nil,
			func(n *MemberProfile, e *Member) { n.Edges.Member = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mpq *MemberProfileQuery) loadMember(ctx context.Context, query *MemberQuery, nodes []*MemberProfile, init func(*MemberProfile), assign func(*MemberProfile, *Member)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MemberProfile)
	for i := range nodes {
		fk := nodes[i].MemberID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(member.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mpq *MemberProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mpq.querySpec()
	if len(mpq.modifiers) > 0 {
		_spec.Modifiers = mpq.modifiers
	}
	_spec.Node.Columns = mpq.ctx.Fields
	if len(mpq.ctx.Fields) > 0 {
		_spec.Unique = mpq.ctx.Unique != nil && *mpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mpq.driver, _spec)
}

func (mpq *MemberProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(memberprofile.Table, memberprofile.Columns, sqlgraph.NewFieldSpec(memberprofile.FieldID, field.TypeInt64))
	_spec.From = mpq.sql
	if unique := mpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mpq.path != nil {
		_spec.Unique = true
	}
	if fields := mpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, memberprofile.FieldID)
		for i := range fields {
			if fields[i] != memberprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mpq.withMember != nil {
			_spec.Node.AddColumnOnce(memberprofile.FieldMemberID)
		}
	}
	if ps := mpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mpq *MemberProfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mpq.driver.Dialect())
	t1 := builder.Table(memberprofile.Table)
	columns := mpq.ctx.Fields
	if len(columns) == 0 {
		columns = memberprofile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mpq.sql != nil {
		selector = mpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mpq.ctx.Unique != nil && *mpq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mpq.modifiers {
		m(selector)
	}
	for _, p := range mpq.predicates {
		p(selector)
	}
	for _, p := range mpq.order {
		p(selector)
	}
	if offset := mpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mpq *MemberProfileQuery) Modify(modifiers ...func(s *sql.Selector)) *MemberProfileSelect {
	mpq.modifiers = append(mpq.modifiers, modifiers...)
	return mpq.Select()
}

// MemberProfileGroupBy is the group-by builder for MemberProfile entities.
type MemberProfileGroupBy struct {
	selector
	build *MemberProfileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mpgb *MemberProfileGroupBy) Aggregate(fns ...AggregateFunc) *MemberProfileGroupBy {
	mpgb.fns = append(mpgb.fns, fns...)
	return mpgb
}

// Scan applies the selector query and scans the result into the given value.
func (mpgb *MemberProfileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mpgb.build.ctx, ent.OpQueryGroupBy)
	if err := mpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberProfileQuery, *MemberProfileGroupBy](ctx, mpgb.build, mpgb, mpgb.build.inters, v)
}

func (mpgb *MemberProfileGroupBy) sqlScan(ctx context.Context, root *MemberProfileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mpgb.fns))
	for _, fn := range mpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mpgb.flds)+len(mpgb.fns))
		for _, f := range *mpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MemberProfileSelect is the builder for selecting fields of MemberProfile entities.
type MemberProfileSelect struct {
	*MemberProfileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mps *MemberProfileSelect) Aggregate(fns ...AggregateFunc) *MemberProfileSelect {
	mps.fns = append(mps.fns, fns...)
	return mps
}

// Scan applies the selector query and scans the result into the given value.
func (mps *MemberProfileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mps.ctx, ent.OpQuerySelect)
	if err := mps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberProfileQuery, *MemberProfileSelect](ctx, mps.MemberProfileQuery, mps, mps.inters, v)
}

func (mps *MemberProfileSelect) sqlScan(ctx context.Context, root *MemberProfileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mps.fns))
	for _, fn := range mps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mps *MemberProfileSelect) Modify(modifiers ...func(s *sql.Selector)) *MemberProfileSelect {
	mps.modifiers = append(mps.modifiers, modifiers...)
	return mps
}
