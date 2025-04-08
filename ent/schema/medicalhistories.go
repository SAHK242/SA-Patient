package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// MedicalHistories holds the schema definition for the MedicalHistories entity.
type MedicalHistories struct {
	ent.Schema
}

// Fields of the MedicalHistories.
func (MedicalHistories) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("patient_id", uuid.UUID{}),
		field.String("reason").NotEmpty(),
		field.String("diagnosis").NotEmpty(),
		field.Bool("has_treatment").Default(false),
		field.Bool("has_surgery").Default(false),
		field.Bool("has_prescription").Default(false),
		field.String("doctor_notes").Optional(),
		field.Time("medical_end_date").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.UUID("created_by", uuid.UUID{}),
		field.UUID("updated_by", uuid.UUID{}),
	}
}

// Edges of the MedicalHistories.
func (MedicalHistories) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("patient", Patient.Type, "medical_history", "patient_id"),
		hasMany("medical_prescription", MedicalPrescription.Type),
		hasMany("medical_treatment", MedicalTreatment.Type),
		hasMany("medical_surgery", MedicalSurgery.Type),
	}
}

func (MedicalHistories) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("medical_histories"),
	}
}
