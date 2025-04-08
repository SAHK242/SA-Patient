package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// MedicalSurgery holds the schema definition for the MedicalSurgery entity.
type MedicalSurgery struct {
	ent.Schema
}

// Fields of the MedicalSurgery.
func (MedicalSurgery) Fields() []ent.Field {
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

// Edges of the MedicalSurgery.
func (MedicalSurgery) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("medical_histories", MedicalHistories.Type, "medical_surgery", "medical_history_id"),
	}
}

func (MedicalSurgery) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("medical_surgery"),
	}
}
