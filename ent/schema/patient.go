package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	gcommon "patient/proto/gcommon"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("phone_number"),
		field.String("first_name"),
		field.String("last_name"),
		field.Int32("gender").Positive(),
		field.String("address"),
		field.Time("date_of_birth"),
		field.Int32("current_patient_type").Default(int32(gcommon.PatientType_PATIENT_TYPE_OUTPATIENT)),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		hasMany("inpatients", Inpatient.Type),
		hasMany("outpatients", Outpatient.Type),
	}
}

func (Patient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("patient"),
	}
}
