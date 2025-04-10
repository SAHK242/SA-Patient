// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"patient/ent/medicalhistories"
	"patient/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MedicalHistoriesDelete is the builder for deleting a MedicalHistories entity.
type MedicalHistoriesDelete struct {
	config
	hooks    []Hook
	mutation *MedicalHistoriesMutation
}

// Where appends a list predicates to the MedicalHistoriesDelete builder.
func (mhd *MedicalHistoriesDelete) Where(ps ...predicate.MedicalHistories) *MedicalHistoriesDelete {
	mhd.mutation.Where(ps...)
	return mhd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mhd *MedicalHistoriesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mhd.sqlExec, mhd.mutation, mhd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mhd *MedicalHistoriesDelete) ExecX(ctx context.Context) int {
	n, err := mhd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mhd *MedicalHistoriesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(medicalhistories.Table, sqlgraph.NewFieldSpec(medicalhistories.FieldID, field.TypeUUID))
	if ps := mhd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mhd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mhd.mutation.done = true
	return affected, err
}

// MedicalHistoriesDeleteOne is the builder for deleting a single MedicalHistories entity.
type MedicalHistoriesDeleteOne struct {
	mhd *MedicalHistoriesDelete
}

// Where appends a list predicates to the MedicalHistoriesDelete builder.
func (mhdo *MedicalHistoriesDeleteOne) Where(ps ...predicate.MedicalHistories) *MedicalHistoriesDeleteOne {
	mhdo.mhd.mutation.Where(ps...)
	return mhdo
}

// Exec executes the deletion query.
func (mhdo *MedicalHistoriesDeleteOne) Exec(ctx context.Context) error {
	n, err := mhdo.mhd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{medicalhistories.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mhdo *MedicalHistoriesDeleteOne) ExecX(ctx context.Context) {
	if err := mhdo.Exec(ctx); err != nil {
		panic(err)
	}
}
