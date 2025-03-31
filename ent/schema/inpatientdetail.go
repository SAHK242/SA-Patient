package schema

import "entgo.io/ent"

// InpatientDetail holds the schema definition for the InpatientDetail entity.
type InpatientDetail struct {
	ent.Schema
}

// Fields of the InpatientDetail.
func (InpatientDetail) Fields() []ent.Field {
	return nil
}

// Edges of the InpatientDetail.
func (InpatientDetail) Edges() []ent.Edge {
	return nil
}
