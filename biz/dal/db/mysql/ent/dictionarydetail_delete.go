// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kcers/biz/dal/db/mysql/ent/dictionarydetail"
	"kcers/biz/dal/db/mysql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DictionaryDetailDelete is the builder for deleting a DictionaryDetail entity.
type DictionaryDetailDelete struct {
	config
	hooks    []Hook
	mutation *DictionaryDetailMutation
}

// Where appends a list predicates to the DictionaryDetailDelete builder.
func (ddd *DictionaryDetailDelete) Where(ps ...predicate.DictionaryDetail) *DictionaryDetailDelete {
	ddd.mutation.Where(ps...)
	return ddd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ddd *DictionaryDetailDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ddd.sqlExec, ddd.mutation, ddd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ddd *DictionaryDetailDelete) ExecX(ctx context.Context) int {
	n, err := ddd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ddd *DictionaryDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(dictionarydetail.Table, sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeInt64))
	if ps := ddd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ddd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ddd.mutation.done = true
	return affected, err
}

// DictionaryDetailDeleteOne is the builder for deleting a single DictionaryDetail entity.
type DictionaryDetailDeleteOne struct {
	ddd *DictionaryDetailDelete
}

// Where appends a list predicates to the DictionaryDetailDelete builder.
func (dddo *DictionaryDetailDeleteOne) Where(ps ...predicate.DictionaryDetail) *DictionaryDetailDeleteOne {
	dddo.ddd.mutation.Where(ps...)
	return dddo
}

// Exec executes the deletion query.
func (dddo *DictionaryDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := dddo.ddd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dictionarydetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dddo *DictionaryDetailDeleteOne) ExecX(ctx context.Context) {
	if err := dddo.Exec(ctx); err != nil {
		panic(err)
	}
}
