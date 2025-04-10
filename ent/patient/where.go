// Code generated by ent, DO NOT EDIT.

package patient

import (
	"patient/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldID, id))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldPhone, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldLastName, v))
}

// Gender applies equality check predicate on the "gender" field. It's identical to GenderEQ.
func Gender(v int32) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldGender, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldAddress, v))
}

// DateOfBirth applies equality check predicate on the "date_of_birth" field. It's identical to DateOfBirthEQ.
func DateOfBirth(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldDateOfBirth, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldUpdatedBy, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldPhone, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldLastName, v))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v int32) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v int32) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...int32) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...int32) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldGender, vs...))
}

// GenderGT applies the GT predicate on the "gender" field.
func GenderGT(v int32) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldGender, v))
}

// GenderGTE applies the GTE predicate on the "gender" field.
func GenderGTE(v int32) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldGender, v))
}

// GenderLT applies the LT predicate on the "gender" field.
func GenderLT(v int32) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldGender, v))
}

// GenderLTE applies the LTE predicate on the "gender" field.
func GenderLTE(v int32) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldGender, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldAddress, v))
}

// DateOfBirthEQ applies the EQ predicate on the "date_of_birth" field.
func DateOfBirthEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldDateOfBirth, v))
}

// DateOfBirthNEQ applies the NEQ predicate on the "date_of_birth" field.
func DateOfBirthNEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldDateOfBirth, v))
}

// DateOfBirthIn applies the In predicate on the "date_of_birth" field.
func DateOfBirthIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldDateOfBirth, vs...))
}

// DateOfBirthNotIn applies the NotIn predicate on the "date_of_birth" field.
func DateOfBirthNotIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldDateOfBirth, vs...))
}

// DateOfBirthGT applies the GT predicate on the "date_of_birth" field.
func DateOfBirthGT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldDateOfBirth, v))
}

// DateOfBirthGTE applies the GTE predicate on the "date_of_birth" field.
func DateOfBirthGTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldDateOfBirth, v))
}

// DateOfBirthLT applies the LT predicate on the "date_of_birth" field.
func DateOfBirthLT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldDateOfBirth, v))
}

// DateOfBirthLTE applies the LTE predicate on the "date_of_birth" field.
func DateOfBirthLTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldDateOfBirth, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v uuid.UUID) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldUpdatedBy, v))
}

// HasMedicalHistory applies the HasEdge predicate on the "medical_history" edge.
func HasMedicalHistory() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MedicalHistoryTable, MedicalHistoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMedicalHistoryWith applies the HasEdge predicate on the "medical_history" edge with a given conditions (other predicates).
func HasMedicalHistoryWith(preds ...predicate.MedicalHistories) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := newMedicalHistoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Patient) predicate.Patient {
	return predicate.Patient(sql.NotPredicates(p))
}
