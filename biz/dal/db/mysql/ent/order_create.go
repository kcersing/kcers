// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/membercontract"
	entorder "kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/dal/db/mysql/ent/orderamount"
	"kcers/biz/dal/db/mysql/ent/orderitem"
	"kcers/biz/dal/db/mysql/ent/orderpay"
	"kcers/biz/dal/db/mysql/ent/ordersales"
	"kcers/biz/dal/db/mysql/ent/user"
	"kcers/biz/dal/db/mysql/ent/venue"
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

// SetCreatedAt sets the "created_at" field.
func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrderCreate) SetUpdatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetDelete sets the "delete" field.
func (oc *OrderCreate) SetDelete(i int64) *OrderCreate {
	oc.mutation.SetDelete(i)
	return oc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (oc *OrderCreate) SetNillableDelete(i *int64) *OrderCreate {
	if i != nil {
		oc.SetDelete(*i)
	}
	return oc
}

// SetCreatedID sets the "created_id" field.
func (oc *OrderCreate) SetCreatedID(i int64) *OrderCreate {
	oc.mutation.SetCreatedID(i)
	return oc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedID(i *int64) *OrderCreate {
	if i != nil {
		oc.SetCreatedID(*i)
	}
	return oc
}

// SetOrderSn sets the "order_sn" field.
func (oc *OrderCreate) SetOrderSn(s string) *OrderCreate {
	oc.mutation.SetOrderSn(s)
	return oc
}

// SetNillableOrderSn sets the "order_sn" field if the given value is not nil.
func (oc *OrderCreate) SetNillableOrderSn(s *string) *OrderCreate {
	if s != nil {
		oc.SetOrderSn(*s)
	}
	return oc
}

// SetVenueID sets the "venue_id" field.
func (oc *OrderCreate) SetVenueID(i int64) *OrderCreate {
	oc.mutation.SetVenueID(i)
	return oc
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableVenueID(i *int64) *OrderCreate {
	if i != nil {
		oc.SetVenueID(*i)
	}
	return oc
}

// SetMemberID sets the "member_id" field.
func (oc *OrderCreate) SetMemberID(i int64) *OrderCreate {
	oc.mutation.SetMemberID(i)
	return oc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableMemberID(i *int64) *OrderCreate {
	if i != nil {
		oc.SetMemberID(*i)
	}
	return oc
}

// SetMemberProductID sets the "member_product_id" field.
func (oc *OrderCreate) SetMemberProductID(i int64) *OrderCreate {
	oc.mutation.SetMemberProductID(i)
	return oc
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableMemberProductID(i *int64) *OrderCreate {
	if i != nil {
		oc.SetMemberProductID(*i)
	}
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(i int64) *OrderCreate {
	oc.mutation.SetStatus(i)
	return oc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oc *OrderCreate) SetNillableStatus(i *int64) *OrderCreate {
	if i != nil {
		oc.SetStatus(*i)
	}
	return oc
}

// SetSource sets the "source" field.
func (oc *OrderCreate) SetSource(s string) *OrderCreate {
	oc.mutation.SetSource(s)
	return oc
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (oc *OrderCreate) SetNillableSource(s *string) *OrderCreate {
	if s != nil {
		oc.SetSource(*s)
	}
	return oc
}

// SetDevice sets the "device" field.
func (oc *OrderCreate) SetDevice(s string) *OrderCreate {
	oc.mutation.SetDevice(s)
	return oc
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (oc *OrderCreate) SetNillableDevice(s *string) *OrderCreate {
	if s != nil {
		oc.SetDevice(*s)
	}
	return oc
}

// SetNature sets the "nature" field.
func (oc *OrderCreate) SetNature(i int64) *OrderCreate {
	oc.mutation.SetNature(i)
	return oc
}

// SetNillableNature sets the "nature" field if the given value is not nil.
func (oc *OrderCreate) SetNillableNature(i *int64) *OrderCreate {
	if i != nil {
		oc.SetNature(*i)
	}
	return oc
}

// SetCompletionAt sets the "completion_at" field.
func (oc *OrderCreate) SetCompletionAt(t time.Time) *OrderCreate {
	oc.mutation.SetCompletionAt(t)
	return oc
}

// SetNillableCompletionAt sets the "completion_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCompletionAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCompletionAt(*t)
	}
	return oc
}

// SetCloseAt sets the "close_at" field.
func (oc *OrderCreate) SetCloseAt(t time.Time) *OrderCreate {
	oc.mutation.SetCloseAt(t)
	return oc
}

// SetNillableCloseAt sets the "close_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCloseAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCloseAt(*t)
	}
	return oc
}

// SetRefundAt sets the "refund_at" field.
func (oc *OrderCreate) SetRefundAt(t time.Time) *OrderCreate {
	oc.mutation.SetRefundAt(t)
	return oc
}

// SetNillableRefundAt sets the "refund_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableRefundAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetRefundAt(*t)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OrderCreate) SetID(i int64) *OrderCreate {
	oc.mutation.SetID(i)
	return oc
}

// AddAmountIDs adds the "amount" edge to the OrderAmount entity by IDs.
func (oc *OrderCreate) AddAmountIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddAmountIDs(ids...)
	return oc
}

// AddAmount adds the "amount" edges to the OrderAmount entity.
func (oc *OrderCreate) AddAmount(o ...*OrderAmount) *OrderCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddAmountIDs(ids...)
}

// AddItemIDs adds the "item" edge to the OrderItem entity by IDs.
func (oc *OrderCreate) AddItemIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddItemIDs(ids...)
	return oc
}

