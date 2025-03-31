package schema

import "entgo.io/ent"

// TreatMedication holds the schema definition for the TreatMedication entity.
type TreatMedication struct {
	ent.Schema
}

// Fields of the TreatMedication.
func (TreatMedication) Fields() []ent.Field {
	return nil
}

// Edges of the TreatMedication.
func (TreatMedication) Edges() []ent.Edge {
	return nil
}
