// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"kcers/biz/dal/db/mysql/ent/sms"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SmsUpdate is the builder for updating Sms entities.
type SmsUpdate struct {
	config
	hooks     []Hook
	mutation  *SmsMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SmsUpdate builder.
func (su *SmsUpdate) Where(ps ...predicate.Sms) *SmsUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SmsUpdate) SetUpdatedAt(t time.Time) *SmsUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (su *SmsUpdate) ClearUpdatedAt() *SmsUpdate {
	su.mutation.ClearUpdatedAt()
	return su
}

// SetDelete sets the "delete" field.
func (su *SmsUpdate) SetDelete(i int64) *SmsUpdate {
	su.mutation.ResetDelete()
	su.mutation.SetDelete(i)
	return su
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (su *SmsUpdate) SetNillableDelete(i *int64) *SmsUpdate {
	if i != nil {
		su.SetDelete(*i)
	}
	return su
}

// AddDelete adds i to the "delete" field.
func (su *SmsUpdate) AddDelete(i int64) *SmsUpdate {
	su.mutation.AddDelete(i)
	return su
}

// ClearDelete clears the value of the "delete" field.
func (su *SmsUpdate) ClearDelete() *SmsUpdate {
	su.mutation.ClearDelete()
	return su
}

// SetCreatedID sets the "created_id" field.
func (su *SmsUpdate) SetCreatedID(i int64) *SmsUpdate {
	su.mutation.ResetCreatedID()
	su.mutation.SetCreatedID(i)
	return su
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (su *SmsUpdate) SetNillableCreatedID(i *int64) *SmsUpdate {
	if i != nil {
		su.SetCreatedID(*i)
	}
	return su
}

// AddCreatedID adds i to the "created_id" field.
func (su *SmsUpdate) AddCreatedID(i int64) *SmsUpdate {
	su.mutation.AddCreatedID(i)
	return su
}

// ClearCreatedID clears the value of the "created_id" field.
func (su *SmsUpdate) ClearCreatedID() *SmsUpdate {
	su.mutation.ClearCreatedID()
	return su
}

// SetNoticeCount sets the "notice_count" field.
func (su *SmsUpdate) SetNoticeCount(i int64) *SmsUpdate {
	su.mutation.ResetNoticeCount()
	su.mutation.SetNoticeCount(i)
	return su
}

// SetNillableNoticeCount sets the "notice_count" field if the given value is not nil.
func (su *SmsUpdate) SetNillableNoticeCount(i *int64) *SmsUpdate {
	if i != nil {
		su.SetNoticeCount(*i)
	}
	return su
}

// AddNoticeCount adds i to the "notice_count" field.
func (su *SmsUpdate) AddNoticeCount(i int64) *SmsUpdate {
	su.mutation.AddNoticeCount(i)
	return su
}

// SetUsedNotice sets the "used_notice" field.
func (su *SmsUpdate) SetUsedNotice(i int64) *SmsUpdate {
	su.mutation.ResetUsedNotice()
	su.mutation.SetUsedNotice(i)
	return su
}

// SetNillableUsedNotice sets the "used_notice" field if the given value is not nil.
func (su *SmsUpdate) SetNillableUsedNotice(i *int64) *SmsUpdate {
	if i != nil {
		su.SetUsedNotice(*i)
	}
	return su
}

// AddUsedNotice adds i to the "used_notice" field.
func (su *SmsUpdate) AddUsedNotice(i int64) *SmsUpdate {
	su.mutation.AddUsedNotice(i)
	return su
}

// Mutation returns the SmsMutation object of the builder.
func (su *SmsUpdate) Mutation() *SmsMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SmsUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SmsUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SmsUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SmsUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SmsUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok && !su.mutation.UpdatedAtCleared() {
		v := sms.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SmsUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SmsUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SmsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sms.Table, sms.Columns, sqlgraph.NewFieldSpec(sms.FieldID, field.TypeInt64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if su.mutation.CreatedAtCleared() {
		_spec.ClearField(sms.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(sms.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.UpdatedAtCleared() {
		_spec.ClearField(sms.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := su.mutation.Delete(); ok {
		_spec.SetField(sms.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedDelete(); ok {
		_spec.AddField(sms.FieldDelete, field.TypeInt64, value)
	}
	if su.mutation.DeleteCleared() {
		_spec.ClearField(sms.FieldDelete, field.TypeInt64)
	}
	if value, ok := su.mutation.CreatedID(); ok {
		_spec.SetField(sms.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedCreatedID(); ok {
		_spec.AddField(sms.FieldCreatedID, field.TypeInt64, value)
	}
	if su.mutation.CreatedIDCleared() {
		_spec.ClearField(sms.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := su.mutation.NoticeCount(); ok {
		_spec.SetField(sms.FieldNoticeCount, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedNoticeCount(); ok {
		_spec.AddField(sms.FieldNoticeCount, field.TypeInt64, value)
	}
	if value, ok := su.mutation.UsedNotice(); ok {
		_spec.SetField(sms.FieldUsedNotice, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedUsedNotice(); ok {
		_spec.AddField(sms.FieldUsedNotice, field.TypeInt64, value)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sms.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SmsUpdateOne is the builder for updating a single Sms entity.
type SmsUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SmsMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SmsUpdateOne) SetUpdatedAt(t time.Time) *SmsUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (suo *SmsUpdateOne) ClearUpdatedAt() *SmsUpdateOne {
	suo.mutation.ClearUpdatedAt()
	return suo
}

// SetDelete sets the "delete" field.
func (suo *SmsUpdateOne) SetDelete(i int64) *SmsUpdateOne {
	suo.mutation.ResetDelete()
	suo.mutation.SetDelete(i)
	return suo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (suo *SmsUpdateOne) SetNillableDelete(i *int64) *SmsUpdateOne {
	if i != nil {
		suo.SetDelete(*i)
	}
	return suo
}

// AddDelete adds i to the "delete" field.
func (suo *SmsUpdateOne) AddDelete(i int64) *SmsUpdateOne {
	suo.mutation.AddDelete(i)
	return suo
}

// ClearDelete clears the value of the "delete" field.
func (suo *SmsUpdateOne) ClearDelete() *SmsUpdateOne {
	suo.mutation.ClearDelete()
	return suo
}

// SetCreatedID sets the "created_id" field.
func (suo *SmsUpdateOne) SetCreatedID(i int64) *SmsUpdateOne {
	suo.mutation.ResetCreatedID()
	suo.mutation.SetCreatedID(i)
	return suo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (suo *SmsUpdateOne) SetNillableCreatedID(i *int64) *SmsUpdateOne {
	if i != nil {
		suo.SetCreatedID(*i)
	}
	return suo
}

// AddCreatedID adds i to the "created_id" field.
func (suo *SmsUpdateOne) AddCreatedID(i int64) *SmsUpdateOne {
	suo.mutation.AddCreatedID(i)
	return suo
}

// ClearCreatedID clears the value of the "created_id" field.
func (suo *SmsUpdateOne) ClearCreatedID() *SmsUpdateOne {
	suo.mutation.ClearCreatedID()
	return suo
}

// SetNoticeCount sets the "notice_count" field.
func (suo *SmsUpdateOne) SetNoticeCount(i int64) *SmsUpdateOne {
	suo.mutation.ResetNoticeCount()
	suo.mutation.SetNoticeCount(i)
	return suo
}

// SetNillableNoticeCount sets the "notice_count" field if the given value is not nil.
func (suo *SmsUpdateOne) SetNillableNoticeCount(i *int64) *SmsUpdateOne {
	if i != nil {
		suo.SetNoticeCount(*i)
	}
	return suo
}

// AddNoticeCount adds i to the "notice_count" field.
func (suo *SmsUpdateOne) AddNoticeCount(i int64) *SmsUpdateOne {
	suo.mutation.AddNoticeCount(i)
	return suo
}

// SetUsedNotice sets the "used_notice" field.
func (suo *SmsUpdateOne) SetUsedNotice(i int64) *SmsUpdateOne {
	suo.mutation.ResetUsedNotice()
	suo.mutation.SetUsedNotice(i)
	return suo
}

// SetNillableUsedNotice sets the "used_notice" field if the given value is not nil.
func (suo *SmsUpdateOne) SetNillableUsedNotice(i *int64) *SmsUpdateOne {
	if i != nil {
		suo.SetUsedNotice(*i)
	}
	return suo
}

// AddUsedNotice adds i to the "used_notice" field.
func (suo *SmsUpdateOne) AddUsedNotice(i int64) *SmsUpdateOne {
	suo.mutation.AddUsedNotice(i)
	return suo
}

// Mutation returns the SmsMutation object of the builder.
func (suo *SmsUpdateOne) Mutation() *SmsMutation {
	return suo.mutation
}

// Where appends a list predicates to the SmsUpdate builder.
func (suo *SmsUpdateOne) Where(ps ...predicate.Sms) *SmsUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SmsUpdateOne) Select(field string, fields ...string) *SmsUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Sms entity.
func (suo *SmsUpdateOne) Save(ctx context.Context) (*Sms, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SmsUpdateOne) SaveX(ctx context.Context) *Sms {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SmsUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SmsUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SmsUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok && !suo.mutation.UpdatedAtCleared() {
		v := sms.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SmsUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SmsUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SmsUpdateOne) sqlSave(ctx context.Context) (_node *Sms, err error) {
	_spec := sqlgraph.NewUpdateSpec(sms.Table, sms.Columns, sqlgraph.NewFieldSpec(sms.FieldID, field.TypeInt64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Sms.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sms.FieldID)
		for _, f := range fields {
			if !sms.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sms.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if suo.mutation.CreatedAtCleared() {
		_spec.ClearField(sms.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(sms.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.UpdatedAtCleared() {
		_spec.ClearField(sms.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.Delete(); ok {
		_spec.SetField(sms.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedDelete(); ok {
		_spec.AddField(sms.FieldDelete, field.TypeInt64, value)
	}
	if suo.mutation.DeleteCleared() {
		_spec.ClearField(sms.FieldDelete, field.TypeInt64)
	}
	if value, ok := suo.mutation.CreatedID(); ok {
		_spec.SetField(sms.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedCreatedID(); ok {
		_spec.AddField(sms.FieldCreatedID, field.TypeInt64, value)
	}
	if suo.mutation.CreatedIDCleared() {
		_spec.ClearField(sms.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := suo.mutation.NoticeCount(); ok {
		_spec.SetField(sms.FieldNoticeCount, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedNoticeCount(); ok {
		_spec.AddField(sms.FieldNoticeCount, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.UsedNotice(); ok {
		_spec.SetField(sms.FieldUsedNotice, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedUsedNotice(); ok {
		_spec.AddField(sms.FieldUsedNotice, field.TypeInt64, value)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Sms{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sms.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
