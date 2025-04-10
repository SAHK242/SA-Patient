// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"patient/ent/medicalprescription"
	"patient/ent/medication"
	"patient/ent/predicate"
	"patient/ent/prescriptionmedication"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PrescriptionMedicationUpdate is the builder for updating PrescriptionMedication entities.
type PrescriptionMedicationUpdate struct {
	config
	hooks    []Hook
	mutation *PrescriptionMedicationMutation
}

// Where appends a list predicates to the PrescriptionMedicationUpdate builder.
func (pmu *PrescriptionMedicationUpdate) Where(ps ...predicate.PrescriptionMedication) *PrescriptionMedicationUpdate {
	pmu.mutation.Where(ps...)
	return pmu
}

// SetPrescriptionID sets the "prescription_id" field.
func (pmu *PrescriptionMedicationUpdate) SetPrescriptionID(u uuid.UUID) *PrescriptionMedicationUpdate {
	pmu.mutation.SetPrescriptionID(u)
	return pmu
}

// SetNillablePrescriptionID sets the "prescription_id" field if the given value is not nil.
func (pmu *PrescriptionMedicationUpdate) SetNillablePrescriptionID(u *uuid.UUID) *PrescriptionMedicationUpdate {
	if u != nil {
		pmu.SetPrescriptionID(*u)
	}
	return pmu
}

// SetMedicationID sets the "medication_id" field.
func (pmu *PrescriptionMedicationUpdate) SetMedicationID(u uuid.UUID) *PrescriptionMedicationUpdate {
	pmu.mutation.SetMedicationID(u)
	return pmu
}

// SetNillableMedicationID sets the "medication_id" field if the given value is not nil.
func (pmu *PrescriptionMedicationUpdate) SetNillableMedicationID(u *uuid.UUID) *PrescriptionMedicationUpdate {
	if u != nil {
		pmu.SetMedicationID(*u)
	}
	return pmu
}

