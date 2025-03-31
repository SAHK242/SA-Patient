package schema

import "entgo.io/ent"

// OutpatientDetail holds the schema definition for the OutpatientDetail entity.
type OutpatientDetail struct {
	ent.Schema
}

// Fields of the OutpatientDetail.
func (OutpatientDetail) Fields() []ent.Field {
	return nil
}

// Edges of the OutpatientDetail.
func (OutpatientDetail) Edges() []ent.Edge {
	return nil
}
