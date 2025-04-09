// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	entorder "kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/dal/db/mysql/ent/orderpay"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// OrderPayUpdate is the builder for updating OrderPay entities.
type OrderPayUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderPayMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderPayUpdate builder.
func (opu *OrderPayUpdate) Where(ps ...predicate.OrderPay) *OrderPayUpdate {
	opu.mutation.Where(ps...)
	return opu
}

// SetUpdatedAt sets the "updated_at" field.
func (opu *OrderPayUpdate) SetUpdatedAt(t time.Time) *OrderPayUpdate {
	opu.mutation.SetUpdatedAt(t)
	return opu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (opu *OrderPayUpdate) ClearUpdatedAt() *OrderPayUpdate {
	opu.mutation.ClearUpdatedAt()
	return opu
}

// SetDelete sets the "delete" field.
func (opu *OrderPayUpdate) SetDelete(i int64) *OrderPayUpdate {
	opu.mutation.ResetDelete()
	opu.mutation.SetDelete(i)
	return opu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillableDelete(i *int64) *OrderPayUpdate {
	if i != nil {
		opu.SetDelete(*i)
	}
	return opu
}

// AddDelete adds i to the "delete" field.
func (opu *OrderPayUpdate) AddDelete(i int64) *OrderPayUpdate {
	opu.mutation.AddDelete(i)
	return opu
}

// ClearDelete clears the value of the "delete" field.
func (opu *OrderPayUpdate) ClearDelete() *OrderPayUpdate {
	opu.mutation.ClearDelete()
	return opu
}

// SetCreatedID sets the "created_id" field.
func (opu *OrderPayUpdate) SetCreatedID(i int64) *OrderPayUpdate {
	opu.mutation.ResetCreatedID()
	opu.mutation.SetCreatedID(i)
	return opu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillableCreatedID(i *int64) *OrderPayUpdate {
	if i != nil {
		opu.SetCreatedID(*i)
	}
	return opu
}

// AddCreatedID adds i to the "created_id" field.
func (opu *OrderPayUpdate) AddCreatedID(i int64) *OrderPayUpdate {
	opu.mutation.AddCreatedID(i)
	return opu
}

// ClearCreatedID clears the value of the "created_id" field.
func (opu *OrderPayUpdate) ClearCreatedID() *OrderPayUpdate {
	opu.mutation.ClearCreatedID()
	return opu
}

// SetOrderID sets the "order_id" field.
func (opu *OrderPayUpdate) SetOrderID(i int64) *OrderPayUpdate {
	opu.mutation.SetOrderID(i)
	return opu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillableOrderID(i *int64) *OrderPayUpdate {
	if i != nil {
		opu.SetOrderID(*i)
	}
	return opu
}

// ClearOrderID clears the value of the "order_id" field.
func (opu *OrderPayUpdate) ClearOrderID() *OrderPayUpdate {
	opu.mutation.ClearOrderID()
	return opu
}

// SetRemission sets the "remission" field.
func (opu *OrderPayUpdate) SetRemission(f float64) *OrderPayUpdate {
	opu.mutation.ResetRemission()
	opu.mutation.SetRemission(f)
	return opu
}

// SetNillableRemission sets the "remission" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillableRemission(f *float64) *OrderPayUpdate {
	if f != nil {
		opu.SetRemission(*f)
	}
	return opu
}

// AddRemission adds f to the "remission" field.
func (opu *OrderPayUpdate) AddRemission(f float64) *OrderPayUpdate {
	opu.mutation.AddRemission(f)
	return opu
}

// ClearRemission clears the value of the "remission" field.
func (opu *OrderPayUpdate) ClearRemission() *OrderPayUpdate {
	opu.mutation.ClearRemission()
	return opu
}

// SetPay sets the "pay" field.
func (opu *OrderPayUpdate) SetPay(f float64) *OrderPayUpdate {
	opu.mutation.ResetPay()
	opu.mutation.SetPay(f)
	return opu
}

// SetNillablePay sets the "pay" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillablePay(f *float64) *OrderPayUpdate {
	if f != nil {
		opu.SetPay(*f)
	}
	return opu
}

// AddPay adds f to the "pay" field.
func (opu *OrderPayUpdate) AddPay(f float64) *OrderPayUpdate {
	opu.mutation.AddPay(f)
	return opu
}

// ClearPay clears the value of the "pay" field.
func (opu *OrderPayUpdate) ClearPay() *OrderPayUpdate {
	opu.mutation.ClearPay()
	return opu
}

// SetNote sets the "note" field.
func (opu *OrderPayUpdate) SetNote(s string) *OrderPayUpdate {
	opu.mutation.SetNote(s)
	return opu
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillableNote(s *string) *OrderPayUpdate {
	if s != nil {
		opu.SetNote(*s)
	}
	return opu
}

// ClearNote clears the value of the "note" field.
func (opu *OrderPayUpdate) ClearNote() *OrderPayUpdate {
	opu.mutation.ClearNote()
	return opu
}

// SetPayAt sets the "pay_at" field.
func (opu *OrderPayUpdate) SetPayAt(t time.Time) *OrderPayUpdate {
	opu.mutation.SetPayAt(t)
	return opu
}

// SetNillablePayAt sets the "pay_at" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillablePayAt(t *time.Time) *OrderPayUpdate {
	if t != nil {
		opu.SetPayAt(*t)
	}
	return opu
}

// ClearPayAt clears the value of the "pay_at" field.
func (opu *OrderPayUpdate) ClearPayAt() *OrderPayUpdate {
	opu.mutation.ClearPayAt()
	return opu
}

// SetPayWay sets the "pay_way" field.
func (opu *OrderPayUpdate) SetPayWay(s string) *OrderPayUpdate {
	opu.mutation.SetPayWay(s)
	return opu
}

// SetNillablePayWay sets the "pay_way" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillablePayWay(s *string) *OrderPayUpdate {
	if s != nil {
		opu.SetPayWay(*s)
	}
	return opu
}

// ClearPayWay clears the value of the "pay_way" field.
func (opu *OrderPayUpdate) ClearPayWay() *OrderPayUpdate {
	opu.mutation.ClearPayWay()
	return opu
}

// SetPaySn sets the "pay_sn" field.
func (opu *OrderPayUpdate) SetPaySn(s string) *OrderPayUpdate {
	opu.mutation.SetPaySn(s)
	return opu
}

// SetNillablePaySn sets the "pay_sn" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillablePaySn(s *string) *OrderPayUpdate {
	if s != nil {
		opu.SetPaySn(*s)
	}
	return opu
}

