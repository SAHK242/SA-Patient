package schema

import "entgo.io/ent"

// Outpatient holds the schema definition for the Outpatient entity.
type Outpatient struct {
	ent.Schema
}

// Fields of the Outpatient.
func (Outpatient) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Outpatient.
func (Outpatient) Edges() []ent.Edge {
	return nil
}