// SetQuantity sets the "quantity" field.
func (pmu *PrescriptionMedicationUpdate) SetQuantity(i int64) *PrescriptionMedicationUpdate {
	pmu.mutation.ResetQuantity()
	pmu.mutation.SetQuantity(i)
	return pmu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (pmu *PrescriptionMedicationUpdate) SetNillableQuantity(i *int64) *PrescriptionMedicationUpdate {
	if i != nil {
		pmu.SetQuantity(*i)
	}
	return pmu
}

// AddQuantity adds i to the "quantity" field.
func (pmu *PrescriptionMedicationUpdate) AddQuantity(i int64) *PrescriptionMedicationUpdate {
	pmu.mutation.AddQuantity(i)
	return pmu
}

// SetMedicalPrescriptionID sets the "medical_prescription" edge to the MedicalPrescription entity by ID.
func (pmu *PrescriptionMedicationUpdate) SetMedicalPrescriptionID(id uuid.UUID) *PrescriptionMedicationUpdate {
	pmu.mutation.SetMedicalPrescriptionID(id)
	return pmu
}

// SetMedicalPrescription sets the "medical_prescription" edge to the MedicalPrescription entity.
func (pmu *PrescriptionMedicationUpdate) SetMedicalPrescription(m *MedicalPrescription) *PrescriptionMedicationUpdate {
	return pmu.SetMedicalPrescriptionID(m.ID)
}

// SetMedication sets the "medication" edge to the Medication entity.
func (pmu *PrescriptionMedicationUpdate) SetMedication(m *Medication) *PrescriptionMedicationUpdate {
	return pmu.SetMedicationID(m.ID)
}

// Mutation returns the PrescriptionMedicationMutation object of the builder.
func (pmu *PrescriptionMedicationUpdate) Mutation() *PrescriptionMedicationMutation {
	return pmu.mutation
}

// ClearMedicalPrescription clears the "medical_prescription" edge to the MedicalPrescription entity.
func (pmu *PrescriptionMedicationUpdate) ClearMedicalPrescription() *PrescriptionMedicationUpdate {
	pmu.mutation.ClearMedicalPrescription()
	return pmu
}

// ClearMedication clears the "medication" edge to the Medication entity.
func (pmu *PrescriptionMedicationUpdate) ClearMedication() *PrescriptionMedicationUpdate {
	pmu.mutation.ClearMedication()
	return pmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pmu *PrescriptionMedicationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pmu.sqlSave, pmu.mutation, pmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmu *PrescriptionMedicationUpdate) SaveX(ctx context.Context) int {
	affected, err := pmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pmu *PrescriptionMedicationUpdate) Exec(ctx context.Context) error {
	_, err := pmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmu *PrescriptionMedicationUpdate) ExecX(ctx context.Context) {
	if err := pmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmu *PrescriptionMedicationUpdate) check() error {
	if pmu.mutation.MedicalPrescriptionCleared() && len(pmu.mutation.MedicalPrescriptionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PrescriptionMedication.medical_prescription"`)
	}
	if pmu.mutation.MedicationCleared() && len(pmu.mutation.MedicationIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PrescriptionMedication.medication"`)
	}
	return nil
}

func (pmu *PrescriptionMedicationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(prescriptionmedication.Table, prescriptionmedication.Columns, sqlgraph.NewFieldSpec(prescriptionmedication.FieldID, field.TypeUUID))
	if ps := pmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmu.mutation.Quantity(); ok {
		_spec.SetField(prescriptionmedication.FieldQuantity, field.TypeInt64, value)
	}
	if value, ok := pmu.mutation.AddedQuantity(); ok {
		_spec.AddField(prescriptionmedication.FieldQuantity, field.TypeInt64, value)
	}
	if pmu.mutation.MedicalPrescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionmedication.MedicalPrescriptionTable,
			Columns: []string{prescriptionmedication.MedicalPrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medicalprescription.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.MedicalPrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionmedication.MedicalPrescriptionTable,
			Columns: []string{prescriptionmedication.MedicalPrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medicalprescription.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pmu.mutation.MedicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionmedication.MedicationTable,
			Columns: []string{prescriptionmedication.MedicationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medication.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.MedicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionmedication.MedicationTable,
			Columns: []string{prescriptionmedication.MedicationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medication.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescriptionmedication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pmu.mutation.done = true
	return n, nil
}

// PrescriptionMedicationUpdateOne is the builder for updating a single PrescriptionMedication entity.
type PrescriptionMedicationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PrescriptionMedicationMutation
}

// SetPrescriptionID sets the "prescription_id" field.
func (pmuo *PrescriptionMedicationUpdateOne) SetPrescriptionID(u uuid.UUID) *PrescriptionMedicationUpdateOne {
	pmuo.mutation.SetPrescriptionID(u)
	return pmuo
}

// SetNillablePrescriptionID sets the "prescription_id" field if the given value is not nil.
func (pmuo *PrescriptionMedicationUpdateOne) SetNillablePrescriptionID(u *uuid.UUID) *PrescriptionMedicationUpdateOne {
	if u != nil {
		pmuo.SetPrescriptionID(*u)
	}
	return pmuo
}

// SetMedicationID sets the "medication_id" field.
func (pmuo *PrescriptionMedicationUpdateOne) SetMedicationID(u uuid.UUID) *PrescriptionMedicationUpdateOne {
	pmuo.mutation.SetMedicationID(u)
	return pmuo
}

// SetNillableMedicationID sets the "medication_id" field if the given value is not nil.
func (pmuo *PrescriptionMedicationUpdateOne) SetNillableMedicationID(u *uuid.UUID) *PrescriptionMedicationUpdateOne {
	if u != nil {
		pmuo.SetMedicationID(*u)
	}
	return pmuo
}

// SetQuantity sets the "quantity" field.
func (pmuo *PrescriptionMedicationUpdateOne) SetQuantity(i int64) *PrescriptionMedicationUpdateOne {
	pmuo.mutation.ResetQuantity()
	pmuo.mutation.SetQuantity(i)
	return pmuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (pmuo *PrescriptionMedicationUpdateOne) SetNillableQuantity(i *int64) *PrescriptionMedicationUpdateOne {
	if i != nil {
		pmuo.SetQuantity(*i)
	}
	return pmuo
}

// AddQuantity adds i to the "quantity" field.
func (pmuo *PrescriptionMedicationUpdateOne) AddQuantity(i int64) *PrescriptionMedicationUpdateOne {
	pmuo.mutation.AddQuantity(i)
	return pmuo
}

// SetMedicalPrescriptionID sets the "medical_prescription" edge to the MedicalPrescription entity by ID.
func (pmuo *PrescriptionMedicationUpdateOne) SetMedicalPrescriptionID(id uuid.UUID) *PrescriptionMedicationUpdateOne {
	pmuo.mutation.SetMedicalPrescriptionID(id)
	return pmuo
}

// SetMedicalPrescription sets the "medical_prescription" edge to the MedicalPrescription entity.
func (pmuo *PrescriptionMedicationUpdateOne) SetMedicalPrescription(m *MedicalPrescription) *PrescriptionMedicationUpdateOne {
	return pmuo.SetMedicalPrescriptionID(m.ID)
}

// SetMedication sets the "medication" edge to the Medication entity.
func (pmuo *PrescriptionMedicationUpdateOne) SetMedication(m *Medication) *PrescriptionMedicationUpdateOne {
	return pmuo.SetMedicationID(m.ID)
}

// Mutation returns the PrescriptionMedicationMutation object of the builder.
func (pmuo *PrescriptionMedicationUpdateOne) Mutation() *PrescriptionMedicationMutation {
	return pmuo.mutation
}

// ClearMedicalPrescription clears the "medical_prescription" edge to the MedicalPrescription entity.
func (pmuo *PrescriptionMedicationUpdateOne) ClearMedicalPrescription() *PrescriptionMedicationUpdateOne {
	pmuo.mutation.ClearMedicalPrescription()
	return pmuo
}

// ClearMedication clears the "medication" edge to the Medication entity.
func (pmuo *PrescriptionMedicationUpdateOne) ClearMedication() *PrescriptionMedicationUpdateOne {
	pmuo.mutation.ClearMedication()
	return pmuo
}

// Where appends a list predicates to the PrescriptionMedicationUpdate builder.
func (pmuo *PrescriptionMedicationUpdateOne) Where(ps ...predicate.PrescriptionMedication) *PrescriptionMedicationUpdateOne {
	pmuo.mutation.Where(ps...)
	return pmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pmuo *PrescriptionMedicationUpdateOne) Select(field string, fields ...string) *PrescriptionMedicationUpdateOne {
	pmuo.fields = append([]string{field}, fields...)
	return pmuo
}

// Save executes the query and returns the updated PrescriptionMedication entity.
func (pmuo *PrescriptionMedicationUpdateOne) Save(ctx context.Context) (*PrescriptionMedication, error) {
	return withHooks(ctx, pmuo.sqlSave, pmuo.mutation, pmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmuo *PrescriptionMedicationUpdateOne) SaveX(ctx context.Context) *PrescriptionMedication {
	node, err := pmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pmuo *PrescriptionMedicationUpdateOne) Exec(ctx context.Context) error {
	_, err := pmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmuo *PrescriptionMedicationUpdateOne) ExecX(ctx context.Context) {
	if err := pmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmuo *PrescriptionMedicationUpdateOne) check() error {
	if pmuo.mutation.MedicalPrescriptionCleared() && len(pmuo.mutation.MedicalPrescriptionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PrescriptionMedication.medical_prescription"`)
	}
	if pmuo.mutation.MedicationCleared() && len(pmuo.mutation.MedicationIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PrescriptionMedication.medication"`)
	}
	return nil
}

func (pmuo *PrescriptionMedicationUpdateOne) sqlSave(ctx context.Context) (_node *PrescriptionMedication, err error) {
	if err := pmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(prescriptionmedication.Table, prescriptionmedication.Columns, sqlgraph.NewFieldSpec(prescriptionmedication.FieldID, field.TypeUUID))
	id, ok := pmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PrescriptionMedication.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prescriptionmedication.FieldID)
		for _, f := range fields {
			if !prescriptionmedication.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != prescriptionmedication.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmuo.mutation.Quantity(); ok {
		_spec.SetField(prescriptionmedication.FieldQuantity, field.TypeInt64, value)
	}
	if value, ok := pmuo.mutation.AddedQuantity(); ok {
		_spec.AddField(prescriptionmedication.FieldQuantity, field.TypeInt64, value)
	}
	if pmuo.mutation.MedicalPrescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionmedication.MedicalPrescriptionTable,
			Columns: []string{prescriptionmedication.MedicalPrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medicalprescription.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.MedicalPrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionmedication.MedicalPrescriptionTable,
			Columns: []string{prescriptionmedication.MedicalPrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medicalprescription.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pmuo.mutation.MedicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionmedication.MedicationTable,
			Columns: []string{prescriptionmedication.MedicationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medication.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.MedicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionmedication.MedicationTable,
			Columns: []string{prescriptionmedication.MedicationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(medication.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PrescriptionMedication{config: pmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescriptionmedication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pmuo.mutation.done = true
	return _node, nil
}
