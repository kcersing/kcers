// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/dal/db/mysql/ent/orderamount"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderAmountCreate is the builder for creating a OrderAmount entity.
type OrderAmountCreate struct {
	config
	mutation *OrderAmountMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oac *OrderAmountCreate) SetCreatedAt(t time.Time) *OrderAmountCreate {
	oac.mutation.SetCreatedAt(t)
	return oac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oac *OrderAmountCreate) SetNillableCreatedAt(t *time.Time) *OrderAmountCreate {
	if t != nil {
		oac.SetCreatedAt(*t)
	}
	return oac
}

// SetUpdatedAt sets the "updated_at" field.
func (oac *OrderAmountCreate) SetUpdatedAt(t time.Time) *OrderAmountCreate {
	oac.mutation.SetUpdatedAt(t)
	return oac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oac *OrderAmountCreate) SetNillableUpdatedAt(t *time.Time) *OrderAmountCreate {
	if t != nil {
		oac.SetUpdatedAt(*t)
	}
	return oac
}

// SetOrderID sets the "order_id" field.
func (oac *OrderAmountCreate) SetOrderID(i int64) *OrderAmountCreate {
	oac.mutation.SetOrderID(i)
	return oac
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (oac *OrderAmountCreate) SetNillableOrderID(i *int64) *OrderAmountCreate {
	if i != nil {
		oac.SetOrderID(*i)
	}
	return oac
}

// SetTotal sets the "total" field.
func (oac *OrderAmountCreate) SetTotal(f float64) *OrderAmountCreate {
	oac.mutation.SetTotal(f)
	return oac
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (oac *OrderAmountCreate) SetNillableTotal(f *float64) *OrderAmountCreate {
	if f != nil {
		oac.SetTotal(*f)
	}
	return oac
}

// SetActual sets the "actual" field.
func (oac *OrderAmountCreate) SetActual(f float64) *OrderAmountCreate {
	oac.mutation.SetActual(f)
	return oac
}

// SetNillableActual sets the "actual" field if the given value is not nil.
func (oac *OrderAmountCreate) SetNillableActual(f *float64) *OrderAmountCreate {
	if f != nil {
		oac.SetActual(*f)
	}
	return oac
}

// SetResidue sets the "residue" field.
func (oac *OrderAmountCreate) SetResidue(f float64) *OrderAmountCreate {
	oac.mutation.SetResidue(f)
	return oac
}

// SetNillableResidue sets the "residue" field if the given value is not nil.
func (oac *OrderAmountCreate) SetNillableResidue(f *float64) *OrderAmountCreate {
	if f != nil {
		oac.SetResidue(*f)
	}
	return oac
}

// SetRemission sets the "remission" field.
func (oac *OrderAmountCreate) SetRemission(f float64) *OrderAmountCreate {
	oac.mutation.SetRemission(f)
	return oac
}

// SetNillableRemission sets the "remission" field if the given value is not nil.
func (oac *OrderAmountCreate) SetNillableRemission(f *float64) *OrderAmountCreate {
	if f != nil {
		oac.SetRemission(*f)
	}
	return oac
}

// SetID sets the "id" field.
func (oac *OrderAmountCreate) SetID(i int64) *OrderAmountCreate {
	oac.mutation.SetID(i)
	return oac
}

// SetOrder sets the "order" edge to the Order entity.
func (oac *OrderAmountCreate) SetOrder(o *Order) *OrderAmountCreate {
	return oac.SetOrderID(o.ID)
}

// Mutation returns the OrderAmountMutation object of the builder.
func (oac *OrderAmountCreate) Mutation() *OrderAmountMutation {
	return oac.mutation
}

// Save creates the OrderAmount in the database.
func (oac *OrderAmountCreate) Save(ctx context.Context) (*OrderAmount, error) {
	oac.defaults()
	return withHooks(ctx, oac.sqlSave, oac.mutation, oac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oac *OrderAmountCreate) SaveX(ctx context.Context) *OrderAmount {
	v, err := oac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oac *OrderAmountCreate) Exec(ctx context.Context) error {
	_, err := oac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oac *OrderAmountCreate) ExecX(ctx context.Context) {
	if err := oac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oac *OrderAmountCreate) defaults() {
	if _, ok := oac.mutation.CreatedAt(); !ok {
		v := orderamount.DefaultCreatedAt()
		oac.mutation.SetCreatedAt(v)
	}
	if _, ok := oac.mutation.UpdatedAt(); !ok {
		v := orderamount.DefaultUpdatedAt()
		oac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oac *OrderAmountCreate) check() error {
	if _, ok := oac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderAmount.created_at"`)}
	}
	if _, ok := oac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderAmount.updated_at"`)}
	}
	if v, ok := oac.mutation.OrderID(); ok {
		if err := orderamount.OrderIDValidator(v); err != nil {
			return &ValidationError{Name: "order_id", err: fmt.Errorf(`ent: validator failed for field "OrderAmount.order_id": %w`, err)}
		}
	}
	return nil
}

func (oac *OrderAmountCreate) sqlSave(ctx context.Context) (*OrderAmount, error) {
	if err := oac.check(); err != nil {
		return nil, err
	}
	_node, _spec := oac.createSpec()
	if err := sqlgraph.CreateNode(ctx, oac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	oac.mutation.id = &_node.ID
	oac.mutation.done = true
	return _node, nil
}

func (oac *OrderAmountCreate) createSpec() (*OrderAmount, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderAmount{config: oac.config}
		_spec = sqlgraph.NewCreateSpec(orderamount.Table, sqlgraph.NewFieldSpec(orderamount.FieldID, field.TypeInt64))
	)
	if id, ok := oac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oac.mutation.CreatedAt(); ok {
		_spec.SetField(orderamount.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oac.mutation.UpdatedAt(); ok {
		_spec.SetField(orderamount.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oac.mutation.Total(); ok {
		_spec.SetField(orderamount.FieldTotal, field.TypeFloat64, value)
		_node.Total = value
	}
	if value, ok := oac.mutation.Actual(); ok {
		_spec.SetField(orderamount.FieldActual, field.TypeFloat64, value)
		_node.Actual = value
	}
	if value, ok := oac.mutation.Residue(); ok {
		_spec.SetField(orderamount.FieldResidue, field.TypeFloat64, value)
		_node.Residue = value
	}
	if value, ok := oac.mutation.Remission(); ok {
		_spec.SetField(orderamount.FieldRemission, field.TypeFloat64, value)
		_node.Remission = value
	}
	if nodes := oac.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderamount.OrderTable,
			Columns: []string{orderamount.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderAmountCreateBulk is the builder for creating many OrderAmount entities in bulk.
type OrderAmountCreateBulk struct {
	config
	err      error
	builders []*OrderAmountCreate
}

// Save creates the OrderAmount entities in the database.
func (oacb *OrderAmountCreateBulk) Save(ctx context.Context) ([]*OrderAmount, error) {
	if oacb.err != nil {
		return nil, oacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oacb.builders))
	nodes := make([]*OrderAmount, len(oacb.builders))
	mutators := make([]Mutator, len(oacb.builders))
	for i := range oacb.builders {
		func(i int, root context.Context) {
			builder := oacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderAmountMutation)
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
					_, err = mutators[i+1].Mutate(root, oacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oacb *OrderAmountCreateBulk) SaveX(ctx context.Context) []*OrderAmount {
	v, err := oacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oacb *OrderAmountCreateBulk) Exec(ctx context.Context) error {
	_, err := oacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oacb *OrderAmountCreateBulk) ExecX(ctx context.Context) {
	if err := oacb.Exec(ctx); err != nil {
		panic(err)
	}
}
