// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/membercontract"
	"kcers/biz/dal/db/mysql/ent/membercontractcontent"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberContractContentUpdate is the builder for updating MemberContractContent entities.
type MemberContractContentUpdate struct {
	config
	hooks    []Hook
	mutation *MemberContractContentMutation
}

// Where appends a list predicates to the MemberContractContentUpdate builder.
func (mccu *MemberContractContentUpdate) Where(ps ...predicate.MemberContractContent) *MemberContractContentUpdate {
	mccu.mutation.Where(ps...)
	return mccu
}

// SetUpdatedAt sets the "updated_at" field.
func (mccu *MemberContractContentUpdate) SetUpdatedAt(t time.Time) *MemberContractContentUpdate {
	mccu.mutation.SetUpdatedAt(t)
	return mccu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mccu *MemberContractContentUpdate) ClearUpdatedAt() *MemberContractContentUpdate {
	mccu.mutation.ClearUpdatedAt()
	return mccu
}

// SetDelete sets the "delete" field.
func (mccu *MemberContractContentUpdate) SetDelete(i int64) *MemberContractContentUpdate {
	mccu.mutation.ResetDelete()
	mccu.mutation.SetDelete(i)
	return mccu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mccu *MemberContractContentUpdate) SetNillableDelete(i *int64) *MemberContractContentUpdate {
	if i != nil {
		mccu.SetDelete(*i)
	}
	return mccu
}

// AddDelete adds i to the "delete" field.
func (mccu *MemberContractContentUpdate) AddDelete(i int64) *MemberContractContentUpdate {
	mccu.mutation.AddDelete(i)
	return mccu
}

// ClearDelete clears the value of the "delete" field.
func (mccu *MemberContractContentUpdate) ClearDelete() *MemberContractContentUpdate {
	mccu.mutation.ClearDelete()
	return mccu
}

// SetCreatedID sets the "created_id" field.
func (mccu *MemberContractContentUpdate) SetCreatedID(i int64) *MemberContractContentUpdate {
	mccu.mutation.ResetCreatedID()
	mccu.mutation.SetCreatedID(i)
	return mccu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mccu *MemberContractContentUpdate) SetNillableCreatedID(i *int64) *MemberContractContentUpdate {
	if i != nil {
		mccu.SetCreatedID(*i)
	}
	return mccu
}

// AddCreatedID adds i to the "created_id" field.
func (mccu *MemberContractContentUpdate) AddCreatedID(i int64) *MemberContractContentUpdate {
	mccu.mutation.AddCreatedID(i)
	return mccu
}

// ClearCreatedID clears the value of the "created_id" field.
func (mccu *MemberContractContentUpdate) ClearCreatedID() *MemberContractContentUpdate {
	mccu.mutation.ClearCreatedID()
	return mccu
}

// SetMemberContractID sets the "member_contract_id" field.
func (mccu *MemberContractContentUpdate) SetMemberContractID(i int64) *MemberContractContentUpdate {
	mccu.mutation.SetMemberContractID(i)
	return mccu
}

// SetNillableMemberContractID sets the "member_contract_id" field if the given value is not nil.
func (mccu *MemberContractContentUpdate) SetNillableMemberContractID(i *int64) *MemberContractContentUpdate {
	if i != nil {
		mccu.SetMemberContractID(*i)
	}
	return mccu
}

// ClearMemberContractID clears the value of the "member_contract_id" field.
func (mccu *MemberContractContentUpdate) ClearMemberContractID() *MemberContractContentUpdate {
	mccu.mutation.ClearMemberContractID()
	return mccu
}

// SetContent sets the "content" field.
func (mccu *MemberContractContentUpdate) SetContent(s string) *MemberContractContentUpdate {
	mccu.mutation.SetContent(s)
	return mccu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mccu *MemberContractContentUpdate) SetNillableContent(s *string) *MemberContractContentUpdate {
	if s != nil {
		mccu.SetContent(*s)
	}
	return mccu
}

// ClearContent clears the value of the "content" field.
func (mccu *MemberContractContentUpdate) ClearContent() *MemberContractContentUpdate {
	mccu.mutation.ClearContent()
	return mccu
}

// SetSignImg sets the "sign_img" field.
func (mccu *MemberContractContentUpdate) SetSignImg(s string) *MemberContractContentUpdate {
	mccu.mutation.SetSignImg(s)
	return mccu
}

// SetNillableSignImg sets the "sign_img" field if the given value is not nil.
func (mccu *MemberContractContentUpdate) SetNillableSignImg(s *string) *MemberContractContentUpdate {
	if s != nil {
		mccu.SetSignImg(*s)
	}
	return mccu
}

