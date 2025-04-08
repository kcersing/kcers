// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/dal/db/mysql/ent/orderitem"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderItemUpdate is the builder for updating OrderItem entities.
type OrderItemUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiu *OrderItemUpdate) Where(ps ...predicate.OrderItem) *OrderItemUpdate {
	oiu.mutation.Where(ps...)
	return oiu
}

// SetUpdatedAt sets the "updated_at" field.
func (oiu *OrderItemUpdate) SetUpdatedAt(t time.Time) *OrderItemUpdate {
	oiu.mutation.SetUpdatedAt(t)
	return oiu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (oiu *OrderItemUpdate) ClearUpdatedAt() *OrderItemUpdate {
	oiu.mutation.ClearUpdatedAt()
	return oiu
}

// SetDelete sets the "delete" field.
func (oiu *OrderItemUpdate) SetDelete(i int64) *OrderItemUpdate {
	oiu.mutation.ResetDelete()
	oiu.mutation.SetDelete(i)
	return oiu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableDelete(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetDelete(*i)
	}
	return oiu
}

// AddDelete adds i to the "delete" field.
func (oiu *OrderItemUpdate) AddDelete(i int64) *OrderItemUpdate {
	oiu.mutation.AddDelete(i)
	return oiu
}

// ClearDelete clears the value of the "delete" field.
func (oiu *OrderItemUpdate) ClearDelete() *OrderItemUpdate {
	oiu.mutation.ClearDelete()
	return oiu
}

// SetCreatedID sets the "created_id" field.
func (oiu *OrderItemUpdate) SetCreatedID(i int64) *OrderItemUpdate {
	oiu.mutation.ResetCreatedID()
	oiu.mutation.SetCreatedID(i)
	return oiu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableCreatedID(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetCreatedID(*i)
	}
	return oiu
}

// AddCreatedID adds i to the "created_id" field.
func (oiu *OrderItemUpdate) AddCreatedID(i int64) *OrderItemUpdate {
	oiu.mutation.AddCreatedID(i)
	return oiu
}

// ClearCreatedID clears the value of the "created_id" field.
func (oiu *OrderItemUpdate) ClearCreatedID() *OrderItemUpdate {
	oiu.mutation.ClearCreatedID()
	return oiu
}

// SetOrderID sets the "order_id" field.
func (oiu *OrderItemUpdate) SetOrderID(i int64) *OrderItemUpdate {
	oiu.mutation.SetOrderID(i)
	return oiu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableOrderID(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetOrderID(*i)
	}
	return oiu
}

// ClearOrderID clears the value of the "order_id" field.
func (oiu *OrderItemUpdate) ClearOrderID() *OrderItemUpdate {
	oiu.mutation.ClearOrderID()
	return oiu
}

// SetProductID sets the "product_id" field.
func (oiu *OrderItemUpdate) SetProductID(i int64) *OrderItemUpdate {
	oiu.mutation.ResetProductID()
	oiu.mutation.SetProductID(i)
	return oiu
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableProductID(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetProductID(*i)
	}
	return oiu
}

// AddProductID adds i to the "product_id" field.
func (oiu *OrderItemUpdate) AddProductID(i int64) *OrderItemUpdate {
	oiu.mutation.AddProductID(i)
	return oiu
}

// ClearProductID clears the value of the "product_id" field.
func (oiu *OrderItemUpdate) ClearProductID() *OrderItemUpdate {
	oiu.mutation.ClearProductID()
	return oiu
}

// SetRelatedUserProductID sets the "related_user_product_id" field.
func (oiu *OrderItemUpdate) SetRelatedUserProductID(i int64) *OrderItemUpdate {
	oiu.mutation.ResetRelatedUserProductID()
	oiu.mutation.SetRelatedUserProductID(i)
	return oiu
}

// SetNillableRelatedUserProductID sets the "related_user_product_id" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableRelatedUserProductID(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetRelatedUserProductID(*i)
	}
	return oiu
}

