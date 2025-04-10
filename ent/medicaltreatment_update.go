// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"patient/ent/medicalhistories"
	"patient/ent/medicaltreatment"
	"patient/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MedicalTreatmentUpdate is the builder for updating MedicalTreatment entities.
type MedicalTreatmentUpdate struct {
	config
	hooks    []Hook
	mutation *MedicalTreatmentMutation
}

// Where appends a list predicates to the MedicalTreatmentUpdate builder.
func (mtu *MedicalTreatmentUpdate) Where(ps ...predicate.MedicalTreatment) *MedicalTreatmentUpdate {
	mtu.mutation.Where(ps...)
	return mtu
}

// SetMedicalHistoryID sets the "medical_history_id" field.
func (mtu *MedicalTreatmentUpdate) SetMedicalHistoryID(u uuid.UUID) *MedicalTreatmentUpdate {
	mtu.mutation.SetMedicalHistoryID(u)
	return mtu
}

// SetNillableMedicalHistoryID sets the "medical_history_id" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableMedicalHistoryID(u *uuid.UUID) *MedicalTreatmentUpdate {
	if u != nil {
		mtu.SetMedicalHistoryID(*u)
	}
	return mtu
}

// SetStartDate sets the "start_date" field.
func (mtu *MedicalTreatmentUpdate) SetStartDate(t time.Time) *MedicalTreatmentUpdate {
	mtu.mutation.SetStartDate(t)
	return mtu
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableStartDate(t *time.Time) *MedicalTreatmentUpdate {
	if t != nil {
		mtu.SetStartDate(*t)
	}
	return mtu
}

// SetEndDate sets the "end_date" field.
func (mtu *MedicalTreatmentUpdate) SetEndDate(t time.Time) *MedicalTreatmentUpdate {
	mtu.mutation.SetEndDate(t)
	return mtu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableEndDate(t *time.Time) *MedicalTreatmentUpdate {
	if t != nil {
		mtu.SetEndDate(*t)
	}
	return mtu
}

// SetName sets the "name" field.
func (mtu *MedicalTreatmentUpdate) SetName(s string) *MedicalTreatmentUpdate {
	mtu.mutation.SetName(s)
	return mtu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableName(s *string) *MedicalTreatmentUpdate {
	if s != nil {
		mtu.SetName(*s)
	}
	return mtu
}

// SetResult sets the "result" field.
func (mtu *MedicalTreatmentUpdate) SetResult(s string) *MedicalTreatmentUpdate {
	mtu.mutation.SetResult(s)
	return mtu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableResult(s *string) *MedicalTreatmentUpdate {
	if s != nil {
		mtu.SetResult(*s)
	}
	return mtu
}

// SetDescription sets the "description" field.
func (mtu *MedicalTreatmentUpdate) SetDescription(s string) *MedicalTreatmentUpdate {
	mtu.mutation.SetDescription(s)
	return mtu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableDescription(s *string) *MedicalTreatmentUpdate {
	if s != nil {
		mtu.SetDescription(*s)
	}
	return mtu
}

// SetFee sets the "fee" field.
func (mtu *MedicalTreatmentUpdate) SetFee(f float64) *MedicalTreatmentUpdate {
	mtu.mutation.ResetFee()
	mtu.mutation.SetFee(f)
	return mtu
}

// SetNillableFee sets the "fee" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableFee(f *float64) *MedicalTreatmentUpdate {
	if f != nil {
		mtu.SetFee(*f)
	}
	return mtu
}

// AddFee adds f to the "fee" field.
func (mtu *MedicalTreatmentUpdate) AddFee(f float64) *MedicalTreatmentUpdate {
	mtu.mutation.AddFee(f)
	return mtu
}

// SetMainDoctorID sets the "main_doctor_id" field.
func (mtu *MedicalTreatmentUpdate) SetMainDoctorID(u uuid.UUID) *MedicalTreatmentUpdate {
	mtu.mutation.SetMainDoctorID(u)
	return mtu
}

// SetNillableMainDoctorID sets the "main_doctor_id" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableMainDoctorID(u *uuid.UUID) *MedicalTreatmentUpdate {
	if u != nil {
		mtu.SetMainDoctorID(*u)
	}
	return mtu
}

// SetSupportDoctorIds sets the "support_doctor_ids" field.
func (mtu *MedicalTreatmentUpdate) SetSupportDoctorIds(s string) *MedicalTreatmentUpdate {
	mtu.mutation.SetSupportDoctorIds(s)
	return mtu
}

// SetNillableSupportDoctorIds sets the "support_doctor_ids" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableSupportDoctorIds(s *string) *MedicalTreatmentUpdate {
	if s != nil {
		mtu.SetSupportDoctorIds(*s)
	}
	return mtu
}

// SetSupportNurseIds sets the "support_nurse_ids" field.
func (mtu *MedicalTreatmentUpdate) SetSupportNurseIds(s string) *MedicalTreatmentUpdate {
	mtu.mutation.SetSupportNurseIds(s)
	return mtu
}

// SetNillableSupportNurseIds sets the "support_nurse_ids" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableSupportNurseIds(s *string) *MedicalTreatmentUpdate {
	if s != nil {
		mtu.SetSupportNurseIds(*s)
	}
	return mtu
}

// SetCreatedAt sets the "created_at" field.
func (mtu *MedicalTreatmentUpdate) SetCreatedAt(t time.Time) *MedicalTreatmentUpdate {
	mtu.mutation.SetCreatedAt(t)
	return mtu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableCreatedAt(t *time.Time) *MedicalTreatmentUpdate {
	if t != nil {
		mtu.SetCreatedAt(*t)
	}
	return mtu
}

// SetCreatedBy sets the "created_by" field.
func (mtu *MedicalTreatmentUpdate) SetCreatedBy(u uuid.UUID) *MedicalTreatmentUpdate {
	mtu.mutation.SetCreatedBy(u)
	return mtu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableCreatedBy(u *uuid.UUID) *MedicalTreatmentUpdate {
	if u != nil {
		mtu.SetCreatedBy(*u)
	}
	return mtu
}

// SetUpdatedBy sets the "updated_by" field.
func (mtu *MedicalTreatmentUpdate) SetUpdatedBy(u uuid.UUID) *MedicalTreatmentUpdate {
	mtu.mutation.SetUpdatedBy(u)
	return mtu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mtu *MedicalTreatmentUpdate) SetNillableUpdatedBy(u *uuid.UUID) *MedicalTreatmentUpdate {
	if u != nil {
		mtu.SetUpdatedBy(*u)
	}
	return mtu
}

// SetUpdatedAt sets the "updated_at" field.
func (mtu *MedicalTreatmentUpdate) SetUpdatedAt(t time.Time) *MedicalTreatmentUpdate {
	mtu.mutation.SetUpdatedAt(t)
	return mtu
}

// SetMedicalHistoriesID sets the "medical_histories" edge to the MedicalHistories entity by ID.
func (mtu *MedicalTreatmentUpdate) SetMedicalHistoriesID(id uuid.UUID) *MedicalTreatmentUpdate {
	mtu.mutation.SetMedicalHistoriesID(id)
	return mtu
}

// SetMedicalHistories sets the "medical_histories" edge to the MedicalHistories entity.
func (mtu *MedicalTreatmentUpdate) SetMedicalHistories(m *MedicalHistories) *MedicalTreatmentUpdate {
	return mtu.SetMedicalHistoriesID(m.ID)
}

// Mutation returns the MedicalTreatmentMutation object of the builder.
func (mtu *MedicalTreatmentUpdate) Mutation() *MedicalTreatmentMutation {
	return mtu.mutation
}

// ClearMedicalHistories clears the "medical_histories" edge to the MedicalHistories entity.
func (mtu *MedicalTreatmentUpdate) ClearMedicalHistories() *MedicalTreatmentUpdate {
	mtu.mutation.ClearMedicalHistories()
	return mtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mtu *MedicalTreatmentUpdate) Save(ctx context.Context) (int, error) {
	mtu.defaults()
	return withHooks(ctx, mtu.sqlSave, mtu.mutation, mtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mtu *MedicalTreatmentUpdate) SaveX(ctx context.Context) int {
	affected, err := mtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mtu *MedicalTreatmentUpdate) Exec(ctx context.Context) error {
	_, err := mtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtu *MedicalTreatmentUpdate) ExecX(ctx context.Context) {
	if err := mtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mtu *MedicalTreatmentUpdate) defaults() {
	if _, ok := mtu.mutation.UpdatedAt(); !ok {
		v := medicaltreatment.UpdateDefaultUpdatedAt()
		mtu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mtu *MedicalTreatmentUpdate) check() error {
	if v, ok := mtu.mutation.Name(); ok {
		if err := medicaltreatment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MedicalTreatment.name": %w`, err)}
		}
	}
	if mtu.mutation.MedicalHistoriesCleared() && len(mtu.mutation.MedicalHistoriesIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MedicalTreatment.medical_histories"`)
	}
	return nil
}

func (mtu *MedicalTreatmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(medicaltreatment.Table, medicaltreatment.Columns, sqlgraph.NewFieldSpec(medicaltreatment.FieldID, field.TypeUUID))
	if ps := mtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mtu.mutation.StartDate(); ok {
		_spec.SetField(medicaltreatment.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := mtu.mutation.EndDate(); ok {
		_spec.SetField(medicaltreatment.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := mtu.mutation.Name(); ok {
		_spec.SetField(medicaltreatment.FieldName, field.TypeString, value)
	}
	if value, ok := mtu.mutation.Result(); ok {
		_spec.SetField(medicaltreatment.FieldResult, field.TypeString, value)
	}
	if value, ok := mtu.mutation.Description(); ok {
		_spec.SetField(medicaltreatment.FieldDescription, field.TypeString, value)
	}
	if value, ok := mtu.mutation.Fee(); ok {
		_spec.SetField(medicaltreatment.FieldFee, field.TypeFloat64, value)
	}
	if value, ok := mtu.mutation.AddedFee(); ok {
		_spec.AddField(medicaltreatment.FieldFee, field.TypeFloat64, value)
	}
	if value, ok := mtu.mutation.MainDoctorID(); ok {
		_spec.SetField(medicaltreatment.FieldMainDoctorID, field.TypeUUID, value)
	}
	if value, ok := mtu.mutation.SupportDoctorIds(); ok {
		_spec.SetField(medicaltreatment.FieldSupportDoctorIds, field.TypeString, value)
	}
	if value, ok := mtu.mutation.SupportNurseIds(); ok {
		_spec.SetField(medicaltreatment.FieldSupportNurseIds, field.TypeString, value)
	}
	if value, ok := mtu.mutation.CreatedAt(); ok {
		_spec.SetField(medicaltreatment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mtu.mutation.CreatedBy(); ok {
		_spec.SetField(medicaltreatment.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := mtu.mutation.UpdatedBy(); ok {
		_spec.SetField(medicaltreatment.FieldUpdatedBy, field.TypeUUID, value)
	}
	if value, ok := mtu.mutation.UpdatedAt(); ok {
		_spec.SetField(medicaltreatment.FieldUpdatedAt, field.TypeTime, value)
	}
	if mtu.mutation.MedicalHistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicaltreatment.MedicalHistoriesTable,
			Columns: []string{medicaltreatment.MedicalHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medicalhistories.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mtu.mutation.MedicalHistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicaltreatment.MedicalHistoriesTable,
			Columns: []string{medicaltreatment.MedicalHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medicalhistories.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicaltreatment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mtu.mutation.done = true
	return n, nil
}

// MedicalTreatmentUpdateOne is the builder for updating a single MedicalTreatment entity.
type MedicalTreatmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MedicalTreatmentMutation
}

// SetMedicalHistoryID sets the "medical_history_id" field.
func (mtuo *MedicalTreatmentUpdateOne) SetMedicalHistoryID(u uuid.UUID) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetMedicalHistoryID(u)
	return mtuo
}

// SetNillableMedicalHistoryID sets the "medical_history_id" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableMedicalHistoryID(u *uuid.UUID) *MedicalTreatmentUpdateOne {
	if u != nil {
		mtuo.SetMedicalHistoryID(*u)
	}
	return mtuo
}

// SetStartDate sets the "start_date" field.
func (mtuo *MedicalTreatmentUpdateOne) SetStartDate(t time.Time) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetStartDate(t)
	return mtuo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableStartDate(t *time.Time) *MedicalTreatmentUpdateOne {
	if t != nil {
		mtuo.SetStartDate(*t)
	}
	return mtuo
}

// SetEndDate sets the "end_date" field.
func (mtuo *MedicalTreatmentUpdateOne) SetEndDate(t time.Time) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetEndDate(t)
	return mtuo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableEndDate(t *time.Time) *MedicalTreatmentUpdateOne {
	if t != nil {
		mtuo.SetEndDate(*t)
	}
	return mtuo
}

// SetName sets the "name" field.
func (mtuo *MedicalTreatmentUpdateOne) SetName(s string) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetName(s)
	return mtuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableName(s *string) *MedicalTreatmentUpdateOne {
	if s != nil {
		mtuo.SetName(*s)
	}
	return mtuo
}

// SetResult sets the "result" field.
func (mtuo *MedicalTreatmentUpdateOne) SetResult(s string) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetResult(s)
	return mtuo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableResult(s *string) *MedicalTreatmentUpdateOne {
	if s != nil {
		mtuo.SetResult(*s)
	}
	return mtuo
}

// SetDescription sets the "description" field.
func (mtuo *MedicalTreatmentUpdateOne) SetDescription(s string) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetDescription(s)
	return mtuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableDescription(s *string) *MedicalTreatmentUpdateOne {
	if s != nil {
		mtuo.SetDescription(*s)
	}
	return mtuo
}

// SetFee sets the "fee" field.
func (mtuo *MedicalTreatmentUpdateOne) SetFee(f float64) *MedicalTreatmentUpdateOne {
	mtuo.mutation.ResetFee()
	mtuo.mutation.SetFee(f)
	return mtuo
}

// SetNillableFee sets the "fee" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableFee(f *float64) *MedicalTreatmentUpdateOne {
	if f != nil {
		mtuo.SetFee(*f)
	}
	return mtuo
}

// AddFee adds f to the "fee" field.
func (mtuo *MedicalTreatmentUpdateOne) AddFee(f float64) *MedicalTreatmentUpdateOne {
	mtuo.mutation.AddFee(f)
	return mtuo
}

// SetMainDoctorID sets the "main_doctor_id" field.
func (mtuo *MedicalTreatmentUpdateOne) SetMainDoctorID(u uuid.UUID) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetMainDoctorID(u)
	return mtuo
}

// SetNillableMainDoctorID sets the "main_doctor_id" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableMainDoctorID(u *uuid.UUID) *MedicalTreatmentUpdateOne {
	if u != nil {
		mtuo.SetMainDoctorID(*u)
	}
	return mtuo
}

// SetSupportDoctorIds sets the "support_doctor_ids" field.
func (mtuo *MedicalTreatmentUpdateOne) SetSupportDoctorIds(s string) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetSupportDoctorIds(s)
	return mtuo
}

// SetNillableSupportDoctorIds sets the "support_doctor_ids" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableSupportDoctorIds(s *string) *MedicalTreatmentUpdateOne {
	if s != nil {
		mtuo.SetSupportDoctorIds(*s)
	}
	return mtuo
}

// SetSupportNurseIds sets the "support_nurse_ids" field.
func (mtuo *MedicalTreatmentUpdateOne) SetSupportNurseIds(s string) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetSupportNurseIds(s)
	return mtuo
}

// SetNillableSupportNurseIds sets the "support_nurse_ids" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableSupportNurseIds(s *string) *MedicalTreatmentUpdateOne {
	if s != nil {
		mtuo.SetSupportNurseIds(*s)
	}
	return mtuo
}

// SetCreatedAt sets the "created_at" field.
func (mtuo *MedicalTreatmentUpdateOne) SetCreatedAt(t time.Time) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetCreatedAt(t)
	return mtuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableCreatedAt(t *time.Time) *MedicalTreatmentUpdateOne {
	if t != nil {
		mtuo.SetCreatedAt(*t)
	}
	return mtuo
}

// SetCreatedBy sets the "created_by" field.
func (mtuo *MedicalTreatmentUpdateOne) SetCreatedBy(u uuid.UUID) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetCreatedBy(u)
	return mtuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *MedicalTreatmentUpdateOne {
	if u != nil {
		mtuo.SetCreatedBy(*u)
	}
	return mtuo
}

// SetUpdatedBy sets the "updated_by" field.
func (mtuo *MedicalTreatmentUpdateOne) SetUpdatedBy(u uuid.UUID) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetUpdatedBy(u)
	return mtuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mtuo *MedicalTreatmentUpdateOne) SetNillableUpdatedBy(u *uuid.UUID) *MedicalTreatmentUpdateOne {
	if u != nil {
		mtuo.SetUpdatedBy(*u)
	}
	return mtuo
}

// SetUpdatedAt sets the "updated_at" field.
func (mtuo *MedicalTreatmentUpdateOne) SetUpdatedAt(t time.Time) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetUpdatedAt(t)
	return mtuo
}

// SetMedicalHistoriesID sets the "medical_histories" edge to the MedicalHistories entity by ID.
func (mtuo *MedicalTreatmentUpdateOne) SetMedicalHistoriesID(id uuid.UUID) *MedicalTreatmentUpdateOne {
	mtuo.mutation.SetMedicalHistoriesID(id)
	return mtuo
}

// SetMedicalHistories sets the "medical_histories" edge to the MedicalHistories entity.
func (mtuo *MedicalTreatmentUpdateOne) SetMedicalHistories(m *MedicalHistories) *MedicalTreatmentUpdateOne {
	return mtuo.SetMedicalHistoriesID(m.ID)
}

// Mutation returns the MedicalTreatmentMutation object of the builder.
func (mtuo *MedicalTreatmentUpdateOne) Mutation() *MedicalTreatmentMutation {
	return mtuo.mutation
}

// ClearMedicalHistories clears the "medical_histories" edge to the MedicalHistories entity.
func (mtuo *MedicalTreatmentUpdateOne) ClearMedicalHistories() *MedicalTreatmentUpdateOne {
	mtuo.mutation.ClearMedicalHistories()
	return mtuo
}

// Where appends a list predicates to the MedicalTreatmentUpdate builder.
func (mtuo *MedicalTreatmentUpdateOne) Where(ps ...predicate.MedicalTreatment) *MedicalTreatmentUpdateOne {
	mtuo.mutation.Where(ps...)
	return mtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mtuo *MedicalTreatmentUpdateOne) Select(field string, fields ...string) *MedicalTreatmentUpdateOne {
	mtuo.fields = append([]string{field}, fields...)
	return mtuo
}

// Save executes the query and returns the updated MedicalTreatment entity.
func (mtuo *MedicalTreatmentUpdateOne) Save(ctx context.Context) (*MedicalTreatment, error) {
	mtuo.defaults()
	return withHooks(ctx, mtuo.sqlSave, mtuo.mutation, mtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mtuo *MedicalTreatmentUpdateOne) SaveX(ctx context.Context) *MedicalTreatment {
	node, err := mtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mtuo *MedicalTreatmentUpdateOne) Exec(ctx context.Context) error {
	_, err := mtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtuo *MedicalTreatmentUpdateOne) ExecX(ctx context.Context) {
	if err := mtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mtuo *MedicalTreatmentUpdateOne) defaults() {
	if _, ok := mtuo.mutation.UpdatedAt(); !ok {
		v := medicaltreatment.UpdateDefaultUpdatedAt()
		mtuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mtuo *MedicalTreatmentUpdateOne) check() error {
	if v, ok := mtuo.mutation.Name(); ok {
		if err := medicaltreatment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MedicalTreatment.name": %w`, err)}
		}
	}
	if mtuo.mutation.MedicalHistoriesCleared() && len(mtuo.mutation.MedicalHistoriesIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MedicalTreatment.medical_histories"`)
	}
	return nil
}

func (mtuo *MedicalTreatmentUpdateOne) sqlSave(ctx context.Context) (_node *MedicalTreatment, err error) {
	if err := mtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(medicaltreatment.Table, medicaltreatment.Columns, sqlgraph.NewFieldSpec(medicaltreatment.FieldID, field.TypeUUID))
	id, ok := mtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MedicalTreatment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, medicaltreatment.FieldID)
		for _, f := range fields {
			if !medicaltreatment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != medicaltreatment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mtuo.mutation.StartDate(); ok {
		_spec.SetField(medicaltreatment.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := mtuo.mutation.EndDate(); ok {
		_spec.SetField(medicaltreatment.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := mtuo.mutation.Name(); ok {
		_spec.SetField(medicaltreatment.FieldName, field.TypeString, value)
	}
	if value, ok := mtuo.mutation.Result(); ok {
		_spec.SetField(medicaltreatment.FieldResult, field.TypeString, value)
	}
	if value, ok := mtuo.mutation.Description(); ok {
		_spec.SetField(medicaltreatment.FieldDescription, field.TypeString, value)
	}
	if value, ok := mtuo.mutation.Fee(); ok {
		_spec.SetField(medicaltreatment.FieldFee, field.TypeFloat64, value)
	}
	if value, ok := mtuo.mutation.AddedFee(); ok {
		_spec.AddField(medicaltreatment.FieldFee, field.TypeFloat64, value)
	}
	if value, ok := mtuo.mutation.MainDoctorID(); ok {
		_spec.SetField(medicaltreatment.FieldMainDoctorID, field.TypeUUID, value)
	}
	if value, ok := mtuo.mutation.SupportDoctorIds(); ok {
		_spec.SetField(medicaltreatment.FieldSupportDoctorIds, field.TypeString, value)
	}
	if value, ok := mtuo.mutation.SupportNurseIds(); ok {
		_spec.SetField(medicaltreatment.FieldSupportNurseIds, field.TypeString, value)
	}
	if value, ok := mtuo.mutation.CreatedAt(); ok {
		_spec.SetField(medicaltreatment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mtuo.mutation.CreatedBy(); ok {
		_spec.SetField(medicaltreatment.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := mtuo.mutation.UpdatedBy(); ok {
		_spec.SetField(medicaltreatment.FieldUpdatedBy, field.TypeUUID, value)
	}
	if value, ok := mtuo.mutation.UpdatedAt(); ok {
		_spec.SetField(medicaltreatment.FieldUpdatedAt, field.TypeTime, value)
	}
	if mtuo.mutation.MedicalHistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicaltreatment.MedicalHistoriesTable,
			Columns: []string{medicaltreatment.MedicalHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medicalhistories.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mtuo.mutation.MedicalHistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicaltreatment.MedicalHistoriesTable,
			Columns: []string{medicaltreatment.MedicalHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medicalhistories.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MedicalTreatment{config: mtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicaltreatment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mtuo.mutation.done = true
	return _node, nil
}
