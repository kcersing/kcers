// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kcers/biz/dal/db/mysql/ent/membercontract"
	"kcers/biz/dal/db/mysql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberContractDelete is the builder for deleting a MemberContract entity.
type MemberContractDelete struct {
	config
	hooks    []Hook
	mutation *MemberContractMutation
}

// Where appends a list predicates to the MemberContractDelete builder.
func (mcd *MemberContractDelete) Where(ps ...predicate.MemberContract) *MemberContractDelete {
	mcd.mutation.Where(ps...)
	return mcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mcd *MemberContractDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mcd.sqlExec, mcd.mutation, mcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mcd *MemberContractDelete) ExecX(ctx context.Context) int {
	n, err := mcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mcd *MemberContractDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(membercontract.Table, sqlgraph.NewFieldSpec(membercontract.FieldID, field.TypeInt64))
	if ps := mcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mcd.mutation.done = true
	return affected, err
}

// MemberContractDeleteOne is the builder for deleting a single MemberContract entity.
type MemberContractDeleteOne struct {
	mcd *MemberContractDelete
}

// Where appends a list predicates to the MemberContractDelete builder.
func (mcdo *MemberContractDeleteOne) Where(ps ...predicate.MemberContract) *MemberContractDeleteOne {
	mcdo.mcd.mutation.Where(ps...)
	return mcdo
}

// Exec executes the deletion query.
func (mcdo *MemberContractDeleteOne) Exec(ctx context.Context) error {
	n, err := mcdo.mcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{membercontract.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mcdo *MemberContractDeleteOne) ExecX(ctx context.Context) {
	if err := mcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
