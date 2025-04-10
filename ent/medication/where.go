// Code generated by ent, DO NOT EDIT.

package medication

import (
	"patient/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldName, v))
}

// Effects applies equality check predicate on the "effects" field. It's identical to EffectsEQ.
func Effects(v string) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldEffects, v))
}

// ExpiredDate applies equality check predicate on the "expired_date" field. It's identical to ExpiredDateEQ.
func ExpiredDate(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldExpiredDate, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int64) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldQuantity, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldPrice, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Medication {
	return predicate.Medication(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Medication {
	return predicate.Medication(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Medication {
	return predicate.Medication(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Medication {
	return predicate.Medication(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Medication {
	return predicate.Medication(sql.FieldContainsFold(FieldName, v))
}

// EffectsEQ applies the EQ predicate on the "effects" field.
func EffectsEQ(v string) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldEffects, v))
}

// EffectsNEQ applies the NEQ predicate on the "effects" field.
func EffectsNEQ(v string) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldEffects, v))
}

// EffectsIn applies the In predicate on the "effects" field.
func EffectsIn(vs ...string) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldEffects, vs...))
}

// EffectsNotIn applies the NotIn predicate on the "effects" field.
func EffectsNotIn(vs ...string) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldEffects, vs...))
}

// EffectsGT applies the GT predicate on the "effects" field.
func EffectsGT(v string) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldEffects, v))
}

// EffectsGTE applies the GTE predicate on the "effects" field.
func EffectsGTE(v string) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldEffects, v))
}

// EffectsLT applies the LT predicate on the "effects" field.
func EffectsLT(v string) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldEffects, v))
}

// EffectsLTE applies the LTE predicate on the "effects" field.
func EffectsLTE(v string) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldEffects, v))
}

// EffectsContains applies the Contains predicate on the "effects" field.
func EffectsContains(v string) predicate.Medication {
	return predicate.Medication(sql.FieldContains(FieldEffects, v))
}

// EffectsHasPrefix applies the HasPrefix predicate on the "effects" field.
func EffectsHasPrefix(v string) predicate.Medication {
	return predicate.Medication(sql.FieldHasPrefix(FieldEffects, v))
}

// EffectsHasSuffix applies the HasSuffix predicate on the "effects" field.
func EffectsHasSuffix(v string) predicate.Medication {
	return predicate.Medication(sql.FieldHasSuffix(FieldEffects, v))
}

// EffectsEqualFold applies the EqualFold predicate on the "effects" field.
func EffectsEqualFold(v string) predicate.Medication {
	return predicate.Medication(sql.FieldEqualFold(FieldEffects, v))
}

// EffectsContainsFold applies the ContainsFold predicate on the "effects" field.
func EffectsContainsFold(v string) predicate.Medication {
	return predicate.Medication(sql.FieldContainsFold(FieldEffects, v))
}

// ExpiredDateEQ applies the EQ predicate on the "expired_date" field.
func ExpiredDateEQ(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldExpiredDate, v))
}

// ExpiredDateNEQ applies the NEQ predicate on the "expired_date" field.
func ExpiredDateNEQ(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldExpiredDate, v))
}

// ExpiredDateIn applies the In predicate on the "expired_date" field.
func ExpiredDateIn(vs ...time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldExpiredDate, vs...))
}

// ExpiredDateNotIn applies the NotIn predicate on the "expired_date" field.
func ExpiredDateNotIn(vs ...time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldExpiredDate, vs...))
}

// ExpiredDateGT applies the GT predicate on the "expired_date" field.
func ExpiredDateGT(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldExpiredDate, v))
}

// ExpiredDateGTE applies the GTE predicate on the "expired_date" field.
func ExpiredDateGTE(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldExpiredDate, v))
}

// ExpiredDateLT applies the LT predicate on the "expired_date" field.
func ExpiredDateLT(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldExpiredDate, v))
}

// ExpiredDateLTE applies the LTE predicate on the "expired_date" field.
func ExpiredDateLTE(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldExpiredDate, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int64) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int64) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int64) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int64) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int64) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int64) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int64) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int64) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldQuantity, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldPrice, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v uuid.UUID) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Medication {
	return predicate.Medication(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasPrescriptionMedication applies the HasEdge predicate on the "prescription_medication" edge.
func HasPrescriptionMedication() predicate.Medication {
	return predicate.Medication(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PrescriptionMedicationTable, PrescriptionMedicationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrescriptionMedicationWith applies the HasEdge predicate on the "prescription_medication" edge with a given conditions (other predicates).
func HasPrescriptionMedicationWith(preds ...predicate.PrescriptionMedication) predicate.Medication {
	return predicate.Medication(func(s *sql.Selector) {
		step := newPrescriptionMedicationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Medication) predicate.Medication {
	return predicate.Medication(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Medication) predicate.Medication {
	return predicate.Medication(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Medication) predicate.Medication {
	return predicate.Medication(sql.NotPredicates(p))
}