// AddItem adds the "item" edges to the OrderItem entity.
func (oc *OrderCreate) AddItem(o ...*OrderItem) *OrderCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddItemIDs(ids...)
}

// AddPayIDs adds the "pay" edge to the OrderPay entity by IDs.
func (oc *OrderCreate) AddPayIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddPayIDs(ids...)
	return oc
}

// AddPay adds the "pay" edges to the OrderPay entity.
func (oc *OrderCreate) AddPay(o ...*OrderPay) *OrderCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddPayIDs(ids...)
}

// AddOrderContentIDs adds the "order_contents" edge to the MemberContract entity by IDs.
func (oc *OrderCreate) AddOrderContentIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddOrderContentIDs(ids...)
	return oc
}

// AddOrderContents adds the "order_contents" edges to the MemberContract entity.
func (oc *OrderCreate) AddOrderContents(m ...*MemberContract) *OrderCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return oc.AddOrderContentIDs(ids...)
}

// AddSaleIDs adds the "sales" edge to the OrderSales entity by IDs.
func (oc *OrderCreate) AddSaleIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddSaleIDs(ids...)
	return oc
}

// AddSales adds the "sales" edges to the OrderSales entity.
func (oc *OrderCreate) AddSales(o ...*OrderSales) *OrderCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddSaleIDs(ids...)
}

// SetOrderVenuesID sets the "order_venues" edge to the Venue entity by ID.
func (oc *OrderCreate) SetOrderVenuesID(id int64) *OrderCreate {
	oc.mutation.SetOrderVenuesID(id)
	return oc
}

// SetNillableOrderVenuesID sets the "order_venues" edge to the Venue entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableOrderVenuesID(id *int64) *OrderCreate {
	if id != nil {
		oc = oc.SetOrderVenuesID(*id)
	}
	return oc
}

// SetOrderVenues sets the "order_venues" edge to the Venue entity.
func (oc *OrderCreate) SetOrderVenues(v *Venue) *OrderCreate {
	return oc.SetOrderVenuesID(v.ID)
}

// SetOrderMembersID sets the "order_members" edge to the Member entity by ID.
func (oc *OrderCreate) SetOrderMembersID(id int64) *OrderCreate {
	oc.mutation.SetOrderMembersID(id)
	return oc
}

// SetNillableOrderMembersID sets the "order_members" edge to the Member entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableOrderMembersID(id *int64) *OrderCreate {
	if id != nil {
		oc = oc.SetOrderMembersID(*id)
	}
	return oc
}

