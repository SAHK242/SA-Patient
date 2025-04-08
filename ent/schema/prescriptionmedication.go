package schema

import "entgo.io/ent"
import "entgo.io/ent/schema"
import "entgo.io/ent/schema/field"
import "github.com/google/uuid"

// PrescriptionMedication holds the schema definition for the PrescriptionMedication entity.
type PrescriptionMedication struct {
	ent.Schema
}

// Fields of the PrescriptionMedication.
func (PrescriptionMedication) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("prescription_id", uuid.UUID{}),
		field.UUID("medication_id", uuid.UUID{}),
		field.Int64("quantity").Default(0),
	}
}

// Edges of the PrescriptionMedication.
func (PrescriptionMedication) Edges() []ent.Edge {
	return []ent.Edge{
		mustBelongToOne("medical_prescription", MedicalPrescription.Type, "prescription_medication", "prescription_id"),
		mustBelongToOne("medication", Medication.Type, "prescription_medication", "medication_id"),
	}
}

func (PrescriptionMedication) Annotations() []schema.Annotation {
	return []schema.Annotation{
		table("prescription_medication"),
	}
}
