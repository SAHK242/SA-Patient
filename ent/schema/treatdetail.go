package schema

import "entgo.io/ent"

// TreatDetail holds the schema definition for the TreatDetail entity.
type TreatDetail struct {
	ent.Schema
}

// Fields of the TreatDetail.
func (TreatDetail) Fields() []ent.Field {
	return nil
}

// Edges of the TreatDetail.
func (TreatDetail) Edges() []ent.Edge {
	return nil
}
