// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/membercontract"
	"kcers/biz/dal/db/mysql/ent/membercontractcontent"
	"kcers/biz/dal/db/mysql/ent/memberproduct"
	entorder "kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberContractUpdate is the builder for updating MemberContract entities.
type MemberContractUpdate struct {
	config
	hooks     []Hook
	mutation  *MemberContractMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MemberContractUpdate builder.
func (mcu *MemberContractUpdate) Where(ps ...predicate.MemberContract) *MemberContractUpdate {
	mcu.mutation.Where(ps...)
	return mcu
}

// SetUpdatedAt sets the "updated_at" field.
func (mcu *MemberContractUpdate) SetUpdatedAt(t time.Time) *MemberContractUpdate {
	mcu.mutation.SetUpdatedAt(t)
	return mcu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mcu *MemberContractUpdate) ClearUpdatedAt() *MemberContractUpdate {
	mcu.mutation.ClearUpdatedAt()
	return mcu
}

// SetDelete sets the "delete" field.
func (mcu *MemberContractUpdate) SetDelete(i int64) *MemberContractUpdate {
	mcu.mutation.ResetDelete()
	mcu.mutation.SetDelete(i)
	return mcu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableDelete(i *int64) *MemberContractUpdate {
	if i != nil {
		mcu.SetDelete(*i)
	}
	return mcu
}

// AddDelete adds i to the "delete" field.
func (mcu *MemberContractUpdate) AddDelete(i int64) *MemberContractUpdate {
	mcu.mutation.AddDelete(i)
	return mcu
}

// ClearDelete clears the value of the "delete" field.
func (mcu *MemberContractUpdate) ClearDelete() *MemberContractUpdate {
	mcu.mutation.ClearDelete()
	return mcu
}

// SetCreatedID sets the "created_id" field.
func (mcu *MemberContractUpdate) SetCreatedID(i int64) *MemberContractUpdate {
	mcu.mutation.ResetCreatedID()
	mcu.mutation.SetCreatedID(i)
	return mcu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableCreatedID(i *int64) *MemberContractUpdate {
	if i != nil {
		mcu.SetCreatedID(*i)
	}
	return mcu
}

// AddCreatedID adds i to the "created_id" field.
func (mcu *MemberContractUpdate) AddCreatedID(i int64) *MemberContractUpdate {
	mcu.mutation.AddCreatedID(i)
	return mcu
}

// ClearCreatedID clears the value of the "created_id" field.
func (mcu *MemberContractUpdate) ClearCreatedID() *MemberContractUpdate {
	mcu.mutation.ClearCreatedID()
	return mcu
}

// SetStatus sets the "status" field.
func (mcu *MemberContractUpdate) SetStatus(i int64) *MemberContractUpdate {
	mcu.mutation.ResetStatus()
	mcu.mutation.SetStatus(i)
	return mcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableStatus(i *int64) *MemberContractUpdate {
	if i != nil {
		mcu.SetStatus(*i)
	}
	return mcu
}

// AddStatus adds i to the "status" field.
func (mcu *MemberContractUpdate) AddStatus(i int64) *MemberContractUpdate {
	mcu.mutation.AddStatus(i)
	return mcu
}

// ClearStatus clears the value of the "status" field.
func (mcu *MemberContractUpdate) ClearStatus() *MemberContractUpdate {
	mcu.mutation.ClearStatus()
	return mcu
}

// SetMemberID sets the "member_id" field.
func (mcu *MemberContractUpdate) SetMemberID(i int64) *MemberContractUpdate {
	mcu.mutation.SetMemberID(i)
	return mcu
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableMemberID(i *int64) *MemberContractUpdate {
	if i != nil {
		mcu.SetMemberID(*i)
	}
	return mcu
}

// ClearMemberID clears the value of the "member_id" field.
func (mcu *MemberContractUpdate) ClearMemberID() *MemberContractUpdate {
	mcu.mutation.ClearMemberID()
	return mcu
}

// SetContractID sets the "contract_id" field.
func (mcu *MemberContractUpdate) SetContractID(i int64) *MemberContractUpdate {
	mcu.mutation.ResetContractID()
	mcu.mutation.SetContractID(i)
	return mcu
}

// SetNillableContractID sets the "contract_id" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableContractID(i *int64) *MemberContractUpdate {
	if i != nil {
		mcu.SetContractID(*i)
	}
	return mcu
}

// AddContractID adds i to the "contract_id" field.
func (mcu *MemberContractUpdate) AddContractID(i int64) *MemberContractUpdate {
	mcu.mutation.AddContractID(i)
	return mcu
}

// ClearContractID clears the value of the "contract_id" field.
func (mcu *MemberContractUpdate) ClearContractID() *MemberContractUpdate {
	mcu.mutation.ClearContractID()
	return mcu
}

// SetOrderID sets the "order_id" field.
func (mcu *MemberContractUpdate) SetOrderID(i int64) *MemberContractUpdate {
	mcu.mutation.SetOrderID(i)
	return mcu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableOrderID(i *int64) *MemberContractUpdate {
	if i != nil {
		mcu.SetOrderID(*i)
	}
	return mcu
}

// ClearOrderID clears the value of the "order_id" field.
func (mcu *MemberContractUpdate) ClearOrderID() *MemberContractUpdate {
	mcu.mutation.ClearOrderID()
	return mcu
}

// SetVenueID sets the "venue_id" field.
func (mcu *MemberContractUpdate) SetVenueID(i int64) *MemberContractUpdate {
	mcu.mutation.ResetVenueID()
	mcu.mutation.SetVenueID(i)
	return mcu
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableVenueID(i *int64) *MemberContractUpdate {
	if i != nil {
		mcu.SetVenueID(*i)
	}
	return mcu
}

// AddVenueID adds i to the "venue_id" field.
func (mcu *MemberContractUpdate) AddVenueID(i int64) *MemberContractUpdate {
	mcu.mutation.AddVenueID(i)
	return mcu
}

// ClearVenueID clears the value of the "venue_id" field.
func (mcu *MemberContractUpdate) ClearVenueID() *MemberContractUpdate {
	mcu.mutation.ClearVenueID()
	return mcu
}

// SetMemberProductID sets the "member_product_id" field.
func (mcu *MemberContractUpdate) SetMemberProductID(i int64) *MemberContractUpdate {
	mcu.mutation.SetMemberProductID(i)
	return mcu
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableMemberProductID(i *int64) *MemberContractUpdate {
	if i != nil {
		mcu.SetMemberProductID(*i)
	}
	return mcu
}

// ClearMemberProductID clears the value of the "member_product_id" field.
func (mcu *MemberContractUpdate) ClearMemberProductID() *MemberContractUpdate {
	mcu.mutation.ClearMemberProductID()
	return mcu
}

// SetName sets the "name" field.
func (mcu *MemberContractUpdate) SetName(s string) *MemberContractUpdate {
	mcu.mutation.SetName(s)
	return mcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableName(s *string) *MemberContractUpdate {
	if s != nil {
		mcu.SetName(*s)
	}
	return mcu
}

// ClearName clears the value of the "name" field.
func (mcu *MemberContractUpdate) ClearName() *MemberContractUpdate {
	mcu.mutation.ClearName()
	return mcu
}

// SetSign sets the "sign" field.
func (mcu *MemberContractUpdate) SetSign(s string) *MemberContractUpdate {
	mcu.mutation.SetSign(s)
	return mcu
}

// SetNillableSign sets the "sign" field if the given value is not nil.
func (mcu *MemberContractUpdate) SetNillableSign(s *string) *MemberContractUpdate {
	if s != nil {
		mcu.SetSign(*s)
	}
	return mcu
}

// ClearSign clears the value of the "sign" field.
func (mcu *MemberContractUpdate) ClearSign() *MemberContractUpdate {
	mcu.mutation.ClearSign()
	return mcu
}

// AddContentIDs adds the "content" edge to the MemberContractContent entity by IDs.
func (mcu *MemberContractUpdate) AddContentIDs(ids ...int64) *MemberContractUpdate {
	mcu.mutation.AddContentIDs(ids...)
	return mcu
}

// AddContent adds the "content" edges to the MemberContractContent entity.
func (mcu *MemberContractUpdate) AddContent(m ...*MemberContractContent) *MemberContractUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcu.AddContentIDs(ids...)
}

// SetMember sets the "member" edge to the Member entity.
func (mcu *MemberContractUpdate) SetMember(m *Member) *MemberContractUpdate {
	return mcu.SetMemberID(m.ID)
}

// SetOrder sets the "order" edge to the Order entity.
func (mcu *MemberContractUpdate) SetOrder(o *Order) *MemberContractUpdate {
	return mcu.SetOrderID(o.ID)
}

// SetMemberProduct sets the "member_product" edge to the MemberProduct entity.
func (mcu *MemberContractUpdate) SetMemberProduct(m *MemberProduct) *MemberContractUpdate {
	return mcu.SetMemberProductID(m.ID)
}

// Mutation returns the MemberContractMutation object of the builder.
func (mcu *MemberContractUpdate) Mutation() *MemberContractMutation {
	return mcu.mutation
}

// ClearContent clears all "content" edges to the MemberContractContent entity.
func (mcu *MemberContractUpdate) ClearContent() *MemberContractUpdate {
	mcu.mutation.ClearContent()
	return mcu
}

// RemoveContentIDs removes the "content" edge to MemberContractContent entities by IDs.
func (mcu *MemberContractUpdate) RemoveContentIDs(ids ...int64) *MemberContractUpdate {
	mcu.mutation.RemoveContentIDs(ids...)
	return mcu
}

// RemoveContent removes "content" edges to MemberContractContent entities.
func (mcu *MemberContractUpdate) RemoveContent(m ...*MemberContractContent) *MemberContractUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcu.RemoveContentIDs(ids...)
}

// ClearMember clears the "member" edge to the Member entity.
func (mcu *MemberContractUpdate) ClearMember() *MemberContractUpdate {
	mcu.mutation.ClearMember()
	return mcu
}

// ClearOrder clears the "order" edge to the Order entity.
func (mcu *MemberContractUpdate) ClearOrder() *MemberContractUpdate {
	mcu.mutation.ClearOrder()
	return mcu
}

// ClearMemberProduct clears the "member_product" edge to the MemberProduct entity.
func (mcu *MemberContractUpdate) ClearMemberProduct() *MemberContractUpdate {
	mcu.mutation.ClearMemberProduct()
	return mcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mcu *MemberContractUpdate) Save(ctx context.Context) (int, error) {
	mcu.defaults()
	return withHooks(ctx, mcu.sqlSave, mcu.mutation, mcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcu *MemberContractUpdate) SaveX(ctx context.Context) int {
	affected, err := mcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mcu *MemberContractUpdate) Exec(ctx context.Context) error {
	_, err := mcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcu *MemberContractUpdate) ExecX(ctx context.Context) {
	if err := mcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcu *MemberContractUpdate) defaults() {
	if _, ok := mcu.mutation.UpdatedAt(); !ok && !mcu.mutation.UpdatedAtCleared() {
		v := membercontract.UpdateDefaultUpdatedAt()
		mcu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mcu *MemberContractUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MemberContractUpdate {
	mcu.modifiers = append(mcu.modifiers, modifiers...)
	return mcu
}

func (mcu *MemberContractUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(membercontract.Table, membercontract.Columns, sqlgraph.NewFieldSpec(membercontract.FieldID, field.TypeInt64))
	if ps := mcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mcu.mutation.CreatedAtCleared() {
		_spec.ClearField(membercontract.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mcu.mutation.UpdatedAt(); ok {
		_spec.SetField(membercontract.FieldUpdatedAt, field.TypeTime, value)
	}
	if mcu.mutation.UpdatedAtCleared() {
		_spec.ClearField(membercontract.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := mcu.mutation.Delete(); ok {
		_spec.SetField(membercontract.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := mcu.mutation.AddedDelete(); ok {
		_spec.AddField(membercontract.FieldDelete, field.TypeInt64, value)
	}
	if mcu.mutation.DeleteCleared() {
		_spec.ClearField(membercontract.FieldDelete, field.TypeInt64)
	}
	if value, ok := mcu.mutation.CreatedID(); ok {
		_spec.SetField(membercontract.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := mcu.mutation.AddedCreatedID(); ok {
		_spec.AddField(membercontract.FieldCreatedID, field.TypeInt64, value)
	}
	if mcu.mutation.CreatedIDCleared() {
		_spec.ClearField(membercontract.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := mcu.mutation.Status(); ok {
		_spec.SetField(membercontract.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := mcu.mutation.AddedStatus(); ok {
		_spec.AddField(membercontract.FieldStatus, field.TypeInt64, value)
	}
	if mcu.mutation.StatusCleared() {
		_spec.ClearField(membercontract.FieldStatus, field.TypeInt64)
	}
	if value, ok := mcu.mutation.ContractID(); ok {
		_spec.SetField(membercontract.FieldContractID, field.TypeInt64, value)
	}
	if value, ok := mcu.mutation.AddedContractID(); ok {
		_spec.AddField(membercontract.FieldContractID, field.TypeInt64, value)
	}
	if mcu.mutation.ContractIDCleared() {
		_spec.ClearField(membercontract.FieldContractID, field.TypeInt64)
	}
	if value, ok := mcu.mutation.VenueID(); ok {
		_spec.SetField(membercontract.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := mcu.mutation.AddedVenueID(); ok {
		_spec.AddField(membercontract.FieldVenueID, field.TypeInt64, value)
	}
	if mcu.mutation.VenueIDCleared() {
		_spec.ClearField(membercontract.FieldVenueID, field.TypeInt64)
	}
	if value, ok := mcu.mutation.Name(); ok {
		_spec.SetField(membercontract.FieldName, field.TypeString, value)
	}
	if mcu.mutation.NameCleared() {
		_spec.ClearField(membercontract.FieldName, field.TypeString)
	}
	if value, ok := mcu.mutation.Sign(); ok {
		_spec.SetField(membercontract.FieldSign, field.TypeString, value)
	}
	if mcu.mutation.SignCleared() {
		_spec.ClearField(membercontract.FieldSign, field.TypeString)
	}
	if mcu.mutation.ContentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   membercontract.ContentTable,
			Columns: []string{membercontract.ContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.RemovedContentIDs(); len(nodes) > 0 && !mcu.mutation.ContentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   membercontract.ContentTable,
			Columns: []string{membercontract.ContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.ContentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   membercontract.ContentTable,
			Columns: []string{membercontract.ContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcu.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.MemberTable,
			Columns: []string{membercontract.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.MemberTable,
			Columns: []string{membercontract.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.OrderTable,
			Columns: []string{membercontract.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.OrderTable,
			Columns: []string{membercontract.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcu.mutation.MemberProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.MemberProductTable,
			Columns: []string{membercontract.MemberProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.MemberProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.MemberProductTable,
			Columns: []string{membercontract.MemberProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membercontract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mcu.mutation.done = true
	return n, nil
}

// MemberContractUpdateOne is the builder for updating a single MemberContract entity.
type MemberContractUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MemberContractMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (mcuo *MemberContractUpdateOne) SetUpdatedAt(t time.Time) *MemberContractUpdateOne {
	mcuo.mutation.SetUpdatedAt(t)
	return mcuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mcuo *MemberContractUpdateOne) ClearUpdatedAt() *MemberContractUpdateOne {
	mcuo.mutation.ClearUpdatedAt()
	return mcuo
}

// SetDelete sets the "delete" field.
func (mcuo *MemberContractUpdateOne) SetDelete(i int64) *MemberContractUpdateOne {
	mcuo.mutation.ResetDelete()
	mcuo.mutation.SetDelete(i)
	return mcuo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableDelete(i *int64) *MemberContractUpdateOne {
	if i != nil {
		mcuo.SetDelete(*i)
	}
	return mcuo
}

// AddDelete adds i to the "delete" field.
func (mcuo *MemberContractUpdateOne) AddDelete(i int64) *MemberContractUpdateOne {
	mcuo.mutation.AddDelete(i)
	return mcuo
}

// ClearDelete clears the value of the "delete" field.
func (mcuo *MemberContractUpdateOne) ClearDelete() *MemberContractUpdateOne {
	mcuo.mutation.ClearDelete()
	return mcuo
}

// SetCreatedID sets the "created_id" field.
func (mcuo *MemberContractUpdateOne) SetCreatedID(i int64) *MemberContractUpdateOne {
	mcuo.mutation.ResetCreatedID()
	mcuo.mutation.SetCreatedID(i)
	return mcuo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableCreatedID(i *int64) *MemberContractUpdateOne {
	if i != nil {
		mcuo.SetCreatedID(*i)
	}
	return mcuo
}

// AddCreatedID adds i to the "created_id" field.
func (mcuo *MemberContractUpdateOne) AddCreatedID(i int64) *MemberContractUpdateOne {
	mcuo.mutation.AddCreatedID(i)
	return mcuo
}

// ClearCreatedID clears the value of the "created_id" field.
func (mcuo *MemberContractUpdateOne) ClearCreatedID() *MemberContractUpdateOne {
	mcuo.mutation.ClearCreatedID()
	return mcuo
}

// SetStatus sets the "status" field.
func (mcuo *MemberContractUpdateOne) SetStatus(i int64) *MemberContractUpdateOne {
	mcuo.mutation.ResetStatus()
	mcuo.mutation.SetStatus(i)
	return mcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableStatus(i *int64) *MemberContractUpdateOne {
	if i != nil {
		mcuo.SetStatus(*i)
	}
	return mcuo
}

// AddStatus adds i to the "status" field.
func (mcuo *MemberContractUpdateOne) AddStatus(i int64) *MemberContractUpdateOne {
	mcuo.mutation.AddStatus(i)
	return mcuo
}

// ClearStatus clears the value of the "status" field.
func (mcuo *MemberContractUpdateOne) ClearStatus() *MemberContractUpdateOne {
	mcuo.mutation.ClearStatus()
	return mcuo
}

// SetMemberID sets the "member_id" field.
func (mcuo *MemberContractUpdateOne) SetMemberID(i int64) *MemberContractUpdateOne {
	mcuo.mutation.SetMemberID(i)
	return mcuo
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableMemberID(i *int64) *MemberContractUpdateOne {
	if i != nil {
		mcuo.SetMemberID(*i)
	}
	return mcuo
}

// ClearMemberID clears the value of the "member_id" field.
func (mcuo *MemberContractUpdateOne) ClearMemberID() *MemberContractUpdateOne {
	mcuo.mutation.ClearMemberID()
	return mcuo
}

// SetContractID sets the "contract_id" field.
func (mcuo *MemberContractUpdateOne) SetContractID(i int64) *MemberContractUpdateOne {
	mcuo.mutation.ResetContractID()
	mcuo.mutation.SetContractID(i)
	return mcuo
}

// SetNillableContractID sets the "contract_id" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableContractID(i *int64) *MemberContractUpdateOne {
	if i != nil {
		mcuo.SetContractID(*i)
	}
	return mcuo
}

// AddContractID adds i to the "contract_id" field.
func (mcuo *MemberContractUpdateOne) AddContractID(i int64) *MemberContractUpdateOne {
	mcuo.mutation.AddContractID(i)
	return mcuo
}

// ClearContractID clears the value of the "contract_id" field.
func (mcuo *MemberContractUpdateOne) ClearContractID() *MemberContractUpdateOne {
	mcuo.mutation.ClearContractID()
	return mcuo
}

// SetOrderID sets the "order_id" field.
func (mcuo *MemberContractUpdateOne) SetOrderID(i int64) *MemberContractUpdateOne {
	mcuo.mutation.SetOrderID(i)
	return mcuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableOrderID(i *int64) *MemberContractUpdateOne {
	if i != nil {
		mcuo.SetOrderID(*i)
	}
	return mcuo
}

// ClearOrderID clears the value of the "order_id" field.
func (mcuo *MemberContractUpdateOne) ClearOrderID() *MemberContractUpdateOne {
	mcuo.mutation.ClearOrderID()
	return mcuo
}

// SetVenueID sets the "venue_id" field.
func (mcuo *MemberContractUpdateOne) SetVenueID(i int64) *MemberContractUpdateOne {
	mcuo.mutation.ResetVenueID()
	mcuo.mutation.SetVenueID(i)
	return mcuo
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableVenueID(i *int64) *MemberContractUpdateOne {
	if i != nil {
		mcuo.SetVenueID(*i)
	}
	return mcuo
}

// AddVenueID adds i to the "venue_id" field.
func (mcuo *MemberContractUpdateOne) AddVenueID(i int64) *MemberContractUpdateOne {
	mcuo.mutation.AddVenueID(i)
	return mcuo
}

// ClearVenueID clears the value of the "venue_id" field.
func (mcuo *MemberContractUpdateOne) ClearVenueID() *MemberContractUpdateOne {
	mcuo.mutation.ClearVenueID()
	return mcuo
}

// SetMemberProductID sets the "member_product_id" field.
func (mcuo *MemberContractUpdateOne) SetMemberProductID(i int64) *MemberContractUpdateOne {
	mcuo.mutation.SetMemberProductID(i)
	return mcuo
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableMemberProductID(i *int64) *MemberContractUpdateOne {
	if i != nil {
		mcuo.SetMemberProductID(*i)
	}
	return mcuo
}

// ClearMemberProductID clears the value of the "member_product_id" field.
func (mcuo *MemberContractUpdateOne) ClearMemberProductID() *MemberContractUpdateOne {
	mcuo.mutation.ClearMemberProductID()
	return mcuo
}

// SetName sets the "name" field.
func (mcuo *MemberContractUpdateOne) SetName(s string) *MemberContractUpdateOne {
	mcuo.mutation.SetName(s)
	return mcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableName(s *string) *MemberContractUpdateOne {
	if s != nil {
		mcuo.SetName(*s)
	}
	return mcuo
}

// ClearName clears the value of the "name" field.
func (mcuo *MemberContractUpdateOne) ClearName() *MemberContractUpdateOne {
	mcuo.mutation.ClearName()
	return mcuo
}

// SetSign sets the "sign" field.
func (mcuo *MemberContractUpdateOne) SetSign(s string) *MemberContractUpdateOne {
	mcuo.mutation.SetSign(s)
	return mcuo
}

// SetNillableSign sets the "sign" field if the given value is not nil.
func (mcuo *MemberContractUpdateOne) SetNillableSign(s *string) *MemberContractUpdateOne {
	if s != nil {
		mcuo.SetSign(*s)
	}
	return mcuo
}

// ClearSign clears the value of the "sign" field.
func (mcuo *MemberContractUpdateOne) ClearSign() *MemberContractUpdateOne {
	mcuo.mutation.ClearSign()
	return mcuo
}

// AddContentIDs adds the "content" edge to the MemberContractContent entity by IDs.
func (mcuo *MemberContractUpdateOne) AddContentIDs(ids ...int64) *MemberContractUpdateOne {
	mcuo.mutation.AddContentIDs(ids...)
	return mcuo
}

// AddContent adds the "content" edges to the MemberContractContent entity.
func (mcuo *MemberContractUpdateOne) AddContent(m ...*MemberContractContent) *MemberContractUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcuo.AddContentIDs(ids...)
}

// SetMember sets the "member" edge to the Member entity.
func (mcuo *MemberContractUpdateOne) SetMember(m *Member) *MemberContractUpdateOne {
	return mcuo.SetMemberID(m.ID)
}

// SetOrder sets the "order" edge to the Order entity.
func (mcuo *MemberContractUpdateOne) SetOrder(o *Order) *MemberContractUpdateOne {
	return mcuo.SetOrderID(o.ID)
}

// SetMemberProduct sets the "member_product" edge to the MemberProduct entity.
func (mcuo *MemberContractUpdateOne) SetMemberProduct(m *MemberProduct) *MemberContractUpdateOne {
	return mcuo.SetMemberProductID(m.ID)
}

// Mutation returns the MemberContractMutation object of the builder.
func (mcuo *MemberContractUpdateOne) Mutation() *MemberContractMutation {
	return mcuo.mutation
}

// ClearContent clears all "content" edges to the MemberContractContent entity.
func (mcuo *MemberContractUpdateOne) ClearContent() *MemberContractUpdateOne {
	mcuo.mutation.ClearContent()
	return mcuo
}

// RemoveContentIDs removes the "content" edge to MemberContractContent entities by IDs.
func (mcuo *MemberContractUpdateOne) RemoveContentIDs(ids ...int64) *MemberContractUpdateOne {
	mcuo.mutation.RemoveContentIDs(ids...)
	return mcuo
}

// RemoveContent removes "content" edges to MemberContractContent entities.
func (mcuo *MemberContractUpdateOne) RemoveContent(m ...*MemberContractContent) *MemberContractUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcuo.RemoveContentIDs(ids...)
}

// ClearMember clears the "member" edge to the Member entity.
func (mcuo *MemberContractUpdateOne) ClearMember() *MemberContractUpdateOne {
	mcuo.mutation.ClearMember()
	return mcuo
}

// ClearOrder clears the "order" edge to the Order entity.
func (mcuo *MemberContractUpdateOne) ClearOrder() *MemberContractUpdateOne {
	mcuo.mutation.ClearOrder()
	return mcuo
}

// ClearMemberProduct clears the "member_product" edge to the MemberProduct entity.
func (mcuo *MemberContractUpdateOne) ClearMemberProduct() *MemberContractUpdateOne {
	mcuo.mutation.ClearMemberProduct()
	return mcuo
}

// Where appends a list predicates to the MemberContractUpdate builder.
func (mcuo *MemberContractUpdateOne) Where(ps ...predicate.MemberContract) *MemberContractUpdateOne {
	mcuo.mutation.Where(ps...)
	return mcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mcuo *MemberContractUpdateOne) Select(field string, fields ...string) *MemberContractUpdateOne {
	mcuo.fields = append([]string{field}, fields...)
	return mcuo
}

// Save executes the query and returns the updated MemberContract entity.
func (mcuo *MemberContractUpdateOne) Save(ctx context.Context) (*MemberContract, error) {
	mcuo.defaults()
	return withHooks(ctx, mcuo.sqlSave, mcuo.mutation, mcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcuo *MemberContractUpdateOne) SaveX(ctx context.Context) *MemberContract {
	node, err := mcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mcuo *MemberContractUpdateOne) Exec(ctx context.Context) error {
	_, err := mcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcuo *MemberContractUpdateOne) ExecX(ctx context.Context) {
	if err := mcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcuo *MemberContractUpdateOne) defaults() {
	if _, ok := mcuo.mutation.UpdatedAt(); !ok && !mcuo.mutation.UpdatedAtCleared() {
		v := membercontract.UpdateDefaultUpdatedAt()
		mcuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mcuo *MemberContractUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MemberContractUpdateOne {
	mcuo.modifiers = append(mcuo.modifiers, modifiers...)
	return mcuo
}

func (mcuo *MemberContractUpdateOne) sqlSave(ctx context.Context) (_node *MemberContract, err error) {
	_spec := sqlgraph.NewUpdateSpec(membercontract.Table, membercontract.Columns, sqlgraph.NewFieldSpec(membercontract.FieldID, field.TypeInt64))
	id, ok := mcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MemberContract.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membercontract.FieldID)
		for _, f := range fields {
			if !membercontract.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != membercontract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mcuo.mutation.CreatedAtCleared() {
		_spec.ClearField(membercontract.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(membercontract.FieldUpdatedAt, field.TypeTime, value)
	}
	if mcuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(membercontract.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := mcuo.mutation.Delete(); ok {
		_spec.SetField(membercontract.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := mcuo.mutation.AddedDelete(); ok {
		_spec.AddField(membercontract.FieldDelete, field.TypeInt64, value)
	}
	if mcuo.mutation.DeleteCleared() {
		_spec.ClearField(membercontract.FieldDelete, field.TypeInt64)
	}
	if value, ok := mcuo.mutation.CreatedID(); ok {
		_spec.SetField(membercontract.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := mcuo.mutation.AddedCreatedID(); ok {
		_spec.AddField(membercontract.FieldCreatedID, field.TypeInt64, value)
	}
	if mcuo.mutation.CreatedIDCleared() {
		_spec.ClearField(membercontract.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := mcuo.mutation.Status(); ok {
		_spec.SetField(membercontract.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := mcuo.mutation.AddedStatus(); ok {
		_spec.AddField(membercontract.FieldStatus, field.TypeInt64, value)
	}
	if mcuo.mutation.StatusCleared() {
		_spec.ClearField(membercontract.FieldStatus, field.TypeInt64)
	}
	if value, ok := mcuo.mutation.ContractID(); ok {
		_spec.SetField(membercontract.FieldContractID, field.TypeInt64, value)
	}
	if value, ok := mcuo.mutation.AddedContractID(); ok {
		_spec.AddField(membercontract.FieldContractID, field.TypeInt64, value)
	}
	if mcuo.mutation.ContractIDCleared() {
		_spec.ClearField(membercontract.FieldContractID, field.TypeInt64)
	}
	if value, ok := mcuo.mutation.VenueID(); ok {
		_spec.SetField(membercontract.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := mcuo.mutation.AddedVenueID(); ok {
		_spec.AddField(membercontract.FieldVenueID, field.TypeInt64, value)
	}
	if mcuo.mutation.VenueIDCleared() {
		_spec.ClearField(membercontract.FieldVenueID, field.TypeInt64)
	}
	if value, ok := mcuo.mutation.Name(); ok {
		_spec.SetField(membercontract.FieldName, field.TypeString, value)
	}
	if mcuo.mutation.NameCleared() {
		_spec.ClearField(membercontract.FieldName, field.TypeString)
	}
	if value, ok := mcuo.mutation.Sign(); ok {
		_spec.SetField(membercontract.FieldSign, field.TypeString, value)
	}
	if mcuo.mutation.SignCleared() {
		_spec.ClearField(membercontract.FieldSign, field.TypeString)
	}
	if mcuo.mutation.ContentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   membercontract.ContentTable,
			Columns: []string{membercontract.ContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.RemovedContentIDs(); len(nodes) > 0 && !mcuo.mutation.ContentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   membercontract.ContentTable,
			Columns: []string{membercontract.ContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.ContentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   membercontract.ContentTable,
			Columns: []string{membercontract.ContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcuo.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.MemberTable,
			Columns: []string{membercontract.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.MemberTable,
			Columns: []string{membercontract.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.OrderTable,
			Columns: []string{membercontract.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.OrderTable,
			Columns: []string{membercontract.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mcuo.mutation.MemberProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.MemberProductTable,
			Columns: []string{membercontract.MemberProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.MemberProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontract.MemberProductTable,
			Columns: []string{membercontract.MemberProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mcuo.modifiers...)
	_node = &MemberContract{config: mcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membercontract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mcuo.mutation.done = true
	return _node, nil
}