// SetOrderMembers sets the "order_members" edge to the Member entity.
func (oc *OrderCreate) SetOrderMembers(m *Member) *OrderCreate {
	return oc.SetOrderMembersID(m.ID)
}

// SetOrderCreatesID sets the "order_creates" edge to the User entity by ID.
func (oc *OrderCreate) SetOrderCreatesID(id int64) *OrderCreate {
	oc.mutation.SetOrderCreatesID(id)
	return oc
}

// SetNillableOrderCreatesID sets the "order_creates" edge to the User entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableOrderCreatesID(id *int64) *OrderCreate {
	if id != nil {
		oc = oc.SetOrderCreatesID(*id)
	}
	return oc
}

// SetOrderCreates sets the "order_creates" edge to the User entity.
func (oc *OrderCreate) SetOrderCreates(u *User) *OrderCreate {
	return oc.SetOrderCreatesID(u.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
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
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := entorder.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := entorder.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oc.mutation.Delete(); !ok {
		v := entorder.DefaultDelete
		oc.mutation.SetDelete(v)
	}
	if _, ok := oc.mutation.CreatedID(); !ok {
		v := entorder.DefaultCreatedID
		oc.mutation.SetCreatedID(v)
	}
	if _, ok := oc.mutation.Status(); !ok {
		v := entorder.DefaultStatus
		oc.mutation.SetStatus(v)
	}
	if _, ok := oc.mutation.Source(); !ok {
		v := entorder.DefaultSource
		oc.mutation.SetSource(v)
	}
	if _, ok := oc.mutation.Device(); !ok {
		v := entorder.DefaultDevice
		oc.mutation.SetDevice(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(entorder.Table, sqlgraph.NewFieldSpec(entorder.FieldID, field.TypeInt64))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(entorder.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(entorder.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.Delete(); ok {
		_spec.SetField(entorder.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := oc.mutation.OrderSn(); ok {
		_spec.SetField(entorder.FieldOrderSn, field.TypeString, value)
		_node.OrderSn = value
	}
	if value, ok := oc.mutation.MemberProductID(); ok {
		_spec.SetField(entorder.FieldMemberProductID, field.TypeInt64, value)
		_node.MemberProductID = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(entorder.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.Source(); ok {
		_spec.SetField(entorder.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if value, ok := oc.mutation.Device(); ok {
		_spec.SetField(entorder.FieldDevice, field.TypeString, value)
		_node.Device = value
	}
	if value, ok := oc.mutation.Nature(); ok {
		_spec.SetField(entorder.FieldNature, field.TypeInt64, value)
		_node.Nature = value
	}
	if value, ok := oc.mutation.CompletionAt(); ok {
		_spec.SetField(entorder.FieldCompletionAt, field.TypeTime, value)
		_node.CompletionAt = value
	}
	if value, ok := oc.mutation.CloseAt(); ok {
		_spec.SetField(entorder.FieldCloseAt, field.TypeTime, value)
		_node.CloseAt = value
	}
	if value, ok := oc.mutation.RefundAt(); ok {
		_spec.SetField(entorder.FieldRefundAt, field.TypeTime, value)
		_node.RefundAt = value
	}
	if nodes := oc.mutation.AmountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entorder.AmountTable,
			Columns: []string{entorder.AmountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderamount.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entorder.ItemTable,
			Columns: []string{entorder.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.PayIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entorder.PayTable,
			Columns: []string{entorder.PayColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderpay.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrderContentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entorder.OrderContentsTable,
			Columns: []string{entorder.OrderContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontract.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.SalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entorder.SalesTable,
			Columns: []string{entorder.SalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ordersales.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrderVenuesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entorder.OrderVenuesTable,
			Columns: []string{entorder.OrderVenuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.VenueID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrderMembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entorder.OrderMembersTable,
			Columns: []string{entorder.OrderMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MemberID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrderCreatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entorder.OrderCreatesTable,
			Columns: []string{entorder.OrderCreatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreatedID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	err      error
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
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
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
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
