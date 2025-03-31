package schema

import "entgo.io/ent"

// ExamineDetail holds the schema definition for the ExamineDetail entity.
type ExamineDetail struct {
	ent.Schema
}

// Fields of the ExamineDetail.
func (ExamineDetail) Fields() []ent.Field {
	return nil
}

// Edges of the ExamineDetail.
func (ExamineDetail) Edges() []ent.Edge {
	return nil
}
