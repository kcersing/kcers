// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers/biz/dal/db/mysql/ent/api"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// APIUpdate is the builder for updating API entities.
type APIUpdate struct {
	config
	hooks    []Hook
	mutation *APIMutation
}

// Where appends a list predicates to the APIUpdate builder.
func (au *APIUpdate) Where(ps ...predicate.API) *APIUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *APIUpdate) SetUpdatedAt(t time.Time) *APIUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetPath sets the "path" field.
func (au *APIUpdate) SetPath(s string) *APIUpdate {
	au.mutation.SetPath(s)
	return au
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (au *APIUpdate) SetNillablePath(s *string) *APIUpdate {
	if s != nil {
		au.SetPath(*s)
	}
	return au
}

// SetTitle sets the "title" field.
func (au *APIUpdate) SetTitle(s string) *APIUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (au *APIUpdate) SetNillableTitle(s *string) *APIUpdate {
	if s != nil {
		au.SetTitle(*s)
	}
	return au
}

// SetDescription sets the "description" field.
func (au *APIUpdate) SetDescription(s string) *APIUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (au *APIUpdate) SetNillableDescription(s *string) *APIUpdate {
	if s != nil {
		au.SetDescription(*s)
	}
	return au
}

// SetAPIGroup sets the "api_group" field.
func (au *APIUpdate) SetAPIGroup(s string) *APIUpdate {
	au.mutation.SetAPIGroup(s)
	return au
}

// SetNillableAPIGroup sets the "api_group" field if the given value is not nil.
func (au *APIUpdate) SetNillableAPIGroup(s *string) *APIUpdate {
	if s != nil {
		au.SetAPIGroup(*s)
	}
	return au
}

// SetMethod sets the "method" field.
func (au *APIUpdate) SetMethod(s string) *APIUpdate {
	au.mutation.SetMethod(s)
	return au
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (au *APIUpdate) SetNillableMethod(s *string) *APIUpdate {
	if s != nil {
		au.SetMethod(*s)
	}
	return au
}

// Mutation returns the APIMutation object of the builder.
func (au *APIUpdate) Mutation() *APIMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *APIUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *APIUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *APIUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *APIUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *APIUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := api.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *APIUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(api.Table, api.Columns, sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(api.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Path(); ok {
		_spec.SetField(api.FieldPath, field.TypeString, value)
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.SetField(api.FieldTitle, field.TypeString, value)
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.SetField(api.FieldDescription, field.TypeString, value)
	}
	if value, ok := au.mutation.APIGroup(); ok {
		_spec.SetField(api.FieldAPIGroup, field.TypeString, value)
	}
	if value, ok := au.mutation.Method(); ok {
		_spec.SetField(api.FieldMethod, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{api.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// APIUpdateOne is the builder for updating a single API entity.
type APIUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *APIMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *APIUpdateOne) SetUpdatedAt(t time.Time) *APIUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetPath sets the "path" field.
func (auo *APIUpdateOne) SetPath(s string) *APIUpdateOne {
	auo.mutation.SetPath(s)
	return auo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (auo *APIUpdateOne) SetNillablePath(s *string) *APIUpdateOne {
	if s != nil {
		auo.SetPath(*s)
	}
	return auo
}

// SetTitle sets the "title" field.
func (auo *APIUpdateOne) SetTitle(s string) *APIUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (auo *APIUpdateOne) SetNillableTitle(s *string) *APIUpdateOne {
	if s != nil {
		auo.SetTitle(*s)
	}
	return auo
}

// SetDescription sets the "description" field.
func (auo *APIUpdateOne) SetDescription(s string) *APIUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (auo *APIUpdateOne) SetNillableDescription(s *string) *APIUpdateOne {
	if s != nil {
		auo.SetDescription(*s)
	}
	return auo
}

// SetAPIGroup sets the "api_group" field.
func (auo *APIUpdateOne) SetAPIGroup(s string) *APIUpdateOne {
	auo.mutation.SetAPIGroup(s)
	return auo
}

// SetNillableAPIGroup sets the "api_group" field if the given value is not nil.
func (auo *APIUpdateOne) SetNillableAPIGroup(s *string) *APIUpdateOne {
	if s != nil {
		auo.SetAPIGroup(*s)
	}
	return auo
}

// SetMethod sets the "method" field.
func (auo *APIUpdateOne) SetMethod(s string) *APIUpdateOne {
	auo.mutation.SetMethod(s)
	return auo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (auo *APIUpdateOne) SetNillableMethod(s *string) *APIUpdateOne {
	if s != nil {
		auo.SetMethod(*s)
	}
	return auo
}

// Mutation returns the APIMutation object of the builder.
func (auo *APIUpdateOne) Mutation() *APIMutation {
	return auo.mutation
}

// Where appends a list predicates to the APIUpdate builder.
func (auo *APIUpdateOne) Where(ps ...predicate.API) *APIUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *APIUpdateOne) Select(field string, fields ...string) *APIUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated API entity.
func (auo *APIUpdateOne) Save(ctx context.Context) (*API, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *APIUpdateOne) SaveX(ctx context.Context) *API {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *APIUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *APIUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *APIUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := api.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *APIUpdateOne) sqlSave(ctx context.Context) (_node *API, err error) {
	_spec := sqlgraph.NewUpdateSpec(api.Table, api.Columns, sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "API.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, api.FieldID)
		for _, f := range fields {
			if !api.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != api.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(api.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Path(); ok {
		_spec.SetField(api.FieldPath, field.TypeString, value)
	}
	if value, ok := auo.mutation.Title(); ok {
		_spec.SetField(api.FieldTitle, field.TypeString, value)
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.SetField(api.FieldDescription, field.TypeString, value)
	}
	if value, ok := auo.mutation.APIGroup(); ok {
		_spec.SetField(api.FieldAPIGroup, field.TypeString, value)
	}
	if value, ok := auo.mutation.Method(); ok {
		_spec.SetField(api.FieldMethod, field.TypeString, value)
	}
	_node = &API{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{api.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
