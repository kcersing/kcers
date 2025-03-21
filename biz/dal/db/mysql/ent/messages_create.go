// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/messages"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessagesCreate is the builder for creating a Messages entity.
type MessagesCreate struct {
	config
	mutation *MessagesMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MessagesCreate) SetCreatedAt(t time.Time) *MessagesCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MessagesCreate) SetNillableCreatedAt(t *time.Time) *MessagesCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MessagesCreate) SetUpdatedAt(t time.Time) *MessagesCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MessagesCreate) SetNillableUpdatedAt(t *time.Time) *MessagesCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetType sets the "type" field.
func (mc *MessagesCreate) SetType(s string) *MessagesCreate {
	mc.mutation.SetType(s)
	return mc
}

// SetToUserID sets the "to_user_id" field.
func (mc *MessagesCreate) SetToUserID(s string) *MessagesCreate {
	mc.mutation.SetToUserID(s)
	return mc
}

// SetFromUserID sets the "from_user_id" field.
func (mc *MessagesCreate) SetFromUserID(s string) *MessagesCreate {
	mc.mutation.SetFromUserID(s)
	return mc
}

// SetContent sets the "content" field.
func (mc *MessagesCreate) SetContent(s string) *MessagesCreate {
	mc.mutation.SetContent(s)
	return mc
}

// SetID sets the "id" field.
func (mc *MessagesCreate) SetID(i int64) *MessagesCreate {
	mc.mutation.SetID(i)
	return mc
}

// Mutation returns the MessagesMutation object of the builder.
func (mc *MessagesCreate) Mutation() *MessagesMutation {
	return mc.mutation
}

// Save creates the Messages in the database.
func (mc *MessagesCreate) Save(ctx context.Context) (*Messages, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessagesCreate) SaveX(ctx context.Context) *Messages {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessagesCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessagesCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessagesCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := messages.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := messages.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessagesCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Messages.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Messages.updated_at"`)}
	}
	if _, ok := mc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Messages.type"`)}
	}
	if _, ok := mc.mutation.ToUserID(); !ok {
		return &ValidationError{Name: "to_user_id", err: errors.New(`ent: missing required field "Messages.to_user_id"`)}
	}
	if _, ok := mc.mutation.FromUserID(); !ok {
		return &ValidationError{Name: "from_user_id", err: errors.New(`ent: missing required field "Messages.from_user_id"`)}
	}
	if _, ok := mc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Messages.content"`)}
	}
	return nil
}

func (mc *MessagesCreate) sqlSave(ctx context.Context) (*Messages, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MessagesCreate) createSpec() (*Messages, *sqlgraph.CreateSpec) {
	var (
		_node = &Messages{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(messages.Table, sqlgraph.NewFieldSpec(messages.FieldID, field.TypeInt64))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(messages.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(messages.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.GetType(); ok {
		_spec.SetField(messages.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := mc.mutation.ToUserID(); ok {
		_spec.SetField(messages.FieldToUserID, field.TypeString, value)
		_node.ToUserID = value
	}
	if value, ok := mc.mutation.FromUserID(); ok {
		_spec.SetField(messages.FieldFromUserID, field.TypeString, value)
		_node.FromUserID = value
	}
	if value, ok := mc.mutation.Content(); ok {
		_spec.SetField(messages.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	return _node, _spec
}

// MessagesCreateBulk is the builder for creating many Messages entities in bulk.
type MessagesCreateBulk struct {
	config
	err      error
	builders []*MessagesCreate
}

// Save creates the Messages entities in the database.
func (mcb *MessagesCreateBulk) Save(ctx context.Context) ([]*Messages, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Messages, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessagesMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessagesCreateBulk) SaveX(ctx context.Context) []*Messages {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessagesCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessagesCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