// ClearSignImg clears the value of the "sign_img" field.
func (mccu *MemberContractContentUpdate) ClearSignImg() *MemberContractContentUpdate {
	mccu.mutation.ClearSignImg()
	return mccu
}

// SetContractID sets the "contract" edge to the MemberContract entity by ID.
func (mccu *MemberContractContentUpdate) SetContractID(id int64) *MemberContractContentUpdate {
	mccu.mutation.SetContractID(id)
	return mccu
}

// SetNillableContractID sets the "contract" edge to the MemberContract entity by ID if the given value is not nil.
func (mccu *MemberContractContentUpdate) SetNillableContractID(id *int64) *MemberContractContentUpdate {
	if id != nil {
		mccu = mccu.SetContractID(*id)
	}
	return mccu
}

// SetContract sets the "contract" edge to the MemberContract entity.
func (mccu *MemberContractContentUpdate) SetContract(m *MemberContract) *MemberContractContentUpdate {
	return mccu.SetContractID(m.ID)
}

// Mutation returns the MemberContractContentMutation object of the builder.
func (mccu *MemberContractContentUpdate) Mutation() *MemberContractContentMutation {
	return mccu.mutation
}

// ClearContract clears the "contract" edge to the MemberContract entity.
func (mccu *MemberContractContentUpdate) ClearContract() *MemberContractContentUpdate {
	mccu.mutation.ClearContract()
	return mccu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mccu *MemberContractContentUpdate) Save(ctx context.Context) (int, error) {
	mccu.defaults()
	return withHooks(ctx, mccu.sqlSave, mccu.mutation, mccu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mccu *MemberContractContentUpdate) SaveX(ctx context.Context) int {
	affected, err := mccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mccu *MemberContractContentUpdate) Exec(ctx context.Context) error {
	_, err := mccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mccu *MemberContractContentUpdate) ExecX(ctx context.Context) {
	if err := mccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mccu *MemberContractContentUpdate) defaults() {
	if _, ok := mccu.mutation.UpdatedAt(); !ok && !mccu.mutation.UpdatedAtCleared() {
		v := membercontractcontent.UpdateDefaultUpdatedAt()
		mccu.mutation.SetUpdatedAt(v)
	}
}

func (mccu *MemberContractContentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(membercontractcontent.Table, membercontractcontent.Columns, sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64))
	if ps := mccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mccu.mutation.CreatedAtCleared() {
		_spec.ClearField(membercontractcontent.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mccu.mutation.UpdatedAt(); ok {
		_spec.SetField(membercontractcontent.FieldUpdatedAt, field.TypeTime, value)
	}
	if mccu.mutation.UpdatedAtCleared() {
		_spec.ClearField(membercontractcontent.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := mccu.mutation.Delete(); ok {
		_spec.SetField(membercontractcontent.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := mccu.mutation.AddedDelete(); ok {
		_spec.AddField(membercontractcontent.FieldDelete, field.TypeInt64, value)
	}
	if mccu.mutation.DeleteCleared() {
		_spec.ClearField(membercontractcontent.FieldDelete, field.TypeInt64)
	}
	if value, ok := mccu.mutation.CreatedID(); ok {
		_spec.SetField(membercontractcontent.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := mccu.mutation.AddedCreatedID(); ok {
		_spec.AddField(membercontractcontent.FieldCreatedID, field.TypeInt64, value)
	}
	if mccu.mutation.CreatedIDCleared() {
		_spec.ClearField(membercontractcontent.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := mccu.mutation.Content(); ok {
		_spec.SetField(membercontractcontent.FieldContent, field.TypeString, value)
	}
	if mccu.mutation.ContentCleared() {
		_spec.ClearField(membercontractcontent.FieldContent, field.TypeString)
	}
	if value, ok := mccu.mutation.SignImg(); ok {
		_spec.SetField(membercontractcontent.FieldSignImg, field.TypeString, value)
	}
	if mccu.mutation.SignImgCleared() {
		_spec.ClearField(membercontractcontent.FieldSignImg, field.TypeString)
	}
	if mccu.mutation.ContractCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mccu.mutation.ContractIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membercontractcontent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mccu.mutation.done = true
	return n, nil
}

// MemberContractContentUpdateOne is the builder for updating a single MemberContractContent entity.
type MemberContractContentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberContractContentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (mccuo *MemberContractContentUpdateOne) SetUpdatedAt(t time.Time) *MemberContractContentUpdateOne {
	mccuo.mutation.SetUpdatedAt(t)
	return mccuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mccuo *MemberContractContentUpdateOne) ClearUpdatedAt() *MemberContractContentUpdateOne {
	mccuo.mutation.ClearUpdatedAt()
	return mccuo
}

// SetDelete sets the "delete" field.
func (mccuo *MemberContractContentUpdateOne) SetDelete(i int64) *MemberContractContentUpdateOne {
	mccuo.mutation.ResetDelete()
	mccuo.mutation.SetDelete(i)
	return mccuo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mccuo *MemberContractContentUpdateOne) SetNillableDelete(i *int64) *MemberContractContentUpdateOne {
	if i != nil {
		mccuo.SetDelete(*i)
	}
	return mccuo
}

// AddDelete adds i to the "delete" field.
func (mccuo *MemberContractContentUpdateOne) AddDelete(i int64) *MemberContractContentUpdateOne {
	mccuo.mutation.AddDelete(i)
	return mccuo
}

// ClearDelete clears the value of the "delete" field.
func (mccuo *MemberContractContentUpdateOne) ClearDelete() *MemberContractContentUpdateOne {
	mccuo.mutation.ClearDelete()
	return mccuo
}

// SetCreatedID sets the "created_id" field.
func (mccuo *MemberContractContentUpdateOne) SetCreatedID(i int64) *MemberContractContentUpdateOne {
	mccuo.mutation.ResetCreatedID()
	mccuo.mutation.SetCreatedID(i)
	return mccuo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mccuo *MemberContractContentUpdateOne) SetNillableCreatedID(i *int64) *MemberContractContentUpdateOne {
	if i != nil {
		mccuo.SetCreatedID(*i)
	}
	return mccuo
}

// AddCreatedID adds i to the "created_id" field.
func (mccuo *MemberContractContentUpdateOne) AddCreatedID(i int64) *MemberContractContentUpdateOne {
	mccuo.mutation.AddCreatedID(i)
	return mccuo
}

// ClearCreatedID clears the value of the "created_id" field.
func (mccuo *MemberContractContentUpdateOne) ClearCreatedID() *MemberContractContentUpdateOne {
	mccuo.mutation.ClearCreatedID()
	return mccuo
}

// SetMemberContractID sets the "member_contract_id" field.
func (mccuo *MemberContractContentUpdateOne) SetMemberContractID(i int64) *MemberContractContentUpdateOne {
	mccuo.mutation.SetMemberContractID(i)
	return mccuo
}

// SetNillableMemberContractID sets the "member_contract_id" field if the given value is not nil.
func (mccuo *MemberContractContentUpdateOne) SetNillableMemberContractID(i *int64) *MemberContractContentUpdateOne {
	if i != nil {
		mccuo.SetMemberContractID(*i)
	}
	return mccuo
}

// ClearMemberContractID clears the value of the "member_contract_id" field.
func (mccuo *MemberContractContentUpdateOne) ClearMemberContractID() *MemberContractContentUpdateOne {
	mccuo.mutation.ClearMemberContractID()
	return mccuo
}

// SetContent sets the "content" field.
func (mccuo *MemberContractContentUpdateOne) SetContent(s string) *MemberContractContentUpdateOne {
	mccuo.mutation.SetContent(s)
	return mccuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mccuo *MemberContractContentUpdateOne) SetNillableContent(s *string) *MemberContractContentUpdateOne {
	if s != nil {
		mccuo.SetContent(*s)
	}
	return mccuo
}

// ClearContent clears the value of the "content" field.
func (mccuo *MemberContractContentUpdateOne) ClearContent() *MemberContractContentUpdateOne {
	mccuo.mutation.ClearContent()
	return mccuo
}

// SetSignImg sets the "sign_img" field.
func (mccuo *MemberContractContentUpdateOne) SetSignImg(s string) *MemberContractContentUpdateOne {
	mccuo.mutation.SetSignImg(s)
	return mccuo
}

// SetNillableSignImg sets the "sign_img" field if the given value is not nil.
func (mccuo *MemberContractContentUpdateOne) SetNillableSignImg(s *string) *MemberContractContentUpdateOne {
	if s != nil {
		mccuo.SetSignImg(*s)
	}
	return mccuo
}

// ClearSignImg clears the value of the "sign_img" field.
func (mccuo *MemberContractContentUpdateOne) ClearSignImg() *MemberContractContentUpdateOne {
	mccuo.mutation.ClearSignImg()
	return mccuo
}

// SetContractID sets the "contract" edge to the MemberContract entity by ID.
func (mccuo *MemberContractContentUpdateOne) SetContractID(id int64) *MemberContractContentUpdateOne {
	mccuo.mutation.SetContractID(id)
	return mccuo
}

// SetNillableContractID sets the "contract" edge to the MemberContract entity by ID if the given value is not nil.
func (mccuo *MemberContractContentUpdateOne) SetNillableContractID(id *int64) *MemberContractContentUpdateOne {
	if id != nil {
		mccuo = mccuo.SetContractID(*id)
	}
	return mccuo
}

// SetContract sets the "contract" edge to the MemberContract entity.
func (mccuo *MemberContractContentUpdateOne) SetContract(m *MemberContract) *MemberContractContentUpdateOne {
	return mccuo.SetContractID(m.ID)
}

// Mutation returns the MemberContractContentMutation object of the builder.
func (mccuo *MemberContractContentUpdateOne) Mutation() *MemberContractContentMutation {
	return mccuo.mutation
}

// ClearContract clears the "contract" edge to the MemberContract entity.
func (mccuo *MemberContractContentUpdateOne) ClearContract() *MemberContractContentUpdateOne {
	mccuo.mutation.ClearContract()
	return mccuo
}

// Where appends a list predicates to the MemberContractContentUpdate builder.
func (mccuo *MemberContractContentUpdateOne) Where(ps ...predicate.MemberContractContent) *MemberContractContentUpdateOne {
	mccuo.mutation.Where(ps...)
	return mccuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mccuo *MemberContractContentUpdateOne) Select(field string, fields ...string) *MemberContractContentUpdateOne {
	mccuo.fields = append([]string{field}, fields...)
	return mccuo
}

// Save executes the query and returns the updated MemberContractContent entity.
func (mccuo *MemberContractContentUpdateOne) Save(ctx context.Context) (*MemberContractContent, error) {
	mccuo.defaults()
	return withHooks(ctx, mccuo.sqlSave, mccuo.mutation, mccuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mccuo *MemberContractContentUpdateOne) SaveX(ctx context.Context) *MemberContractContent {
	node, err := mccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mccuo *MemberContractContentUpdateOne) Exec(ctx context.Context) error {
	_, err := mccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mccuo *MemberContractContentUpdateOne) ExecX(ctx context.Context) {
	if err := mccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mccuo *MemberContractContentUpdateOne) defaults() {
	if _, ok := mccuo.mutation.UpdatedAt(); !ok && !mccuo.mutation.UpdatedAtCleared() {
		v := membercontractcontent.UpdateDefaultUpdatedAt()
		mccuo.mutation.SetUpdatedAt(v)
	}
}

func (mccuo *MemberContractContentUpdateOne) sqlSave(ctx context.Context) (_node *MemberContractContent, err error) {
	_spec := sqlgraph.NewUpdateSpec(membercontractcontent.Table, membercontractcontent.Columns, sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64))
	id, ok := mccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MemberContractContent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membercontractcontent.FieldID)
		for _, f := range fields {
			if !membercontractcontent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != membercontractcontent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mccuo.mutation.CreatedAtCleared() {
		_spec.ClearField(membercontractcontent.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mccuo.mutation.UpdatedAt(); ok {
		_spec.SetField(membercontractcontent.FieldUpdatedAt, field.TypeTime, value)
	}
	if mccuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(membercontractcontent.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := mccuo.mutation.Delete(); ok {
		_spec.SetField(membercontractcontent.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := mccuo.mutation.AddedDelete(); ok {
		_spec.AddField(membercontractcontent.FieldDelete, field.TypeInt64, value)
	}
	if mccuo.mutation.DeleteCleared() {
		_spec.ClearField(membercontractcontent.FieldDelete, field.TypeInt64)
	}
	if value, ok := mccuo.mutation.CreatedID(); ok {
		_spec.SetField(membercontractcontent.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := mccuo.mutation.AddedCreatedID(); ok {
		_spec.AddField(membercontractcontent.FieldCreatedID, field.TypeInt64, value)
	}
	if mccuo.mutation.CreatedIDCleared() {
		_spec.ClearField(membercontractcontent.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := mccuo.mutation.Content(); ok {
		_spec.SetField(membercontractcontent.FieldContent, field.TypeString, value)
	}
	if mccuo.mutation.ContentCleared() {
		_spec.ClearField(membercontractcontent.FieldContent, field.TypeString)
	}
	if value, ok := mccuo.mutation.SignImg(); ok {
		_spec.SetField(membercontractcontent.FieldSignImg, field.TypeString, value)
	}
	if mccuo.mutation.SignImgCleared() {
		_spec.ClearField(membercontractcontent.FieldSignImg, field.TypeString)
	}
	if mccuo.mutation.ContractCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mccuo.mutation.ContractIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MemberContractContent{config: mccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membercontractcontent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mccuo.mutation.done = true
	return _node, nil
}
