package schema

import "entgo.io/ent"
import "entgo.io/ent/schema/field"
import "github.com/google/uuid"
import "entgo.io/ent/schema"

// Outpatient holds the schema definition for the Outpatient entity.
type Outpatient struct {
	ent.Schema
}

// Fields of the Outpatient.
func (Outpatient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("patient_id", uuid.UUID{}),
		field.Time("register_date"),
	}
}

// Edges of the Outpatient.
func (Outpatient) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("patient", Patient.Type, "outpatients", "patient_id"),
	}
}

func (Outpatient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("outpatient"),
	}
}
