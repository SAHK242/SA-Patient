package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return nil
}

func (Patient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("patient"),
	}
}