// AddRelatedUserProductID adds i to the "related_user_product_id" field.
func (oiu *OrderItemUpdate) AddRelatedUserProductID(i int64) *OrderItemUpdate {
	oiu.mutation.AddRelatedUserProductID(i)
	return oiu
}

// ClearRelatedUserProductID clears the value of the "related_user_product_id" field.
func (oiu *OrderItemUpdate) ClearRelatedUserProductID() *OrderItemUpdate {
	oiu.mutation.ClearRelatedUserProductID()
	return oiu
}

// SetOrder sets the "order" edge to the Order entity.
func (oiu *OrderItemUpdate) SetOrder(o *Order) *OrderItemUpdate {
	return oiu.SetOrderID(o.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiu *OrderItemUpdate) Mutation() *OrderItemMutation {
	return oiu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (oiu *OrderItemUpdate) ClearOrder() *OrderItemUpdate {
	oiu.mutation.ClearOrder()
	return oiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oiu *OrderItemUpdate) Save(ctx context.Context) (int, error) {
	oiu.defaults()
	return withHooks(ctx, oiu.sqlSave, oiu.mutation, oiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiu *OrderItemUpdate) SaveX(ctx context.Context) int {
	affected, err := oiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oiu *OrderItemUpdate) Exec(ctx context.Context) error {
	_, err := oiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiu *OrderItemUpdate) ExecX(ctx context.Context) {
	if err := oiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oiu *OrderItemUpdate) defaults() {
	if _, ok := oiu.mutation.UpdatedAt(); !ok && !oiu.mutation.UpdatedAtCleared() {
		v := orderitem.UpdateDefaultUpdatedAt()
		oiu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (oiu *OrderItemUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderItemUpdate {
	oiu.modifiers = append(oiu.modifiers, modifiers...)
	return oiu
}

func (oiu *OrderItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt64))
	if ps := oiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if oiu.mutation.CreatedAtCleared() {
		_spec.ClearField(orderitem.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := oiu.mutation.UpdatedAt(); ok {
		_spec.SetField(orderitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if oiu.mutation.UpdatedAtCleared() {
		_spec.ClearField(orderitem.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := oiu.mutation.Delete(); ok {
		_spec.SetField(orderitem.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := oiu.mutation.AddedDelete(); ok {
		_spec.AddField(orderitem.FieldDelete, field.TypeInt64, value)
	}
	if oiu.mutation.DeleteCleared() {
		_spec.ClearField(orderitem.FieldDelete, field.TypeInt64)
	}
	if value, ok := oiu.mutation.CreatedID(); ok {
		_spec.SetField(orderitem.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := oiu.mutation.AddedCreatedID(); ok {
		_spec.AddField(orderitem.FieldCreatedID, field.TypeInt64, value)
	}
	if oiu.mutation.CreatedIDCleared() {
		_spec.ClearField(orderitem.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := oiu.mutation.ProductID(); ok {
		_spec.SetField(orderitem.FieldProductID, field.TypeInt64, value)
	}
	if value, ok := oiu.mutation.AddedProductID(); ok {
		_spec.AddField(orderitem.FieldProductID, field.TypeInt64, value)
	}
	if oiu.mutation.ProductIDCleared() {
		_spec.ClearField(orderitem.FieldProductID, field.TypeInt64)
	}
	if value, ok := oiu.mutation.RelatedUserProductID(); ok {
		_spec.SetField(orderitem.FieldRelatedUserProductID, field.TypeInt64, value)
	}
	if value, ok := oiu.mutation.AddedRelatedUserProductID(); ok {
		_spec.AddField(orderitem.FieldRelatedUserProductID, field.TypeInt64, value)
	}
	if oiu.mutation.RelatedUserProductIDCleared() {
		_spec.ClearField(orderitem.FieldRelatedUserProductID, field.TypeInt64)
	}
	if oiu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(oiu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, oiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oiu.mutation.done = true
	return n, nil
}

// OrderItemUpdateOne is the builder for updating a single OrderItem entity.
type OrderItemUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (oiuo *OrderItemUpdateOne) SetUpdatedAt(t time.Time) *OrderItemUpdateOne {
	oiuo.mutation.SetUpdatedAt(t)
	return oiuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (oiuo *OrderItemUpdateOne) ClearUpdatedAt() *OrderItemUpdateOne {
	oiuo.mutation.ClearUpdatedAt()
	return oiuo
}

// SetDelete sets the "delete" field.
func (oiuo *OrderItemUpdateOne) SetDelete(i int64) *OrderItemUpdateOne {
	oiuo.mutation.ResetDelete()
	oiuo.mutation.SetDelete(i)
	return oiuo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableDelete(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetDelete(*i)
	}
	return oiuo
}

// AddDelete adds i to the "delete" field.
func (oiuo *OrderItemUpdateOne) AddDelete(i int64) *OrderItemUpdateOne {
	oiuo.mutation.AddDelete(i)
	return oiuo
}

// ClearDelete clears the value of the "delete" field.
func (oiuo *OrderItemUpdateOne) ClearDelete() *OrderItemUpdateOne {
	oiuo.mutation.ClearDelete()
	return oiuo
}

// SetCreatedID sets the "created_id" field.
func (oiuo *OrderItemUpdateOne) SetCreatedID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.ResetCreatedID()
	oiuo.mutation.SetCreatedID(i)
	return oiuo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableCreatedID(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetCreatedID(*i)
	}
	return oiuo
}

// AddCreatedID adds i to the "created_id" field.
func (oiuo *OrderItemUpdateOne) AddCreatedID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.AddCreatedID(i)
	return oiuo
}

// ClearCreatedID clears the value of the "created_id" field.
func (oiuo *OrderItemUpdateOne) ClearCreatedID() *OrderItemUpdateOne {
	oiuo.mutation.ClearCreatedID()
	return oiuo
}

// SetOrderID sets the "order_id" field.
func (oiuo *OrderItemUpdateOne) SetOrderID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.SetOrderID(i)
	return oiuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableOrderID(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetOrderID(*i)
	}
	return oiuo
}

// ClearOrderID clears the value of the "order_id" field.
func (oiuo *OrderItemUpdateOne) ClearOrderID() *OrderItemUpdateOne {
	oiuo.mutation.ClearOrderID()
	return oiuo
}

// SetProductID sets the "product_id" field.
func (oiuo *OrderItemUpdateOne) SetProductID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.ResetProductID()
	oiuo.mutation.SetProductID(i)
	return oiuo
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableProductID(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetProductID(*i)
	}
	return oiuo
}

// AddProductID adds i to the "product_id" field.
func (oiuo *OrderItemUpdateOne) AddProductID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.AddProductID(i)
	return oiuo
}

// ClearProductID clears the value of the "product_id" field.
func (oiuo *OrderItemUpdateOne) ClearProductID() *OrderItemUpdateOne {
	oiuo.mutation.ClearProductID()
	return oiuo
}

// SetRelatedUserProductID sets the "related_user_product_id" field.
func (oiuo *OrderItemUpdateOne) SetRelatedUserProductID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.ResetRelatedUserProductID()
	oiuo.mutation.SetRelatedUserProductID(i)
	return oiuo
}

// SetNillableRelatedUserProductID sets the "related_user_product_id" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableRelatedUserProductID(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetRelatedUserProductID(*i)
	}
	return oiuo
}

// AddRelatedUserProductID adds i to the "related_user_product_id" field.
func (oiuo *OrderItemUpdateOne) AddRelatedUserProductID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.AddRelatedUserProductID(i)
	return oiuo
}

// ClearRelatedUserProductID clears the value of the "related_user_product_id" field.
func (oiuo *OrderItemUpdateOne) ClearRelatedUserProductID() *OrderItemUpdateOne {
	oiuo.mutation.ClearRelatedUserProductID()
	return oiuo
}

// SetOrder sets the "order" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) SetOrder(o *Order) *OrderItemUpdateOne {
	return oiuo.SetOrderID(o.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiuo *OrderItemUpdateOne) Mutation() *OrderItemMutation {
	return oiuo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) ClearOrder() *OrderItemUpdateOne {
	oiuo.mutation.ClearOrder()
	return oiuo
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiuo *OrderItemUpdateOne) Where(ps ...predicate.OrderItem) *OrderItemUpdateOne {
	oiuo.mutation.Where(ps...)
	return oiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oiuo *OrderItemUpdateOne) Select(field string, fields ...string) *OrderItemUpdateOne {
	oiuo.fields = append([]string{field}, fields...)
	return oiuo
}

// Save executes the query and returns the updated OrderItem entity.
func (oiuo *OrderItemUpdateOne) Save(ctx context.Context) (*OrderItem, error) {
	oiuo.defaults()
	return withHooks(ctx, oiuo.sqlSave, oiuo.mutation, oiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) SaveX(ctx context.Context) *OrderItem {
	node, err := oiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oiuo *OrderItemUpdateOne) Exec(ctx context.Context) error {
	_, err := oiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) ExecX(ctx context.Context) {
	if err := oiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oiuo *OrderItemUpdateOne) defaults() {
	if _, ok := oiuo.mutation.UpdatedAt(); !ok && !oiuo.mutation.UpdatedAtCleared() {
		v := orderitem.UpdateDefaultUpdatedAt()
		oiuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (oiuo *OrderItemUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderItemUpdateOne {
	oiuo.modifiers = append(oiuo.modifiers, modifiers...)
	return oiuo
}

func (oiuo *OrderItemUpdateOne) sqlSave(ctx context.Context) (_node *OrderItem, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt64))
	id, ok := oiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.FieldID)
		for _, f := range fields {
			if !orderitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if oiuo.mutation.CreatedAtCleared() {
		_spec.ClearField(orderitem.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := oiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orderitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if oiuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(orderitem.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := oiuo.mutation.Delete(); ok {
		_spec.SetField(orderitem.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := oiuo.mutation.AddedDelete(); ok {
		_spec.AddField(orderitem.FieldDelete, field.TypeInt64, value)
	}
	if oiuo.mutation.DeleteCleared() {
		_spec.ClearField(orderitem.FieldDelete, field.TypeInt64)
	}
	if value, ok := oiuo.mutation.CreatedID(); ok {
		_spec.SetField(orderitem.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := oiuo.mutation.AddedCreatedID(); ok {
		_spec.AddField(orderitem.FieldCreatedID, field.TypeInt64, value)
	}
	if oiuo.mutation.CreatedIDCleared() {
		_spec.ClearField(orderitem.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := oiuo.mutation.ProductID(); ok {
		_spec.SetField(orderitem.FieldProductID, field.TypeInt64, value)
	}
	if value, ok := oiuo.mutation.AddedProductID(); ok {
		_spec.AddField(orderitem.FieldProductID, field.TypeInt64, value)
	}
	if oiuo.mutation.ProductIDCleared() {
		_spec.ClearField(orderitem.FieldProductID, field.TypeInt64)
	}
	if value, ok := oiuo.mutation.RelatedUserProductID(); ok {
		_spec.SetField(orderitem.FieldRelatedUserProductID, field.TypeInt64, value)
	}
	if value, ok := oiuo.mutation.AddedRelatedUserProductID(); ok {
		_spec.AddField(orderitem.FieldRelatedUserProductID, field.TypeInt64, value)
	}
	if oiuo.mutation.RelatedUserProductIDCleared() {
		_spec.ClearField(orderitem.FieldRelatedUserProductID, field.TypeInt64)
	}
	if oiuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(oiuo.modifiers...)
	_node = &OrderItem{config: oiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oiuo.mutation.done = true
	return _node, nil
}
