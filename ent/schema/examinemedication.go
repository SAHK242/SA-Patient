package schema

import "entgo.io/ent"

// ExamineMedication holds the schema definition for the ExamineMedication entity.
type ExamineMedication struct {
	ent.Schema
}

// Fields of the ExamineMedication.
func (ExamineMedication) Fields() []ent.Field {
	return nil
}

// Edges of the ExamineMedication.
func (ExamineMedication) Edges() []ent.Edge {
	return nil
}
