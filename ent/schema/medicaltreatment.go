package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// MedicalTreatment holds the schema definition for the MedicalTreatment entity.
type MedicalTreatment struct {
	ent.Schema
}

// Fields of the MedicalTreatment.
func (MedicalTreatment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("medical_history_id", uuid.UUID{}),
		field.Time("start_date").Default(time.Now),
		field.Time("end_date").Nillable(),
		field.String("name").NotEmpty(),
		field.String("result"),
		field.String("description"),
		field.Float("fee").Default(0),
		field.UUID("main_doctor_id", uuid.UUID{}),
		field.String("support_doctor_ids").Nillable(),
		field.String("support_nurse_ids").Nillable(),
		field.Time("created_at").Default(time.Now),
		field.UUID("created_by", uuid.UUID{}),
		field.UUID("updated_by", uuid.UUID{}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the MedicalTreatment.
func (MedicalTreatment) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("medical_histories", MedicalHistories.Type, "medical_treatment", "medical_history_id"),
	}
}

func (MedicalTreatment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("medical_treatment"),
	}
}
