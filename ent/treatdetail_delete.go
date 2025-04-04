// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"patient/ent/predicate"
	"patient/ent/treatdetail"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TreatDetailDelete is the builder for deleting a TreatDetail entity.
type TreatDetailDelete struct {
	config
	hooks    []Hook
	mutation *TreatDetailMutation
}

// Where appends a list predicates to the TreatDetailDelete builder.
func (tdd *TreatDetailDelete) Where(ps ...predicate.TreatDetail) *TreatDetailDelete {
	tdd.mutation.Where(ps...)
	return tdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tdd *TreatDetailDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tdd.sqlExec, tdd.mutation, tdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tdd *TreatDetailDelete) ExecX(ctx context.Context) int {
	n, err := tdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tdd *TreatDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(treatdetail.Table, sqlgraph.NewFieldSpec(treatdetail.FieldID, field.TypeInt))
	if ps := tdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tdd.mutation.done = true
	return affected, err
}

// TreatDetailDeleteOne is the builder for deleting a single TreatDetail entity.
type TreatDetailDeleteOne struct {
	tdd *TreatDetailDelete
}

// Where appends a list predicates to the TreatDetailDelete builder.
func (tddo *TreatDetailDeleteOne) Where(ps ...predicate.TreatDetail) *TreatDetailDeleteOne {
	tddo.tdd.mutation.Where(ps...)
	return tddo
}

// Exec executes the deletion query.
func (tddo *TreatDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := tddo.tdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{treatdetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tddo *TreatDetailDeleteOne) ExecX(ctx context.Context) {
	if err := tddo.Exec(ctx); err != nil {
		panic(err)
	}
}
