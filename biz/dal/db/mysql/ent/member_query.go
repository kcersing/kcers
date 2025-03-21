// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/entrylogs"
	"kcers/biz/dal/db/mysql/ent/face"
	"kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/membercontract"
	"kcers/biz/dal/db/mysql/ent/memberdetails"
	"kcers/biz/dal/db/mysql/ent/membernote"
	"kcers/biz/dal/db/mysql/ent/memberproduct"
	"kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberQuery is the builder for querying Member entities.
type MemberQuery struct {
	config
	ctx                *QueryContext
	order              []member.OrderOption
	inters             []Interceptor
	predicates         []predicate.Member
	withMemberDetails  *MemberDetailsQuery
	withMemberNotes    *MemberNoteQuery
	withMemberOrders   *OrderQuery
	withMemberProducts *MemberProductQuery
	withMemberEntry    *EntryLogsQuery
	withMemberContents *MemberContractQuery
	withMemberFace     *FaceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MemberQuery builder.
func (mq *MemberQuery) Where(ps ...predicate.Member) *MemberQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MemberQuery) Limit(limit int) *MemberQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MemberQuery) Offset(offset int) *MemberQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MemberQuery) Unique(unique bool) *MemberQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MemberQuery) Order(o ...member.OrderOption) *MemberQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryMemberDetails chains the current query on the "member_details" edge.
func (mq *MemberQuery) QueryMemberDetails() *MemberDetailsQuery {
	query := (&MemberDetailsClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(memberdetails.Table, memberdetails.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberDetailsTable, member.MemberDetailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberNotes chains the current query on the "member_notes" edge.
func (mq *MemberQuery) QueryMemberNotes() *MemberNoteQuery {
	query := (&MemberNoteClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(membernote.Table, membernote.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberNotesTable, member.MemberNotesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberOrders chains the current query on the "member_orders" edge.
func (mq *MemberQuery) QueryMemberOrders() *OrderQuery {
	query := (&OrderClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberOrdersTable, member.MemberOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberProducts chains the current query on the "member_products" edge.
func (mq *MemberQuery) QueryMemberProducts() *MemberProductQuery {
	query := (&MemberProductClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(memberproduct.Table, memberproduct.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberProductsTable, member.MemberProductsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberEntry chains the current query on the "member_entry" edge.
func (mq *MemberQuery) QueryMemberEntry() *EntryLogsQuery {
	query := (&EntryLogsClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(entrylogs.Table, entrylogs.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberEntryTable, member.MemberEntryColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberContents chains the current query on the "member_contents" edge.
func (mq *MemberQuery) QueryMemberContents() *MemberContractQuery {
	query := (&MemberContractClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(membercontract.Table, membercontract.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberContentsTable, member.MemberContentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberFace chains the current query on the "member_face" edge.
func (mq *MemberQuery) QueryMemberFace() *FaceQuery {
	query := (&FaceClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(face.Table, face.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberFaceTable, member.MemberFaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Member entity from the query.
// Returns a *NotFoundError when no Member was found.
func (mq *MemberQuery) First(ctx context.Context) (*Member, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{member.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MemberQuery) FirstX(ctx context.Context) *Member {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Member ID from the query.
// Returns a *NotFoundError when no Member ID was found.
func (mq *MemberQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{member.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MemberQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Member entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Member entity is found.
// Returns a *NotFoundError when no Member entities are found.
func (mq *MemberQuery) Only(ctx context.Context) (*Member, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{member.Label}
	default:
		return nil, &NotSingularError{member.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MemberQuery) OnlyX(ctx context.Context) *Member {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Member ID in the query.
// Returns a *NotSingularError when more than one Member ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MemberQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = &NotSingularError{member.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MemberQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Members.
func (mq *MemberQuery) All(ctx context.Context) ([]*Member, error) {
	ctx = setContextOp(ctx, mq.ctx, ent.OpQueryAll)
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Member, *MemberQuery]()
	return withInterceptors[[]*Member](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MemberQuery) AllX(ctx context.Context) []*Member {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Member IDs.
func (mq *MemberQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, ent.OpQueryIDs)
	if err = mq.Select(member.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MemberQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MemberQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, ent.OpQueryCount)
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MemberQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MemberQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MemberQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, ent.OpQueryExist)
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MemberQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MemberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MemberQuery) Clone() *MemberQuery {
	if mq == nil {
		return nil
	}
	return &MemberQuery{
		config:             mq.config,
		ctx:                mq.ctx.Clone(),
		order:              append([]member.OrderOption{}, mq.order...),
		inters:             append([]Interceptor{}, mq.inters...),
		predicates:         append([]predicate.Member{}, mq.predicates...),
		withMemberDetails:  mq.withMemberDetails.Clone(),
		withMemberNotes:    mq.withMemberNotes.Clone(),
		withMemberOrders:   mq.withMemberOrders.Clone(),
		withMemberProducts: mq.withMemberProducts.Clone(),
		withMemberEntry:    mq.withMemberEntry.Clone(),
		withMemberContents: mq.withMemberContents.Clone(),
		withMemberFace:     mq.withMemberFace.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithMemberDetails tells the query-builder to eager-load the nodes that are connected to
// the "member_details" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberDetails(opts ...func(*MemberDetailsQuery)) *MemberQuery {
	query := (&MemberDetailsClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberDetails = query
	return mq
}

// WithMemberNotes tells the query-builder to eager-load the nodes that are connected to
// the "member_notes" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberNotes(opts ...func(*MemberNoteQuery)) *MemberQuery {
	query := (&MemberNoteClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberNotes = query
	return mq
}

// WithMemberOrders tells the query-builder to eager-load the nodes that are connected to
// the "member_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberOrders(opts ...func(*OrderQuery)) *MemberQuery {
	query := (&OrderClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberOrders = query
	return mq
}

// WithMemberProducts tells the query-builder to eager-load the nodes that are connected to
// the "member_products" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberProducts(opts ...func(*MemberProductQuery)) *MemberQuery {
	query := (&MemberProductClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberProducts = query
	return mq
}

// WithMemberEntry tells the query-builder to eager-load the nodes that are connected to
// the "member_entry" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberEntry(opts ...func(*EntryLogsQuery)) *MemberQuery {
	query := (&EntryLogsClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberEntry = query
	return mq
}

// WithMemberContents tells the query-builder to eager-load the nodes that are connected to
// the "member_contents" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberContents(opts ...func(*MemberContractQuery)) *MemberQuery {
	query := (&MemberContractClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberContents = query
	return mq
}

// WithMemberFace tells the query-builder to eager-load the nodes that are connected to
// the "member_face" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberFace(opts ...func(*FaceQuery)) *MemberQuery {
	query := (&FaceClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberFace = query
	return mq
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
//	client.Member.Query().
//		GroupBy(member.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MemberQuery) GroupBy(field string, fields ...string) *MemberGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MemberGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = member.Label
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
//	client.Member.Query().
//		Select(member.FieldCreatedAt).
//		Scan(ctx, &v)
func (mq *MemberQuery) Select(fields ...string) *MemberSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MemberSelect{MemberQuery: mq}
	sbuild.label = member.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MemberSelect configured with the given aggregations.
func (mq *MemberQuery) Aggregate(fns ...AggregateFunc) *MemberSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MemberQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !member.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MemberQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Member, error) {
	var (
		nodes       = []*Member{}
		_spec       = mq.querySpec()
		loadedTypes = [7]bool{
			mq.withMemberDetails != nil,
			mq.withMemberNotes != nil,
			mq.withMemberOrders != nil,
			mq.withMemberProducts != nil,
			mq.withMemberEntry != nil,
			mq.withMemberContents != nil,
			mq.withMemberFace != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Member).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Member{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withMemberDetails; query != nil {
		if err := mq.loadMemberDetails(ctx, query, nodes,
			func(n *Member) { n.Edges.MemberDetails = []*MemberDetails{} },
			func(n *Member, e *MemberDetails) { n.Edges.MemberDetails = append(n.Edges.MemberDetails, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withMemberNotes; query != nil {
		if err := mq.loadMemberNotes(ctx, query, nodes,
			func(n *Member) { n.Edges.MemberNotes = []*MemberNote{} },
			func(n *Member, e *MemberNote) { n.Edges.MemberNotes = append(n.Edges.MemberNotes, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withMemberOrders; query != nil {
		if err := mq.loadMemberOrders(ctx, query, nodes,
			func(n *Member) { n.Edges.MemberOrders = []*Order{} },
			func(n *Member, e *Order) { n.Edges.MemberOrders = append(n.Edges.MemberOrders, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withMemberProducts; query != nil {
		if err := mq.loadMemberProducts(ctx, query, nodes,
			func(n *Member) { n.Edges.MemberProducts = []*MemberProduct{} },
			func(n *Member, e *MemberProduct) { n.Edges.MemberProducts = append(n.Edges.MemberProducts, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withMemberEntry; query != nil {
		if err := mq.loadMemberEntry(ctx, query, nodes,
			func(n *Member) { n.Edges.MemberEntry = []*EntryLogs{} },
			func(n *Member, e *EntryLogs) { n.Edges.MemberEntry = append(n.Edges.MemberEntry, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withMemberContents; query != nil {
		if err := mq.loadMemberContents(ctx, query, nodes,
			func(n *Member) { n.Edges.MemberContents = []*MemberContract{} },
			func(n *Member, e *MemberContract) { n.Edges.MemberContents = append(n.Edges.MemberContents, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withMemberFace; query != nil {
		if err := mq.loadMemberFace(ctx, query, nodes,
			func(n *Member) { n.Edges.MemberFace = []*Face{} },
			func(n *Member, e *Face) { n.Edges.MemberFace = append(n.Edges.MemberFace, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MemberQuery) loadMemberDetails(ctx context.Context, query *MemberDetailsQuery, nodes []*Member, init func(*Member), assign func(*Member, *MemberDetails)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(memberdetails.FieldMemberID)
	}
	query.Where(predicate.MemberDetails(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.MemberDetailsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadMemberNotes(ctx context.Context, query *MemberNoteQuery, nodes []*Member, init func(*Member), assign func(*Member, *MemberNote)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(membernote.FieldMemberID)
	}
	query.Where(predicate.MemberNote(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.MemberNotesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadMemberOrders(ctx context.Context, query *OrderQuery, nodes []*Member, init func(*Member), assign func(*Member, *Order)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(order.FieldMemberID)
	}
	query.Where(predicate.Order(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.MemberOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadMemberProducts(ctx context.Context, query *MemberProductQuery, nodes []*Member, init func(*Member), assign func(*Member, *MemberProduct)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(memberproduct.FieldMemberID)
	}
	query.Where(predicate.MemberProduct(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.MemberProductsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadMemberEntry(ctx context.Context, query *EntryLogsQuery, nodes []*Member, init func(*Member), assign func(*Member, *EntryLogs)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(entrylogs.FieldMemberID)
	}
	query.Where(predicate.EntryLogs(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.MemberEntryColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadMemberContents(ctx context.Context, query *MemberContractQuery, nodes []*Member, init func(*Member), assign func(*Member, *MemberContract)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(membercontract.FieldMemberID)
	}
	query.Where(predicate.MemberContract(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.MemberContentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadMemberFace(ctx context.Context, query *FaceQuery, nodes []*Member, init func(*Member), assign func(*Member, *Face)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(face.FieldMemberID)
	}
	query.Where(predicate.Face(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.MemberFaceColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MemberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MemberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for i := range fields {
			if fields[i] != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MemberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(member.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = member.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MemberGroupBy is the group-by builder for Member entities.
type MemberGroupBy struct {
	selector
	build *MemberQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MemberGroupBy) Aggregate(fns ...AggregateFunc) *MemberGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MemberGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, ent.OpQueryGroupBy)
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberQuery, *MemberGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MemberGroupBy) sqlScan(ctx context.Context, root *MemberQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MemberSelect is the builder for selecting fields of Member entities.
type MemberSelect struct {
	*MemberQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MemberSelect) Aggregate(fns ...AggregateFunc) *MemberSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MemberSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, ent.OpQuerySelect)
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberQuery, *MemberSelect](ctx, ms.MemberQuery, ms, ms.inters, v)
}

func (ms *MemberSelect) sqlScan(ctx context.Context, root *MemberQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
