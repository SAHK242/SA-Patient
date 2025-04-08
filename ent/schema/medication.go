package schema

import (
	"entgo.io/ent"
	"time"
)
import "entgo.io/ent/schema"
import "entgo.io/ent/schema/field"
import "github.com/google/uuid"

// Medication holds the schema definition for the Medication entity.
type Medication struct {
	ent.Schema
}

// Fields of the Medication.
func (Medication) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.String("effects"),
		field.Time("expired_date"),
		field.Int64("quantity").Default(0),
		field.Float("price").Default(0),
		field.UUID("created_by", uuid.UUID{}),
		field.UUID("updated_by", uuid.UUID{}),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Medication.
func (Medication) Edges() []ent.Edge {
	return []ent.Edge{
		hasMany("prescription_medication", PrescriptionMedication.Type),
	}
}

func (Medication) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("medication"),
	}
}