// ClearPaySn clears the value of the "pay_sn" field.
func (opu *OrderPayUpdate) ClearPaySn() *OrderPayUpdate {
	opu.mutation.ClearPaySn()
	return opu
}

// SetPrepayID sets the "prepay_id" field.
func (opu *OrderPayUpdate) SetPrepayID(s string) *OrderPayUpdate {
	opu.mutation.SetPrepayID(s)
	return opu
}

// SetNillablePrepayID sets the "prepay_id" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillablePrepayID(s *string) *OrderPayUpdate {
	if s != nil {
		opu.SetPrepayID(*s)
	}
	return opu
}

// ClearPrepayID clears the value of the "prepay_id" field.
func (opu *OrderPayUpdate) ClearPrepayID() *OrderPayUpdate {
	opu.mutation.ClearPrepayID()
	return opu
}

// SetPayExtra sets the "pay_extra" field.
func (opu *OrderPayUpdate) SetPayExtra(u []uint8) *OrderPayUpdate {
	opu.mutation.SetPayExtra(u)
	return opu
}

// AppendPayExtra appends u to the "pay_extra" field.
func (opu *OrderPayUpdate) AppendPayExtra(u []uint8) *OrderPayUpdate {
	opu.mutation.AppendPayExtra(u)
	return opu
}

// ClearPayExtra clears the value of the "pay_extra" field.
func (opu *OrderPayUpdate) ClearPayExtra() *OrderPayUpdate {
	opu.mutation.ClearPayExtra()
	return opu
}

