package schema

import "entgo.io/ent"
import "entgo.io/ent/schema/field"
import "github.com/google/uuid"
import "entgo.io/ent/schema"

// Inpatient holds the schema definition for the Inpatient entity.
type Inpatient struct {
	ent.Schema
}

// Fields of the Inpatient.
func (Inpatient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("patient_id", uuid.UUID{}),
		field.Time("register_date"),
	}
}

// Edges of the Inpatient.
func (Inpatient) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("patient", Patient.Type, "inpatients", "patient_id"),
	}
}

func (Inpatient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("inpatient"),
	}
}
