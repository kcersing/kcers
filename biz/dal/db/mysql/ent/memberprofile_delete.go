// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kcers/biz/dal/db/mysql/ent/memberprofile"
	"kcers/biz/dal/db/mysql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberProfileDelete is the builder for deleting a MemberProfile entity.
type MemberProfileDelete struct {
	config
	hooks    []Hook
	mutation *MemberProfileMutation
}

// Where appends a list predicates to the MemberProfileDelete builder.
func (mpd *MemberProfileDelete) Where(ps ...predicate.MemberProfile) *MemberProfileDelete {
	mpd.mutation.Where(ps...)
	return mpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mpd *MemberProfileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mpd.sqlExec, mpd.mutation, mpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mpd *MemberProfileDelete) ExecX(ctx context.Context) int {
	n, err := mpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mpd *MemberProfileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(memberprofile.Table, sqlgraph.NewFieldSpec(memberprofile.FieldID, field.TypeInt64))
	if ps := mpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mpd.mutation.done = true
	return affected, err
}

// MemberProfileDeleteOne is the builder for deleting a single MemberProfile entity.
type MemberProfileDeleteOne struct {
	mpd *MemberProfileDelete
}

// Where appends a list predicates to the MemberProfileDelete builder.
func (mpdo *MemberProfileDeleteOne) Where(ps ...predicate.MemberProfile) *MemberProfileDeleteOne {
	mpdo.mpd.mutation.Where(ps...)
	return mpdo
}

// Exec executes the deletion query.
func (mpdo *MemberProfileDeleteOne) Exec(ctx context.Context) error {
	n, err := mpdo.mpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{memberprofile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mpdo *MemberProfileDeleteOne) ExecX(ctx context.Context) {
	if err := mpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