// SetOrder sets the "order" edge to the Order entity.
func (opu *OrderPayUpdate) SetOrder(o *Order) *OrderPayUpdate {
	return opu.SetOrderID(o.ID)
}

// Mutation returns the OrderPayMutation object of the builder.
func (opu *OrderPayUpdate) Mutation() *OrderPayMutation {
	return opu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (opu *OrderPayUpdate) ClearOrder() *OrderPayUpdate {
	opu.mutation.ClearOrder()
	return opu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (opu *OrderPayUpdate) Save(ctx context.Context) (int, error) {
	opu.defaults()
	return withHooks(ctx, opu.sqlSave, opu.mutation, opu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opu *OrderPayUpdate) SaveX(ctx context.Context) int {
	affected, err := opu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (opu *OrderPayUpdate) Exec(ctx context.Context) error {
	_, err := opu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opu *OrderPayUpdate) ExecX(ctx context.Context) {
	if err := opu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opu *OrderPayUpdate) defaults() {
	if _, ok := opu.mutation.UpdatedAt(); !ok && !opu.mutation.UpdatedAtCleared() {
		v := orderpay.UpdateDefaultUpdatedAt()
		opu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (opu *OrderPayUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderPayUpdate {
	opu.modifiers = append(opu.modifiers, modifiers...)
	return opu
}

func (opu *OrderPayUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderpay.Table, orderpay.Columns, sqlgraph.NewFieldSpec(orderpay.FieldID, field.TypeInt64))
	if ps := opu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if opu.mutation.CreatedAtCleared() {
		_spec.ClearField(orderpay.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := opu.mutation.UpdatedAt(); ok {
		_spec.SetField(orderpay.FieldUpdatedAt, field.TypeTime, value)
	}
	if opu.mutation.UpdatedAtCleared() {
		_spec.ClearField(orderpay.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := opu.mutation.Delete(); ok {
		_spec.SetField(orderpay.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := opu.mutation.AddedDelete(); ok {
		_spec.AddField(orderpay.FieldDelete, field.TypeInt64, value)
	}
	if opu.mutation.DeleteCleared() {
		_spec.ClearField(orderpay.FieldDelete, field.TypeInt64)
	}
	if value, ok := opu.mutation.CreatedID(); ok {
		_spec.SetField(orderpay.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := opu.mutation.AddedCreatedID(); ok {
		_spec.AddField(orderpay.FieldCreatedID, field.TypeInt64, value)
	}
	if opu.mutation.CreatedIDCleared() {
		_spec.ClearField(orderpay.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := opu.mutation.Remission(); ok {
		_spec.SetField(orderpay.FieldRemission, field.TypeFloat64, value)
	}
	if value, ok := opu.mutation.AddedRemission(); ok {
		_spec.AddField(orderpay.FieldRemission, field.TypeFloat64, value)
	}
	if opu.mutation.RemissionCleared() {
		_spec.ClearField(orderpay.FieldRemission, field.TypeFloat64)
	}
	if value, ok := opu.mutation.Pay(); ok {
		_spec.SetField(orderpay.FieldPay, field.TypeFloat64, value)
	}
	if value, ok := opu.mutation.AddedPay(); ok {
		_spec.AddField(orderpay.FieldPay, field.TypeFloat64, value)
	}
	if opu.mutation.PayCleared() {
		_spec.ClearField(orderpay.FieldPay, field.TypeFloat64)
	}
	if value, ok := opu.mutation.Note(); ok {
		_spec.SetField(orderpay.FieldNote, field.TypeString, value)
	}
	if opu.mutation.NoteCleared() {
		_spec.ClearField(orderpay.FieldNote, field.TypeString)
	}
	if value, ok := opu.mutation.PayAt(); ok {
		_spec.SetField(orderpay.FieldPayAt, field.TypeTime, value)
	}
	if opu.mutation.PayAtCleared() {
		_spec.ClearField(orderpay.FieldPayAt, field.TypeTime)
	}
	if value, ok := opu.mutation.PayWay(); ok {
		_spec.SetField(orderpay.FieldPayWay, field.TypeString, value)
	}
	if opu.mutation.PayWayCleared() {
		_spec.ClearField(orderpay.FieldPayWay, field.TypeString)
	}
	if value, ok := opu.mutation.PaySn(); ok {
		_spec.SetField(orderpay.FieldPaySn, field.TypeString, value)
	}
	if opu.mutation.PaySnCleared() {
		_spec.ClearField(orderpay.FieldPaySn, field.TypeString)
	}
	if value, ok := opu.mutation.PrepayID(); ok {
		_spec.SetField(orderpay.FieldPrepayID, field.TypeString, value)
	}
	if opu.mutation.PrepayIDCleared() {
		_spec.ClearField(orderpay.FieldPrepayID, field.TypeString)
	}
	if value, ok := opu.mutation.PayExtra(); ok {
		_spec.SetField(orderpay.FieldPayExtra, field.TypeJSON, value)
	}
	if value, ok := opu.mutation.AppendedPayExtra(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orderpay.FieldPayExtra, value)
		})
	}
	if opu.mutation.PayExtraCleared() {
		_spec.ClearField(orderpay.FieldPayExtra, field.TypeJSON)
	}
	if opu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderpay.OrderTable,
			Columns: []string{orderpay.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderpay.OrderTable,
			Columns: []string{orderpay.OrderColumn},
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
	_spec.AddModifiers(opu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, opu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderpay.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	opu.mutation.done = true
	return n, nil
}

// OrderPayUpdateOne is the builder for updating a single OrderPay entity.
type OrderPayUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderPayMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (opuo *OrderPayUpdateOne) SetUpdatedAt(t time.Time) *OrderPayUpdateOne {
	opuo.mutation.SetUpdatedAt(t)
	return opuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (opuo *OrderPayUpdateOne) ClearUpdatedAt() *OrderPayUpdateOne {
	opuo.mutation.ClearUpdatedAt()
	return opuo
}

// SetDelete sets the "delete" field.
func (opuo *OrderPayUpdateOne) SetDelete(i int64) *OrderPayUpdateOne {
	opuo.mutation.ResetDelete()
	opuo.mutation.SetDelete(i)
	return opuo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillableDelete(i *int64) *OrderPayUpdateOne {
	if i != nil {
		opuo.SetDelete(*i)
	}
	return opuo
}

// AddDelete adds i to the "delete" field.
func (opuo *OrderPayUpdateOne) AddDelete(i int64) *OrderPayUpdateOne {
	opuo.mutation.AddDelete(i)
	return opuo
}

// ClearDelete clears the value of the "delete" field.
func (opuo *OrderPayUpdateOne) ClearDelete() *OrderPayUpdateOne {
	opuo.mutation.ClearDelete()
	return opuo
}

// SetCreatedID sets the "created_id" field.
func (opuo *OrderPayUpdateOne) SetCreatedID(i int64) *OrderPayUpdateOne {
	opuo.mutation.ResetCreatedID()
	opuo.mutation.SetCreatedID(i)
	return opuo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillableCreatedID(i *int64) *OrderPayUpdateOne {
	if i != nil {
		opuo.SetCreatedID(*i)
	}
	return opuo
}

// AddCreatedID adds i to the "created_id" field.
func (opuo *OrderPayUpdateOne) AddCreatedID(i int64) *OrderPayUpdateOne {
	opuo.mutation.AddCreatedID(i)
	return opuo
}

// ClearCreatedID clears the value of the "created_id" field.
func (opuo *OrderPayUpdateOne) ClearCreatedID() *OrderPayUpdateOne {
	opuo.mutation.ClearCreatedID()
	return opuo
}

// SetOrderID sets the "order_id" field.
func (opuo *OrderPayUpdateOne) SetOrderID(i int64) *OrderPayUpdateOne {
	opuo.mutation.SetOrderID(i)
	return opuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillableOrderID(i *int64) *OrderPayUpdateOne {
	if i != nil {
		opuo.SetOrderID(*i)
	}
	return opuo
}

// ClearOrderID clears the value of the "order_id" field.
func (opuo *OrderPayUpdateOne) ClearOrderID() *OrderPayUpdateOne {
	opuo.mutation.ClearOrderID()
	return opuo
}

// SetRemission sets the "remission" field.
func (opuo *OrderPayUpdateOne) SetRemission(f float64) *OrderPayUpdateOne {
	opuo.mutation.ResetRemission()
	opuo.mutation.SetRemission(f)
	return opuo
}

// SetNillableRemission sets the "remission" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillableRemission(f *float64) *OrderPayUpdateOne {
	if f != nil {
		opuo.SetRemission(*f)
	}
	return opuo
}

// AddRemission adds f to the "remission" field.
func (opuo *OrderPayUpdateOne) AddRemission(f float64) *OrderPayUpdateOne {
	opuo.mutation.AddRemission(f)
	return opuo
}

// ClearRemission clears the value of the "remission" field.
func (opuo *OrderPayUpdateOne) ClearRemission() *OrderPayUpdateOne {
	opuo.mutation.ClearRemission()
	return opuo
}

// SetPay sets the "pay" field.
func (opuo *OrderPayUpdateOne) SetPay(f float64) *OrderPayUpdateOne {
	opuo.mutation.ResetPay()
	opuo.mutation.SetPay(f)
	return opuo
}

// SetNillablePay sets the "pay" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillablePay(f *float64) *OrderPayUpdateOne {
	if f != nil {
		opuo.SetPay(*f)
	}
	return opuo
}

// AddPay adds f to the "pay" field.
func (opuo *OrderPayUpdateOne) AddPay(f float64) *OrderPayUpdateOne {
	opuo.mutation.AddPay(f)
	return opuo
}

// ClearPay clears the value of the "pay" field.
func (opuo *OrderPayUpdateOne) ClearPay() *OrderPayUpdateOne {
	opuo.mutation.ClearPay()
	return opuo
}

// SetNote sets the "note" field.
func (opuo *OrderPayUpdateOne) SetNote(s string) *OrderPayUpdateOne {
	opuo.mutation.SetNote(s)
	return opuo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillableNote(s *string) *OrderPayUpdateOne {
	if s != nil {
		opuo.SetNote(*s)
	}
	return opuo
}

// ClearNote clears the value of the "note" field.
func (opuo *OrderPayUpdateOne) ClearNote() *OrderPayUpdateOne {
	opuo.mutation.ClearNote()
	return opuo
}

// SetPayAt sets the "pay_at" field.
func (opuo *OrderPayUpdateOne) SetPayAt(t time.Time) *OrderPayUpdateOne {
	opuo.mutation.SetPayAt(t)
	return opuo
}

// SetNillablePayAt sets the "pay_at" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillablePayAt(t *time.Time) *OrderPayUpdateOne {
	if t != nil {
		opuo.SetPayAt(*t)
	}
	return opuo
}

// ClearPayAt clears the value of the "pay_at" field.
func (opuo *OrderPayUpdateOne) ClearPayAt() *OrderPayUpdateOne {
	opuo.mutation.ClearPayAt()
	return opuo
}

// SetPayWay sets the "pay_way" field.
func (opuo *OrderPayUpdateOne) SetPayWay(s string) *OrderPayUpdateOne {
	opuo.mutation.SetPayWay(s)
	return opuo
}

// SetNillablePayWay sets the "pay_way" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillablePayWay(s *string) *OrderPayUpdateOne {
	if s != nil {
		opuo.SetPayWay(*s)
	}
	return opuo
}

// ClearPayWay clears the value of the "pay_way" field.
func (opuo *OrderPayUpdateOne) ClearPayWay() *OrderPayUpdateOne {
	opuo.mutation.ClearPayWay()
	return opuo
}

// SetPaySn sets the "pay_sn" field.
func (opuo *OrderPayUpdateOne) SetPaySn(s string) *OrderPayUpdateOne {
	opuo.mutation.SetPaySn(s)
	return opuo
}

// SetNillablePaySn sets the "pay_sn" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillablePaySn(s *string) *OrderPayUpdateOne {
	if s != nil {
		opuo.SetPaySn(*s)
	}
	return opuo
}

// ClearPaySn clears the value of the "pay_sn" field.
func (opuo *OrderPayUpdateOne) ClearPaySn() *OrderPayUpdateOne {
	opuo.mutation.ClearPaySn()
	return opuo
}

// SetPrepayID sets the "prepay_id" field.
func (opuo *OrderPayUpdateOne) SetPrepayID(s string) *OrderPayUpdateOne {
	opuo.mutation.SetPrepayID(s)
	return opuo
}

// SetNillablePrepayID sets the "prepay_id" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillablePrepayID(s *string) *OrderPayUpdateOne {
	if s != nil {
		opuo.SetPrepayID(*s)
	}
	return opuo
}

// ClearPrepayID clears the value of the "prepay_id" field.
func (opuo *OrderPayUpdateOne) ClearPrepayID() *OrderPayUpdateOne {
	opuo.mutation.ClearPrepayID()
	return opuo
}

// SetPayExtra sets the "pay_extra" field.
func (opuo *OrderPayUpdateOne) SetPayExtra(u []uint8) *OrderPayUpdateOne {
	opuo.mutation.SetPayExtra(u)
	return opuo
}

// AppendPayExtra appends u to the "pay_extra" field.
func (opuo *OrderPayUpdateOne) AppendPayExtra(u []uint8) *OrderPayUpdateOne {
	opuo.mutation.AppendPayExtra(u)
	return opuo
}

// ClearPayExtra clears the value of the "pay_extra" field.
func (opuo *OrderPayUpdateOne) ClearPayExtra() *OrderPayUpdateOne {
	opuo.mutation.ClearPayExtra()
	return opuo
}

// SetOrder sets the "order" edge to the Order entity.
func (opuo *OrderPayUpdateOne) SetOrder(o *Order) *OrderPayUpdateOne {
	return opuo.SetOrderID(o.ID)
}

// Mutation returns the OrderPayMutation object of the builder.
func (opuo *OrderPayUpdateOne) Mutation() *OrderPayMutation {
	return opuo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (opuo *OrderPayUpdateOne) ClearOrder() *OrderPayUpdateOne {
	opuo.mutation.ClearOrder()
	return opuo
}

// Where appends a list predicates to the OrderPayUpdate builder.
func (opuo *OrderPayUpdateOne) Where(ps ...predicate.OrderPay) *OrderPayUpdateOne {
	opuo.mutation.Where(ps...)
	return opuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (opuo *OrderPayUpdateOne) Select(field string, fields ...string) *OrderPayUpdateOne {
	opuo.fields = append([]string{field}, fields...)
	return opuo
}

// Save executes the query and returns the updated OrderPay entity.
func (opuo *OrderPayUpdateOne) Save(ctx context.Context) (*OrderPay, error) {
	opuo.defaults()
	return withHooks(ctx, opuo.sqlSave, opuo.mutation, opuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opuo *OrderPayUpdateOne) SaveX(ctx context.Context) *OrderPay {
	node, err := opuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (opuo *OrderPayUpdateOne) Exec(ctx context.Context) error {
	_, err := opuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opuo *OrderPayUpdateOne) ExecX(ctx context.Context) {
	if err := opuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opuo *OrderPayUpdateOne) defaults() {
	if _, ok := opuo.mutation.UpdatedAt(); !ok && !opuo.mutation.UpdatedAtCleared() {
		v := orderpay.UpdateDefaultUpdatedAt()
		opuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (opuo *OrderPayUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderPayUpdateOne {
	opuo.modifiers = append(opuo.modifiers, modifiers...)
	return opuo
}

func (opuo *OrderPayUpdateOne) sqlSave(ctx context.Context) (_node *OrderPay, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderpay.Table, orderpay.Columns, sqlgraph.NewFieldSpec(orderpay.FieldID, field.TypeInt64))
	id, ok := opuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderPay.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := opuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderpay.FieldID)
		for _, f := range fields {
			if !orderpay.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderpay.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := opuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if opuo.mutation.CreatedAtCleared() {
		_spec.ClearField(orderpay.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := opuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orderpay.FieldUpdatedAt, field.TypeTime, value)
	}
	if opuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(orderpay.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := opuo.mutation.Delete(); ok {
		_spec.SetField(orderpay.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := opuo.mutation.AddedDelete(); ok {
		_spec.AddField(orderpay.FieldDelete, field.TypeInt64, value)
	}
	if opuo.mutation.DeleteCleared() {
		_spec.ClearField(orderpay.FieldDelete, field.TypeInt64)
	}
	if value, ok := opuo.mutation.CreatedID(); ok {
		_spec.SetField(orderpay.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := opuo.mutation.AddedCreatedID(); ok {
		_spec.AddField(orderpay.FieldCreatedID, field.TypeInt64, value)
	}
	if opuo.mutation.CreatedIDCleared() {
		_spec.ClearField(orderpay.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := opuo.mutation.Remission(); ok {
		_spec.SetField(orderpay.FieldRemission, field.TypeFloat64, value)
	}
	if value, ok := opuo.mutation.AddedRemission(); ok {
		_spec.AddField(orderpay.FieldRemission, field.TypeFloat64, value)
	}
	if opuo.mutation.RemissionCleared() {
		_spec.ClearField(orderpay.FieldRemission, field.TypeFloat64)
	}
	if value, ok := opuo.mutation.Pay(); ok {
		_spec.SetField(orderpay.FieldPay, field.TypeFloat64, value)
	}
	if value, ok := opuo.mutation.AddedPay(); ok {
		_spec.AddField(orderpay.FieldPay, field.TypeFloat64, value)
	}
	if opuo.mutation.PayCleared() {
		_spec.ClearField(orderpay.FieldPay, field.TypeFloat64)
	}
	if value, ok := opuo.mutation.Note(); ok {
		_spec.SetField(orderpay.FieldNote, field.TypeString, value)
	}
	if opuo.mutation.NoteCleared() {
		_spec.ClearField(orderpay.FieldNote, field.TypeString)
	}
	if value, ok := opuo.mutation.PayAt(); ok {
		_spec.SetField(orderpay.FieldPayAt, field.TypeTime, value)
	}
	if opuo.mutation.PayAtCleared() {
		_spec.ClearField(orderpay.FieldPayAt, field.TypeTime)
	}
	if value, ok := opuo.mutation.PayWay(); ok {
		_spec.SetField(orderpay.FieldPayWay, field.TypeString, value)
	}
	if opuo.mutation.PayWayCleared() {
		_spec.ClearField(orderpay.FieldPayWay, field.TypeString)
	}
	if value, ok := opuo.mutation.PaySn(); ok {
		_spec.SetField(orderpay.FieldPaySn, field.TypeString, value)
	}
	if opuo.mutation.PaySnCleared() {
		_spec.ClearField(orderpay.FieldPaySn, field.TypeString)
	}
	if value, ok := opuo.mutation.PrepayID(); ok {
		_spec.SetField(orderpay.FieldPrepayID, field.TypeString, value)
	}
	if opuo.mutation.PrepayIDCleared() {
		_spec.ClearField(orderpay.FieldPrepayID, field.TypeString)
	}
	if value, ok := opuo.mutation.PayExtra(); ok {
		_spec.SetField(orderpay.FieldPayExtra, field.TypeJSON, value)
	}
	if value, ok := opuo.mutation.AppendedPayExtra(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orderpay.FieldPayExtra, value)
		})
	}
	if opuo.mutation.PayExtraCleared() {
		_spec.ClearField(orderpay.FieldPayExtra, field.TypeJSON)
	}
	if opuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderpay.OrderTable,
			Columns: []string{orderpay.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderpay.OrderTable,
			Columns: []string{orderpay.OrderColumn},
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
	_spec.AddModifiers(opuo.modifiers...)
	_node = &OrderPay{config: opuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, opuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderpay.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	opuo.mutation.done = true
	return _node, nil
}
