package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// MedicalPrescription holds the schema definition for the MedicalPrescription entity.
type MedicalPrescription struct {
	ent.Schema
}

// Fields of the MedicalPrescription.
func (MedicalPrescription) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("medical_history_id", uuid.UUID{}),
		field.Time("prescription_date").Default(time.Now),
		field.Float("fee").Default(0),
		field.Time("created_at").Default(time.Now),
		field.UUID("created_by", uuid.UUID{}),
	}
}

// Edges of the MedicalPrescription.
func (MedicalPrescription) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("medical_histories", MedicalHistories.Type, "medical_prescription", "medical_history_id"),
		hasMany("prescription_medication", PrescriptionMedication.Type),
	}
}

func (MedicalPrescription) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("medical_prescription"),
	}
}
