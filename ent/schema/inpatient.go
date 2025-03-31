package schema

import "entgo.io/ent"

// Inpatient holds the schema definition for the Inpatient entity.
type Inpatient struct {
	ent.Schema
}

// Fields of the Inpatient.
func (Inpatient) Fields() []ent.Field {
	return nil
}

// Edges of the Inpatient.
func (Inpatient) Edges() []ent.Edge {
	return nil
}
