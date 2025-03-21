// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/membercontract"
	"kcers/biz/dal/db/mysql/ent/membercontractcontent"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberContractContentCreate is the builder for creating a MemberContractContent entity.
type MemberContractContentCreate struct {
	config
	mutation *MemberContractContentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mccc *MemberContractContentCreate) SetCreatedAt(t time.Time) *MemberContractContentCreate {
	mccc.mutation.SetCreatedAt(t)
	return mccc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mccc *MemberContractContentCreate) SetNillableCreatedAt(t *time.Time) *MemberContractContentCreate {
	if t != nil {
		mccc.SetCreatedAt(*t)
	}
	return mccc
}

// SetUpdatedAt sets the "updated_at" field.
func (mccc *MemberContractContentCreate) SetUpdatedAt(t time.Time) *MemberContractContentCreate {
	mccc.mutation.SetUpdatedAt(t)
	return mccc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mccc *MemberContractContentCreate) SetNillableUpdatedAt(t *time.Time) *MemberContractContentCreate {
	if t != nil {
		mccc.SetUpdatedAt(*t)
	}
	return mccc
}

// SetMemberContractID sets the "member_contract_id" field.
func (mccc *MemberContractContentCreate) SetMemberContractID(i int64) *MemberContractContentCreate {
	mccc.mutation.SetMemberContractID(i)
	return mccc
}

// SetNillableMemberContractID sets the "member_contract_id" field if the given value is not nil.
func (mccc *MemberContractContentCreate) SetNillableMemberContractID(i *int64) *MemberContractContentCreate {
	if i != nil {
		mccc.SetMemberContractID(*i)
	}
	return mccc
}

// SetContent sets the "content" field.
func (mccc *MemberContractContentCreate) SetContent(s string) *MemberContractContentCreate {
	mccc.mutation.SetContent(s)
	return mccc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mccc *MemberContractContentCreate) SetNillableContent(s *string) *MemberContractContentCreate {
	if s != nil {
		mccc.SetContent(*s)
	}
	return mccc
}

// SetSignImg sets the "sign_img" field.
func (mccc *MemberContractContentCreate) SetSignImg(s string) *MemberContractContentCreate {
	mccc.mutation.SetSignImg(s)
	return mccc
}

// SetNillableSignImg sets the "sign_img" field if the given value is not nil.
func (mccc *MemberContractContentCreate) SetNillableSignImg(s *string) *MemberContractContentCreate {
	if s != nil {
		mccc.SetSignImg(*s)
	}
	return mccc
}

// SetID sets the "id" field.
func (mccc *MemberContractContentCreate) SetID(i int64) *MemberContractContentCreate {
	mccc.mutation.SetID(i)
	return mccc
}

// SetContractID sets the "contract" edge to the MemberContract entity by ID.
func (mccc *MemberContractContentCreate) SetContractID(id int64) *MemberContractContentCreate {
	mccc.mutation.SetContractID(id)
	return mccc
}

// SetNillableContractID sets the "contract" edge to the MemberContract entity by ID if the given value is not nil.
func (mccc *MemberContractContentCreate) SetNillableContractID(id *int64) *MemberContractContentCreate {
	if id != nil {
		mccc = mccc.SetContractID(*id)
	}
	return mccc
}

// SetContract sets the "contract" edge to the MemberContract entity.
func (mccc *MemberContractContentCreate) SetContract(m *MemberContract) *MemberContractContentCreate {
	return mccc.SetContractID(m.ID)
}

// Mutation returns the MemberContractContentMutation object of the builder.
func (mccc *MemberContractContentCreate) Mutation() *MemberContractContentMutation {
	return mccc.mutation
}

// Save creates the MemberContractContent in the database.
func (mccc *MemberContractContentCreate) Save(ctx context.Context) (*MemberContractContent, error) {
	mccc.defaults()
	return withHooks(ctx, mccc.sqlSave, mccc.mutation, mccc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mccc *MemberContractContentCreate) SaveX(ctx context.Context) *MemberContractContent {
	v, err := mccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mccc *MemberContractContentCreate) Exec(ctx context.Context) error {
	_, err := mccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mccc *MemberContractContentCreate) ExecX(ctx context.Context) {
	if err := mccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mccc *MemberContractContentCreate) defaults() {
	if _, ok := mccc.mutation.CreatedAt(); !ok {
		v := membercontractcontent.DefaultCreatedAt()
		mccc.mutation.SetCreatedAt(v)
	}
	if _, ok := mccc.mutation.UpdatedAt(); !ok {
		v := membercontractcontent.DefaultUpdatedAt()
		mccc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mccc *MemberContractContentCreate) check() error {
	if _, ok := mccc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MemberContractContent.created_at"`)}
	}
	if _, ok := mccc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MemberContractContent.updated_at"`)}
	}
	return nil
}

func (mccc *MemberContractContentCreate) sqlSave(ctx context.Context) (*MemberContractContent, error) {
	if err := mccc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mccc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mccc.mutation.id = &_node.ID
	mccc.mutation.done = true
	return _node, nil
}

func (mccc *MemberContractContentCreate) createSpec() (*MemberContractContent, *sqlgraph.CreateSpec) {
	var (
		_node = &MemberContractContent{config: mccc.config}
		_spec = sqlgraph.NewCreateSpec(membercontractcontent.Table, sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64))
	)
	if id, ok := mccc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mccc.mutation.CreatedAt(); ok {
		_spec.SetField(membercontractcontent.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mccc.mutation.UpdatedAt(); ok {
		_spec.SetField(membercontractcontent.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mccc.mutation.Content(); ok {
		_spec.SetField(membercontractcontent.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := mccc.mutation.SignImg(); ok {
		_spec.SetField(membercontractcontent.FieldSignImg, field.TypeString, value)
		_node.SignImg = value
	}
	if nodes := mccc.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membercontractcontent.ContractTable,
			Columns: []string{membercontractcontent.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontract.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MemberContractID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberContractContentCreateBulk is the builder for creating many MemberContractContent entities in bulk.
type MemberContractContentCreateBulk struct {
	config
	err      error
	builders []*MemberContractContentCreate
}

// Save creates the MemberContractContent entities in the database.
func (mcccb *MemberContractContentCreateBulk) Save(ctx context.Context) ([]*MemberContractContent, error) {
	if mcccb.err != nil {
		return nil, mcccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcccb.builders))
	nodes := make([]*MemberContractContent, len(mcccb.builders))
	mutators := make([]Mutator, len(mcccb.builders))
	for i := range mcccb.builders {
		func(i int, root context.Context) {
			builder := mcccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberContractContentMutation)
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
					_, err = mutators[i+1].Mutate(root, mcccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcccb *MemberContractContentCreateBulk) SaveX(ctx context.Context) []*MemberContractContent {
	v, err := mcccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcccb *MemberContractContentCreateBulk) Exec(ctx context.Context) error {
	_, err := mcccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcccb *MemberContractContentCreateBulk) ExecX(ctx context.Context) {
	if err := mcccb.Exec(ctx); err != nil {
		panic(err)
	}
}
