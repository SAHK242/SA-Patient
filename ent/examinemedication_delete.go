// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"patient/ent/examinemedication"
	"patient/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamineMedicationDelete is the builder for deleting a ExamineMedication entity.
type ExamineMedicationDelete struct {
	config
	hooks    []Hook
	mutation *ExamineMedicationMutation
}

// Where appends a list predicates to the ExamineMedicationDelete builder.
func (emd *ExamineMedicationDelete) Where(ps ...predicate.ExamineMedication) *ExamineMedicationDelete {
	emd.mutation.Where(ps...)
	return emd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (emd *ExamineMedicationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, emd.sqlExec, emd.mutation, emd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (emd *ExamineMedicationDelete) ExecX(ctx context.Context) int {
	n, err := emd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (emd *ExamineMedicationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(examinemedication.Table, sqlgraph.NewFieldSpec(examinemedication.FieldID, field.TypeInt))
	if ps := emd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, emd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	emd.mutation.done = true
	return affected, err
}

// ExamineMedicationDeleteOne is the builder for deleting a single ExamineMedication entity.
type ExamineMedicationDeleteOne struct {
	emd *ExamineMedicationDelete
}

// Where appends a list predicates to the ExamineMedicationDelete builder.
func (emdo *ExamineMedicationDeleteOne) Where(ps ...predicate.ExamineMedication) *ExamineMedicationDeleteOne {
	emdo.emd.mutation.Where(ps...)
	return emdo
}

// Exec executes the deletion query.
func (emdo *ExamineMedicationDeleteOne) Exec(ctx context.Context) error {
	n, err := emdo.emd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{examinemedication.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (emdo *ExamineMedicationDeleteOne) ExecX(ctx context.Context) {
	if err := emdo.Exec(ctx); err != nil {
		panic(err)
	}
}
