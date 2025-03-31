package schema

import "entgo.io/ent"

// Medication holds the schema definition for the Medication entity.
type Medication struct {
	ent.Schema
}

// Fields of the Medication.
func (Medication) Fields() []ent.Field {
	return nil
}

// Edges of the Medication.
func (Medication) Edges() []ent.Edge {
	return nil
}
