package schema

import "entgo.io/ent"

// MedicationEffect holds the schema definition for the MedicationEffect entity.
type MedicationEffect struct {
	ent.Schema
}

// Fields of the MedicationEffect.
func (MedicationEffect) Fields() []ent.Field {
	return nil
}

// Edges of the MedicationEffect.
func (MedicationEffect) Edges() []ent.Edge {
	return nil
}
